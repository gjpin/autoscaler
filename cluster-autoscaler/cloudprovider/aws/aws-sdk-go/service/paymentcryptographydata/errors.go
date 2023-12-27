// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package paymentcryptographydata

import (
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeAccessDeniedException for service response error code
	// "AccessDeniedException".
	//
	// You do not have sufficient access to perform this action.
	ErrCodeAccessDeniedException = "AccessDeniedException"

	// ErrCodeInternalServerException for service response error code
	// "InternalServerException".
	//
	// The request processing has failed because of an unknown error, exception,
	// or failure.
	ErrCodeInternalServerException = "InternalServerException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// The request was denied due to an invalid resource error.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"

	// ErrCodeThrottlingException for service response error code
	// "ThrottlingException".
	//
	// The request was denied due to request throttling.
	ErrCodeThrottlingException = "ThrottlingException"

	// ErrCodeValidationException for service response error code
	// "ValidationException".
	//
	// The request was denied due to an invalid request error.
	ErrCodeValidationException = "ValidationException"

	// ErrCodeVerificationFailedException for service response error code
	// "VerificationFailedException".
	//
	// This request failed verification.
	ErrCodeVerificationFailedException = "VerificationFailedException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"AccessDeniedException":       newErrorAccessDeniedException,
	"InternalServerException":     newErrorInternalServerException,
	"ResourceNotFoundException":   newErrorResourceNotFoundException,
	"ThrottlingException":         newErrorThrottlingException,
	"ValidationException":         newErrorValidationException,
	"VerificationFailedException": newErrorVerificationFailedException,
}
