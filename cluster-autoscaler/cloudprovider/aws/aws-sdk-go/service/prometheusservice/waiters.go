// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package prometheusservice

import (
	"time"

	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws/request"
)

// WaitUntilScraperActive uses the Amazon Prometheus Service API operation
// DescribeScraper to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *PrometheusService) WaitUntilScraperActive(input *DescribeScraperInput) error {
	return c.WaitUntilScraperActiveWithContext(aws.BackgroundContext(), input)
}

// WaitUntilScraperActiveWithContext is an extended version of WaitUntilScraperActive.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *PrometheusService) WaitUntilScraperActiveWithContext(ctx aws.Context, input *DescribeScraperInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilScraperActive",
		MaxAttempts: 60,
		Delay:       request.ConstantWaiterDelay(2 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:   request.SuccessWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "scraper.status.statusCode",
				Expected: "ACTIVE",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "scraper.status.statusCode",
				Expected: "CREATION_FAILED",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *DescribeScraperInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.DescribeScraperRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}

// WaitUntilScraperDeleted uses the Amazon Prometheus Service API operation
// DescribeScraper to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *PrometheusService) WaitUntilScraperDeleted(input *DescribeScraperInput) error {
	return c.WaitUntilScraperDeletedWithContext(aws.BackgroundContext(), input)
}

// WaitUntilScraperDeletedWithContext is an extended version of WaitUntilScraperDeleted.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *PrometheusService) WaitUntilScraperDeletedWithContext(ctx aws.Context, input *DescribeScraperInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilScraperDeleted",
		MaxAttempts: 60,
		Delay:       request.ConstantWaiterDelay(2 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:    request.SuccessWaiterState,
				Matcher:  request.ErrorWaiterMatch,
				Expected: "ResourceNotFoundException",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "scraper.status.statusCode",
				Expected: "DELETION_FAILED",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *DescribeScraperInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.DescribeScraperRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}

// WaitUntilWorkspaceActive uses the Amazon Prometheus Service API operation
// DescribeWorkspace to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *PrometheusService) WaitUntilWorkspaceActive(input *DescribeWorkspaceInput) error {
	return c.WaitUntilWorkspaceActiveWithContext(aws.BackgroundContext(), input)
}

// WaitUntilWorkspaceActiveWithContext is an extended version of WaitUntilWorkspaceActive.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *PrometheusService) WaitUntilWorkspaceActiveWithContext(ctx aws.Context, input *DescribeWorkspaceInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilWorkspaceActive",
		MaxAttempts: 60,
		Delay:       request.ConstantWaiterDelay(2 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:   request.SuccessWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "workspace.status.statusCode",
				Expected: "ACTIVE",
			},
			{
				State:   request.RetryWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "workspace.status.statusCode",
				Expected: "UPDATING",
			},
			{
				State:   request.RetryWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "workspace.status.statusCode",
				Expected: "CREATING",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *DescribeWorkspaceInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.DescribeWorkspaceRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}

// WaitUntilWorkspaceDeleted uses the Amazon Prometheus Service API operation
// DescribeWorkspace to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *PrometheusService) WaitUntilWorkspaceDeleted(input *DescribeWorkspaceInput) error {
	return c.WaitUntilWorkspaceDeletedWithContext(aws.BackgroundContext(), input)
}

// WaitUntilWorkspaceDeletedWithContext is an extended version of WaitUntilWorkspaceDeleted.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *PrometheusService) WaitUntilWorkspaceDeletedWithContext(ctx aws.Context, input *DescribeWorkspaceInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilWorkspaceDeleted",
		MaxAttempts: 60,
		Delay:       request.ConstantWaiterDelay(2 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:    request.SuccessWaiterState,
				Matcher:  request.ErrorWaiterMatch,
				Expected: "ResourceNotFoundException",
			},
			{
				State:   request.RetryWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "workspace.status.statusCode",
				Expected: "DELETING",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *DescribeWorkspaceInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.DescribeWorkspaceRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}
