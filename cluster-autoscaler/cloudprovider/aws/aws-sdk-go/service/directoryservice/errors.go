// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package directoryservice

import (
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeAccessDeniedException for service response error code
	// "AccessDeniedException".
	//
	// Client authentication is not available in this region at this time.
	ErrCodeAccessDeniedException = "AccessDeniedException"

	// ErrCodeAuthenticationFailedException for service response error code
	// "AuthenticationFailedException".
	//
	// An authentication error occurred.
	ErrCodeAuthenticationFailedException = "AuthenticationFailedException"

	// ErrCodeCertificateAlreadyExistsException for service response error code
	// "CertificateAlreadyExistsException".
	//
	// The certificate has already been registered into the system.
	ErrCodeCertificateAlreadyExistsException = "CertificateAlreadyExistsException"

	// ErrCodeCertificateDoesNotExistException for service response error code
	// "CertificateDoesNotExistException".
	//
	// The certificate is not present in the system for describe or deregister activities.
	ErrCodeCertificateDoesNotExistException = "CertificateDoesNotExistException"

	// ErrCodeCertificateInUseException for service response error code
	// "CertificateInUseException".
	//
	// The certificate is being used for the LDAP security connection and cannot
	// be removed without disabling LDAP security.
	ErrCodeCertificateInUseException = "CertificateInUseException"

	// ErrCodeCertificateLimitExceededException for service response error code
	// "CertificateLimitExceededException".
	//
	// The certificate could not be added because the certificate limit has been
	// reached.
	ErrCodeCertificateLimitExceededException = "CertificateLimitExceededException"

	// ErrCodeClientException for service response error code
	// "ClientException".
	//
	// A client exception has occurred.
	ErrCodeClientException = "ClientException"

	// ErrCodeDirectoryAlreadyInRegionException for service response error code
	// "DirectoryAlreadyInRegionException".
	//
	// The Region you specified is the same Region where the Managed Microsoft AD
	// directory was created. Specify a different Region and try again.
	ErrCodeDirectoryAlreadyInRegionException = "DirectoryAlreadyInRegionException"

	// ErrCodeDirectoryAlreadySharedException for service response error code
	// "DirectoryAlreadySharedException".
	//
	// The specified directory has already been shared with this Amazon Web Services
	// account.
	ErrCodeDirectoryAlreadySharedException = "DirectoryAlreadySharedException"

	// ErrCodeDirectoryDoesNotExistException for service response error code
	// "DirectoryDoesNotExistException".
	//
	// The specified directory does not exist in the system.
	ErrCodeDirectoryDoesNotExistException = "DirectoryDoesNotExistException"

	// ErrCodeDirectoryInDesiredStateException for service response error code
	// "DirectoryInDesiredStateException".
	//
	// The directory is already updated to desired update type settings.
	ErrCodeDirectoryInDesiredStateException = "DirectoryInDesiredStateException"

	// ErrCodeDirectoryLimitExceededException for service response error code
	// "DirectoryLimitExceededException".
	//
	// The maximum number of directories in the region has been reached. You can
	// use the GetDirectoryLimits operation to determine your directory limits in
	// the region.
	ErrCodeDirectoryLimitExceededException = "DirectoryLimitExceededException"

	// ErrCodeDirectoryNotSharedException for service response error code
	// "DirectoryNotSharedException".
	//
	// The specified directory has not been shared with this Amazon Web Services
	// account.
	ErrCodeDirectoryNotSharedException = "DirectoryNotSharedException"

	// ErrCodeDirectoryUnavailableException for service response error code
	// "DirectoryUnavailableException".
	//
	// The specified directory is unavailable or could not be found.
	ErrCodeDirectoryUnavailableException = "DirectoryUnavailableException"

	// ErrCodeDomainControllerLimitExceededException for service response error code
	// "DomainControllerLimitExceededException".
	//
	// The maximum allowed number of domain controllers per directory was exceeded.
	// The default limit per directory is 20 domain controllers.
	ErrCodeDomainControllerLimitExceededException = "DomainControllerLimitExceededException"

	// ErrCodeEntityAlreadyExistsException for service response error code
	// "EntityAlreadyExistsException".
	//
	// The specified entity already exists.
	ErrCodeEntityAlreadyExistsException = "EntityAlreadyExistsException"

	// ErrCodeEntityDoesNotExistException for service response error code
	// "EntityDoesNotExistException".
	//
	// The specified entity could not be found.
	ErrCodeEntityDoesNotExistException = "EntityDoesNotExistException"

	// ErrCodeIncompatibleSettingsException for service response error code
	// "IncompatibleSettingsException".
	//
	// The specified directory setting is not compatible with other settings.
	ErrCodeIncompatibleSettingsException = "IncompatibleSettingsException"

	// ErrCodeInsufficientPermissionsException for service response error code
	// "InsufficientPermissionsException".
	//
	// The account does not have sufficient permission to perform the operation.
	ErrCodeInsufficientPermissionsException = "InsufficientPermissionsException"

	// ErrCodeInvalidCertificateException for service response error code
	// "InvalidCertificateException".
	//
	// The certificate PEM that was provided has incorrect encoding.
	ErrCodeInvalidCertificateException = "InvalidCertificateException"

	// ErrCodeInvalidClientAuthStatusException for service response error code
	// "InvalidClientAuthStatusException".
	//
	// Client authentication is already enabled.
	ErrCodeInvalidClientAuthStatusException = "InvalidClientAuthStatusException"

	// ErrCodeInvalidLDAPSStatusException for service response error code
	// "InvalidLDAPSStatusException".
	//
	// The LDAP activities could not be performed because they are limited by the
	// LDAPS status.
	ErrCodeInvalidLDAPSStatusException = "InvalidLDAPSStatusException"

	// ErrCodeInvalidNextTokenException for service response error code
	// "InvalidNextTokenException".
	//
	// The NextToken value is not valid.
	ErrCodeInvalidNextTokenException = "InvalidNextTokenException"

	// ErrCodeInvalidParameterException for service response error code
	// "InvalidParameterException".
	//
	// One or more parameters are not valid.
	ErrCodeInvalidParameterException = "InvalidParameterException"

	// ErrCodeInvalidPasswordException for service response error code
	// "InvalidPasswordException".
	//
	// The new password provided by the user does not meet the password complexity
	// requirements defined in your directory.
	ErrCodeInvalidPasswordException = "InvalidPasswordException"

	// ErrCodeInvalidTargetException for service response error code
	// "InvalidTargetException".
	//
	// The specified shared target is not valid.
	ErrCodeInvalidTargetException = "InvalidTargetException"

	// ErrCodeIpRouteLimitExceededException for service response error code
	// "IpRouteLimitExceededException".
	//
	// The maximum allowed number of IP addresses was exceeded. The default limit
	// is 100 IP address blocks.
	ErrCodeIpRouteLimitExceededException = "IpRouteLimitExceededException"

	// ErrCodeNoAvailableCertificateException for service response error code
	// "NoAvailableCertificateException".
	//
	// Client authentication setup could not be completed because at least one valid
	// certificate must be registered in the system.
	ErrCodeNoAvailableCertificateException = "NoAvailableCertificateException"

	// ErrCodeOrganizationsException for service response error code
	// "OrganizationsException".
	//
	// Exception encountered while trying to access your Amazon Web Services organization.
	ErrCodeOrganizationsException = "OrganizationsException"

	// ErrCodeRegionLimitExceededException for service response error code
	// "RegionLimitExceededException".
	//
	// You have reached the limit for maximum number of simultaneous Region replications
	// per directory.
	ErrCodeRegionLimitExceededException = "RegionLimitExceededException"

	// ErrCodeServiceException for service response error code
	// "ServiceException".
	//
	// An exception has occurred in Directory Service.
	ErrCodeServiceException = "ServiceException"

	// ErrCodeShareLimitExceededException for service response error code
	// "ShareLimitExceededException".
	//
	// The maximum number of Amazon Web Services accounts that you can share with
	// this directory has been reached.
	ErrCodeShareLimitExceededException = "ShareLimitExceededException"

	// ErrCodeSnapshotLimitExceededException for service response error code
	// "SnapshotLimitExceededException".
	//
	// The maximum number of manual snapshots for the directory has been reached.
	// You can use the GetSnapshotLimits operation to determine the snapshot limits
	// for a directory.
	ErrCodeSnapshotLimitExceededException = "SnapshotLimitExceededException"

	// ErrCodeTagLimitExceededException for service response error code
	// "TagLimitExceededException".
	//
	// The maximum allowed number of tags was exceeded.
	ErrCodeTagLimitExceededException = "TagLimitExceededException"

	// ErrCodeUnsupportedOperationException for service response error code
	// "UnsupportedOperationException".
	//
	// The operation is not supported.
	ErrCodeUnsupportedOperationException = "UnsupportedOperationException"

	// ErrCodeUnsupportedSettingsException for service response error code
	// "UnsupportedSettingsException".
	//
	// The specified directory setting is not supported.
	ErrCodeUnsupportedSettingsException = "UnsupportedSettingsException"

	// ErrCodeUserDoesNotExistException for service response error code
	// "UserDoesNotExistException".
	//
	// The user provided a username that does not exist in your directory.
	ErrCodeUserDoesNotExistException = "UserDoesNotExistException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"AccessDeniedException":                  newErrorAccessDeniedException,
	"AuthenticationFailedException":          newErrorAuthenticationFailedException,
	"CertificateAlreadyExistsException":      newErrorCertificateAlreadyExistsException,
	"CertificateDoesNotExistException":       newErrorCertificateDoesNotExistException,
	"CertificateInUseException":              newErrorCertificateInUseException,
	"CertificateLimitExceededException":      newErrorCertificateLimitExceededException,
	"ClientException":                        newErrorClientException,
	"DirectoryAlreadyInRegionException":      newErrorDirectoryAlreadyInRegionException,
	"DirectoryAlreadySharedException":        newErrorDirectoryAlreadySharedException,
	"DirectoryDoesNotExistException":         newErrorDirectoryDoesNotExistException,
	"DirectoryInDesiredStateException":       newErrorDirectoryInDesiredStateException,
	"DirectoryLimitExceededException":        newErrorDirectoryLimitExceededException,
	"DirectoryNotSharedException":            newErrorDirectoryNotSharedException,
	"DirectoryUnavailableException":          newErrorDirectoryUnavailableException,
	"DomainControllerLimitExceededException": newErrorDomainControllerLimitExceededException,
	"EntityAlreadyExistsException":           newErrorEntityAlreadyExistsException,
	"EntityDoesNotExistException":            newErrorEntityDoesNotExistException,
	"IncompatibleSettingsException":          newErrorIncompatibleSettingsException,
	"InsufficientPermissionsException":       newErrorInsufficientPermissionsException,
	"InvalidCertificateException":            newErrorInvalidCertificateException,
	"InvalidClientAuthStatusException":       newErrorInvalidClientAuthStatusException,
	"InvalidLDAPSStatusException":            newErrorInvalidLDAPSStatusException,
	"InvalidNextTokenException":              newErrorInvalidNextTokenException,
	"InvalidParameterException":              newErrorInvalidParameterException,
	"InvalidPasswordException":               newErrorInvalidPasswordException,
	"InvalidTargetException":                 newErrorInvalidTargetException,
	"IpRouteLimitExceededException":          newErrorIpRouteLimitExceededException,
	"NoAvailableCertificateException":        newErrorNoAvailableCertificateException,
	"OrganizationsException":                 newErrorOrganizationsException,
	"RegionLimitExceededException":           newErrorRegionLimitExceededException,
	"ServiceException":                       newErrorServiceException,
	"ShareLimitExceededException":            newErrorShareLimitExceededException,
	"SnapshotLimitExceededException":         newErrorSnapshotLimitExceededException,
	"TagLimitExceededException":              newErrorTagLimitExceededException,
	"UnsupportedOperationException":          newErrorUnsupportedOperationException,
	"UnsupportedSettingsException":           newErrorUnsupportedSettingsException,
	"UserDoesNotExistException":              newErrorUserDoesNotExistException,
}
