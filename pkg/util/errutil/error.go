package errutil

import (
	"fmt"
)

// Common Errors forming the base of our error system
//
// Many Errors returned by Gitea can be tested against these errors
// using errors.Is.
var (
	ErrArgumentRequired = Error("100000", "argument required")
	ErrInvalidArgument  = Error("100001", "invalid argument")
	ErrPermissionDenied = Error("100002", "permission denied")
	ErrAlreadyExist     = Error("100003", "resource already exists")
	ErrNotExist         = Error("100004", "resource does not exist")
)

// GenericError
type GenericError struct {
	Code    string
	Message string
}

var _ error = &GenericError{}

// 实现 error 接口的 Error 方法
func (e *GenericError) Error() string {
	return fmt.Sprintf("Error Code: %s, Message: %s", e.Code, e.Message)
}

// func WrapErrorf(unwrap error, code string, text string, args ...any) error {
// 	if len(args) == 0 {
// 		return GenericError{Code: code, Message: message, Err: unwrap}
// 	}
// 	return SilentWrap{Message: fmt.Sprintf(message, args...), Err: unwrap}
// }

func Error(code string, text string) error {
	return Errorf(code, text)
}

// Create a new error
func Errorf(code string, text string, args ...any) error {
	if len(args) == 0 {
		return &GenericError{Code: code, Message: text}
	}
	return &GenericError{Code: code, Message: fmt.Sprintf(text, args...)}
}

// SilentWrap provides a simple wrapper for a wrapped error where the wrapped error message plays no part in the error message
// Especially useful for "untyped" errors created with "errors.New(…)" that can be classified as 'invalid argument', 'permission denied', 'exists already', or 'does not exist'
type SilentWrap struct {
	Message string
	Err     error
}

// Error returns the message
func (w SilentWrap) Error() string {
	return w.Message
}

// Unwrap returns the underlying error
func (w SilentWrap) Unwrap() error {
	return w.Err
}

// NewSilentWrapErrorf returns an error that formats as the given text but unwraps as the provided error
func NewSilentWrapErrorf(unwrap error, message string, args ...any) error {
	if len(args) == 0 {
		return SilentWrap{Message: message, Err: unwrap}
	}
	return SilentWrap{Message: fmt.Sprintf(message, args...), Err: unwrap}
}

// NewInvalidArgumentErrorf returns an error that formats as the given text but unwraps as an ErrInvalidArgument
func NewInvalidArgumentErrorf(message string, args ...any) error {
	return NewSilentWrapErrorf(ErrInvalidArgument, message, args...)
}

// NewPermissionDeniedErrorf returns an error that formats as the given text but unwraps as an ErrPermissionDenied
func NewPermissionDeniedErrorf(message string, args ...any) error {
	return NewSilentWrapErrorf(ErrPermissionDenied, message, args...)
}

// NewAlreadyExistErrorf returns an error that formats as the given text but unwraps as an ErrAlreadyExist
func NewAlreadyExistErrorf(message string, args ...any) error {
	return NewSilentWrapErrorf(ErrAlreadyExist, message, args...)
}

// NewNotExistErrorf returns an error that formats as the given text but unwraps as an ErrNotExist
func NewNotExistErrorf(message string, args ...any) error {
	return NewSilentWrapErrorf(ErrNotExist, message, args...)
}
