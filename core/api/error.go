package api

import (
	"fmt"
)

const (
	// ErrorCodeAPICallFailure error code for API call failure
	ErrorCodeAPICallFailure = "Key_APICallFailure"
	// ErrorCodeDatabaseFailure error code for database failure
	ErrorCodeDatabaseFailure = "Key_DBQueryFailure"
	// ErrorCodeDuplicateValue error code for duplicate value
	ErrorCodeDuplicateValue = "Key_AlreadyExists"
	// ErrorCodeEmptyRequestBody error code for empty request body
	ErrorCodeEmptyRequestBody = "Key_EmptyRequestBody"
	// ErrorCodeHTTPCreateRequestFailure error code for http request creation failure
	ErrorCodeHTTPCreateRequestFailure = "Key_HTTPCreateRequestFailure"
	// ErrorCodeInvalidFormData error code for form parsing error
	ErrorCodeInvalidFormData = "Key_InvalidFormData"
	// ErrorCodeInternalError error code for internal error
	ErrorCodeInternalError = "Key_InternalError"
	// ErrorCodeInvalidFields error code for invalid fields
	ErrorCodeInvalidFields = "Key_InvalidFields"
	// ErrorCodeInvalidJSON error code for invalid JSON
	ErrorCodeInvalidJSON = "Key_InvalidJSON"
	// ErrorCodeInvalidRequestPayload error code for invalid request payload
	ErrorCodeInvalidRequestPayload = "Key_InvalidRequestPayload"
	// ErrorCodeResourceNotFound error code for invalid request payload
	ErrorCodeResourceNotFound = "Key_ResourceNotFound"
	// ErrorCodeUnexpected error code for invalid request payload
	ErrorCodeUnexpected = "Key_Unexpected"
)

// NewHTTPResourceNotFound creates an new instance of HTTP Error
func NewHTTPResourceNotFound(resourceName,
	resourceValue, errorMessage string) HTTPResourceNotFound {
	return HTTPResourceNotFound{
		ErrorCodeResourceNotFound,
		resourceName,
		resourceValue,
		errorMessage,
	}
}

// HTTPResourceNotFound represents HTTP 404 error
type HTTPResourceNotFound struct {
	ErrorKey      string `json:"error"`
	ResourceName  string `json:"resource_name"`
	ResourceValue string `json:"resource_value"`
	ErrorMessage  string `json:"message"`
}

// Error returns the error string
func (e HTTPResourceNotFound) Error() string {
	return e.ErrorKey
}

// NewHTTPError creates an new instance of HTTP Error
func NewHTTPError(err, message string) HTTPError {
	return HTTPError{ErrorKey: err, ErrorMessage: message}
}

// HTTPError Represent an error to be sent back on response
type HTTPError struct {
	ErrorKey     string `json:"error"`
	ErrorMessage string `json:"message"`
}

// Error returns the error string
func (err HTTPError) Error() string {
	return fmt.Sprintf("key : %s message %s", err.ErrorKey, err.ErrorMessage)
}
