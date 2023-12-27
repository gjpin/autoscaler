// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package devopsguru provides the client and types for making API
// requests to Amazon DevOps Guru.
//
// Amazon DevOps Guru is a fully managed service that helps you identify anomalous
// behavior in business critical operational applications. You specify the Amazon
// Web Services resources that you want DevOps Guru to cover, then the Amazon
// CloudWatch metrics and Amazon Web Services CloudTrail events related to those
// resources are analyzed. When anomalous behavior is detected, DevOps Guru
// creates an insight that includes recommendations, related events, and related
// metrics that can help you improve your operational applications. For more
// information, see What is Amazon DevOps Guru (https://docs.aws.amazon.com/devops-guru/latest/userguide/welcome.html).
//
// You can specify 1 or 2 Amazon Simple Notification Service topics so you are
// notified every time a new insight is created. You can also enable DevOps
// Guru to generate an OpsItem in Amazon Web Services Systems Manager for each
// insight to help you manage and track your work addressing insights.
//
// To learn about the DevOps Guru workflow, see How DevOps Guru works (https://docs.aws.amazon.com/devops-guru/latest/userguide/welcome.html#how-it-works).
// To learn about DevOps Guru concepts, see Concepts in DevOps Guru (https://docs.aws.amazon.com/devops-guru/latest/userguide/concepts.html).
//
// See https://docs.aws.amazon.com/goto/WebAPI/devops-guru-2020-12-01 for more information on this service.
//
// See devopsguru package documentation for more information.
// https://docs.aws.amazon.com/sdk-for-go/api/service/devopsguru/
//
// # Using the Client
//
// To contact Amazon DevOps Guru with the SDK use the New function to create
// a new service client. With that client you can make API requests to the service.
// These clients are safe to use concurrently.
//
// See the SDK's documentation for more information on how to use the SDK.
// https://docs.aws.amazon.com/sdk-for-go/api/
//
// See aws.Config documentation for more information on configuring SDK clients.
// https://docs.aws.amazon.com/sdk-for-go/api/aws/#Config
//
// See the Amazon DevOps Guru client DevOpsGuru for more
// information on creating client for this service.
// https://docs.aws.amazon.com/sdk-for-go/api/service/devopsguru/#New
package devopsguru
