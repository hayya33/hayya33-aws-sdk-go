// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package workmail

import (
	"github.com/aws/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeDirectoryInUseException for service response error code
	// "DirectoryInUseException".
	//
	// The directory is already in use by another WorkMail organization in the same
	// account and Region.
	ErrCodeDirectoryInUseException = "DirectoryInUseException"

	// ErrCodeDirectoryServiceAuthenticationFailedException for service response error code
	// "DirectoryServiceAuthenticationFailedException".
	//
	// The directory service doesn't recognize the credentials supplied by WorkMail.
	ErrCodeDirectoryServiceAuthenticationFailedException = "DirectoryServiceAuthenticationFailedException"

	// ErrCodeDirectoryUnavailableException for service response error code
	// "DirectoryUnavailableException".
	//
	// The directory is unavailable. It might be located in another Region or deleted.
	ErrCodeDirectoryUnavailableException = "DirectoryUnavailableException"

	// ErrCodeEmailAddressInUseException for service response error code
	// "EmailAddressInUseException".
	//
	// The email address that you're trying to assign is already created for a different
	// user, group, or resource.
	ErrCodeEmailAddressInUseException = "EmailAddressInUseException"

	// ErrCodeEntityAlreadyRegisteredException for service response error code
	// "EntityAlreadyRegisteredException".
	//
	// The user, group, or resource that you're trying to register is already registered.
	ErrCodeEntityAlreadyRegisteredException = "EntityAlreadyRegisteredException"

	// ErrCodeEntityNotFoundException for service response error code
	// "EntityNotFoundException".
	//
	// The identifier supplied for the user, group, or resource does not exist in
	// your organization.
	ErrCodeEntityNotFoundException = "EntityNotFoundException"

	// ErrCodeEntityStateException for service response error code
	// "EntityStateException".
	//
	// You are performing an operation on a user, group, or resource that isn't
	// in the expected state, such as trying to delete an active user.
	ErrCodeEntityStateException = "EntityStateException"

	// ErrCodeInvalidConfigurationException for service response error code
	// "InvalidConfigurationException".
	//
	// The configuration for a resource isn't valid. A resource must either be able
	// to auto-respond to requests or have at least one delegate associated that
	// can do so on its behalf.
	ErrCodeInvalidConfigurationException = "InvalidConfigurationException"

	// ErrCodeInvalidParameterException for service response error code
	// "InvalidParameterException".
	//
	// One or more of the input parameters don't match the service's restrictions.
	ErrCodeInvalidParameterException = "InvalidParameterException"

	// ErrCodeInvalidPasswordException for service response error code
	// "InvalidPasswordException".
	//
	// The supplied password doesn't match the minimum security constraints, such
	// as length or use of special characters.
	ErrCodeInvalidPasswordException = "InvalidPasswordException"

	// ErrCodeLimitExceededException for service response error code
	// "LimitExceededException".
	//
	// The request exceeds the limit of the resource.
	ErrCodeLimitExceededException = "LimitExceededException"

	// ErrCodeMailDomainNotFoundException for service response error code
	// "MailDomainNotFoundException".
	//
	// For an email or alias to be created in Amazon WorkMail, the included domain
	// must be defined in the organization.
	ErrCodeMailDomainNotFoundException = "MailDomainNotFoundException"

	// ErrCodeMailDomainStateException for service response error code
	// "MailDomainStateException".
	//
	// After a domain has been added to the organization, it must be verified. The
	// domain is not yet verified.
	ErrCodeMailDomainStateException = "MailDomainStateException"

	// ErrCodeNameAvailabilityException for service response error code
	// "NameAvailabilityException".
	//
	// The user, group, or resource name isn't unique in Amazon WorkMail.
	ErrCodeNameAvailabilityException = "NameAvailabilityException"

	// ErrCodeOrganizationNotFoundException for service response error code
	// "OrganizationNotFoundException".
	//
	// An operation received a valid organization identifier that either doesn't
	// belong or exist in the system.
	ErrCodeOrganizationNotFoundException = "OrganizationNotFoundException"

	// ErrCodeOrganizationStateException for service response error code
	// "OrganizationStateException".
	//
	// The organization must have a valid state to perform certain operations on
	// the organization or its members.
	ErrCodeOrganizationStateException = "OrganizationStateException"

	// ErrCodeReservedNameException for service response error code
	// "ReservedNameException".
	//
	// This user, group, or resource name is not allowed in Amazon WorkMail.
	ErrCodeReservedNameException = "ReservedNameException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// The resource cannot be found.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"

	// ErrCodeTooManyTagsException for service response error code
	// "TooManyTagsException".
	//
	// The resource can have up to 50 user-applied tags.
	ErrCodeTooManyTagsException = "TooManyTagsException"

	// ErrCodeUnsupportedOperationException for service response error code
	// "UnsupportedOperationException".
	//
	// You can't perform a write operation against a read-only directory.
	ErrCodeUnsupportedOperationException = "UnsupportedOperationException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"DirectoryInUseException":                       newErrorDirectoryInUseException,
	"DirectoryServiceAuthenticationFailedException": newErrorDirectoryServiceAuthenticationFailedException,
	"DirectoryUnavailableException":                 newErrorDirectoryUnavailableException,
	"EmailAddressInUseException":                    newErrorEmailAddressInUseException,
	"EntityAlreadyRegisteredException":              newErrorEntityAlreadyRegisteredException,
	"EntityNotFoundException":                       newErrorEntityNotFoundException,
	"EntityStateException":                          newErrorEntityStateException,
	"InvalidConfigurationException":                 newErrorInvalidConfigurationException,
	"InvalidParameterException":                     newErrorInvalidParameterException,
	"InvalidPasswordException":                      newErrorInvalidPasswordException,
	"LimitExceededException":                        newErrorLimitExceededException,
	"MailDomainNotFoundException":                   newErrorMailDomainNotFoundException,
	"MailDomainStateException":                      newErrorMailDomainStateException,
	"NameAvailabilityException":                     newErrorNameAvailabilityException,
	"OrganizationNotFoundException":                 newErrorOrganizationNotFoundException,
	"OrganizationStateException":                    newErrorOrganizationStateException,
	"ReservedNameException":                         newErrorReservedNameException,
	"ResourceNotFoundException":                     newErrorResourceNotFoundException,
	"TooManyTagsException":                          newErrorTooManyTagsException,
	"UnsupportedOperationException":                 newErrorUnsupportedOperationException,
}
