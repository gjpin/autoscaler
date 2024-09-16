/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package openstack

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	apiv1 "k8s.io/api/core/v1"

	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/openstack/hcloud-go/hcloud"
	"k8s.io/autoscaler/cluster-autoscaler/version"
)

var (
	httpClient = &http.Client{
		Transport: instrumentedRoundTripper(),
	}
)

// openstackManager handles OpenStack communication and data caching of
// node groups
type openstackManager struct {
	client           *hcloud.Client
	nodeGroups       map[string]*openstackNodeGroup
	apiCallContext   context.Context
	clusterConfig    *ClusterConfig
	sshKey           *hcloud.SSHKey
	network          *hcloud.Network
	firewall         *hcloud.Firewall
	createTimeout    time.Duration
	publicIPv4       bool
	publicIPv6       bool
	cachedServerType *serverTypeCache
	cachedServers    *serversCache
}

// ClusterConfig holds the configuration for all the nodepools
type ClusterConfig struct {
	ImagesForArch    ImageList
	NodeConfigs      map[string]*NodeConfig
	IsUsingNewFormat bool
	LegacyConfig     LegacyConfig
}

// ImageList holds the image id/names for the different architectures
type ImageList struct {
	Arm64 string
	Amd64 string
}

// NodeConfig holds the configuration for a single nodepool
type NodeConfig struct {
	CloudInit string
	Taints    []apiv1.Taint
	Labels    map[string]string
}

// LegacyConfig holds the configuration in the legacy format
type LegacyConfig struct {
	CloudInit string
	ImageName string
}

func newManager() (*openstackManager, error) {
	token := os.Getenv("HCLOUD_TOKEN")
	if token == "" {
		return nil, errors.New("`HCLOUD_TOKEN` is not specified")
	}

	client := hcloud.NewClient(
		hcloud.WithToken(token),
		hcloud.WithHTTPClient(httpClient),
		hcloud.WithApplication("cluster-autoscaler", version.ClusterAutoscalerVersion),
		hcloud.WithPollBackoffFunc(hcloud.ExponentialBackoff(2, 500*time.Millisecond)),
		hcloud.WithDebugWriter(&debugWriter{}),
	)

	ctx := context.Background()
	var err error

	clusterConfigBase64 := os.Getenv("HCLOUD_CLUSTER_CONFIG")
	cloudInitBase64 := os.Getenv("HCLOUD_CLOUD_INIT")

	if clusterConfigBase64 == "" && cloudInitBase64 == "" {
		return nil, errors.New("`HCLOUD_CLUSTER_CONFIG` or `HCLOUD_CLOUD_INIT` is not specified")
	}
	var clusterConfig *ClusterConfig = &ClusterConfig{}

	if clusterConfigBase64 != "" {
		clusterConfig.IsUsingNewFormat = true
	}

	if clusterConfig.IsUsingNewFormat {
		clusterConfigEnv, err := base64.StdEncoding.DecodeString(clusterConfigBase64)
		if err != nil {
			return nil, fmt.Errorf("failed to parse cluster config error: %s", err)
		}
		err = json.Unmarshal(clusterConfigEnv, &clusterConfig)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal cluster config JSON: %s", err)
		}
	}

	if !clusterConfig.IsUsingNewFormat {
		cloudInit, err := base64.StdEncoding.DecodeString(cloudInitBase64)
		if err != nil {
			return nil, fmt.Errorf("failed to parse cloud init error: %s", err)
		}

		imageName := os.Getenv("HCLOUD_IMAGE")
		if imageName == "" {
			imageName = "ubuntu-20.04"
		}

		clusterConfig.LegacyConfig.CloudInit = string(cloudInit)
		clusterConfig.LegacyConfig.ImageName = imageName
	}

	publicIPv4 := true
	publicIPv4Str := os.Getenv("HCLOUD_PUBLIC_IPV4")
	if publicIPv4Str != "" {
		publicIPv4, err = strconv.ParseBool(publicIPv4Str)
		if err != nil {
			return nil, fmt.Errorf("failed to parse HCLOUD_PUBLIC_IPV4: %s", err)
		}
	}

	publicIPv6 := true
	publicIPv6Str := os.Getenv("HCLOUD_PUBLIC_IPV6")
	if publicIPv6Str != "" {
		publicIPv6, err = strconv.ParseBool(publicIPv6Str)
		if err != nil {
			return nil, fmt.Errorf("failed to parse HCLOUD_PUBLIC_IPV6: %s", err)
		}
	}

	var sshKey *hcloud.SSHKey
	sshKeyIdOrName := os.Getenv("HCLOUD_SSH_KEY")
	if sshKeyIdOrName != "" {
		sshKey, _, err = client.SSHKey.Get(ctx, sshKeyIdOrName)
		if err != nil {
			return nil, fmt.Errorf("failed to get ssh key error: %s", err)
		}
	}

	var network *hcloud.Network
	networkIdOrName := os.Getenv("HCLOUD_NETWORK")
	if networkIdOrName != "" {
		network, _, err = client.Network.Get(ctx, networkIdOrName)
		if err != nil {
			return nil, fmt.Errorf("failed to get network error: %s", err)
		}

	}

	createTimeout := serverCreateTimeoutDefault
	v, err := strconv.Atoi(os.Getenv("HCLOUD_SERVER_CREATION_TIMEOUT"))
	if err == nil && v != 0 {
		createTimeout = time.Duration(v) * time.Minute
	}

	var firewall *hcloud.Firewall
	firewallIdOrName := os.Getenv("HCLOUD_FIREWALL")
	if firewallIdOrName != "" {
		firewall, _, err = client.Firewall.Get(ctx, firewallIdOrName)
		if err != nil {
			return nil, fmt.Errorf("failed to get firewall error: %s", err)
		}
	}

	m := &openstackManager{
		client:           client,
		nodeGroups:       make(map[string]*openstackNodeGroup),
		sshKey:           sshKey,
		network:          network,
		firewall:         firewall,
		createTimeout:    createTimeout,
		apiCallContext:   ctx,
		publicIPv4:       publicIPv4,
		publicIPv6:       publicIPv6,
		clusterConfig:    clusterConfig,
		cachedServerType: newServerTypeCache(ctx, client),
		cachedServers:    newServersCache(ctx, client),
	}

	m.nodeGroups[drainingNodePoolId] = &openstackNodeGroup{
		manager:      m,
		instanceType: "cx11",
		region:       "fsn1",
		targetSize:   0,
		maxSize:      0,
		minSize:      0,
		id:           drainingNodePoolId,
	}

	return m, nil
}

// Refresh refreshes the cache holding the nodegroups. This is called by the CA
// based on the `--scan-interval`. By default it's 10 seconds.
func (m *openstackManager) Refresh() error {
	return nil
}

func (m *openstackManager) allServers(nodeGroup string) ([]*hcloud.Server, error) {
	servers, err := m.cachedServers.getServersByNodeGroupName(nodeGroup)
	if err != nil {
		return nil, fmt.Errorf("failed to get servers for hcloud: %v", err)
	}

	return servers, nil
}

func (m *openstackManager) deleteByNode(node *apiv1.Node) error {
	server, err := m.serverForNode(node)
	if err != nil {
		return fmt.Errorf("failed to delete node %s error: %v", node.Name, err)
	}

	if server == nil {
		return fmt.Errorf("failed to delete node %s server not found", node.Name)
	}

	return m.deleteServer(server)
}

func (m *openstackManager) deleteServer(server *hcloud.Server) error {
	_, err := m.client.Server.Delete(m.apiCallContext, server)
	return err
}

func (m *openstackManager) addNodeToDrainingPool(node *apiv1.Node) (*openstackNodeGroup, error) {
	m.nodeGroups[drainingNodePoolId].targetSize += 1
	return m.nodeGroups[drainingNodePoolId], nil
}

func (m *openstackManager) validProviderID(providerID string) bool {
	return strings.HasPrefix(providerID, providerIDPrefix)
}

func (m *openstackManager) serverForNode(node *apiv1.Node) (*hcloud.Server, error) {
	var nodeIdOrName string
	if node.Spec.ProviderID != "" {
		if !m.validProviderID(node.Spec.ProviderID) {
			// This cluster-autoscaler provider only handles OpenStack Cloud servers.
			// Any other provider ID prefix is invalid, and we return no server. Returning an error here breaks hybrid
			// clusters with nodes from OpenStack Cloud & Robot (or other providers).
			return nil, nil
		}
		nodeIdOrName = strings.TrimPrefix(node.Spec.ProviderID, providerIDPrefix)
	} else {
		nodeIdOrName = node.Name
	}

	server, err := m.cachedServers.getServer(nodeIdOrName)
	if err != nil {
		return nil, fmt.Errorf("failed to get servers for node %s error: %v", node.Name, err)
	}
	return server, nil
}
