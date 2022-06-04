package repository

import (
	"fmt"
	"runtime"
	"strconv"

	"gorm.io/gorm"

	"deck/core/api"
)

// NewError creates a new database error
func NewError(err error) Error {
	if err == nil {
		return nil
	}

	return &errorImpl{createUnexpectedErrorImpl(api.ErrorCodeDatabaseFailure, err)}
}

// Error represents an database query failure error interface
type Error interface {
	UnexpectedError
	IsRecordNotFoundError() bool
}
type errorImpl struct {
	unexpectedErrorImpl
}

func (e *errorImpl) IsRecordNotFoundError() bool {
	return gorm.ErrRecordNotFound == e.cause
}

// IsUnexpectedError returns whether the given error is an UnexpectedError
func IsUnexpectedError(err error) bool {
	_, ok := err.(UnexpectedError)

	return ok
}

// NewUnexpectedError creates a new unexpected error
func NewUnexpectedError(errCode string, err error) UnexpectedError {
	if err == nil {
		return nil
	}

	return createUnexpectedErrorImpl(errCode, err)
}

// NewReadWriteError creates a new read write error
func NewReadWriteError(err error) UnexpectedError {
	return NewUnexpectedError(api.ErrorCodeUnexpected, err)
}

// UnexpectedError represents an unexpected error interface
type UnexpectedError interface {
	Error() string
	GetErrorCode() string
	GetStackTrace() string
	GetCause() error
}
type unexpectedErrorImpl struct {
	errCode    string
	cause      error
	stackTrace string
}

// Error returns the error string
func (e unexpectedErrorImpl) Error() string {
	return fmt.Sprintf("%v:%v", e.errCode, e.cause)
}

// GetCause GetErrCode returns the error code
func (e unexpectedErrorImpl) GetCause() error {
	return e.cause
}

// GetErrorCode GetErrCode returns the error code
func (e unexpectedErrorImpl) GetErrorCode() string {
	return e.errCode
}

// GetStackTrace returns the error stack trace
func (e unexpectedErrorImpl) GetStackTrace() string {
	return e.stackTrace
}

func createUnexpectedErrorImpl(errCode string, err error) unexpectedErrorImpl {
	const depth = 20

	var ptrs [depth]uintptr
	n := runtime.Callers(2, ptrs[:])
	ptrSlice := ptrs[0:n]
	stack := ""

	for _, pc := range ptrSlice {
		stackFunc := runtime.FuncForPC(pc)
		_, line := stackFunc.FileLine(pc)
		stack = stack + stackFunc.Name() + ":" + strconv.Itoa(line) + "\n"
	}

	return unexpectedErrorImpl{errCode: errCode, cause: err, stackTrace: stack}
}
