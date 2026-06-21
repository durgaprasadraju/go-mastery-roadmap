// Package apperrors provides structured error handling with wrapping.
package apperrors

import (
	"errors"
	"fmt"
)

// Code represents a machine-readable error category.
type Code string

const (
	CodeNotFound     Code = "NOT_FOUND"
	CodeValidation   Code = "VALIDATION"
	CodeUnauthorized Code = "UNAUTHORIZED"
	CodeInternal     Code = "INTERNAL"
	CodeConflict     Code = "CONFLICT"
)

// AppError is a production-grade application error.
type AppError struct {
	Code    Code
	Message string
	Err     error
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %s: %v", e.Code, e.Message, e.Err)
	}
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

func (e *AppError) Unwrap() error { return e.Err }

// New creates an AppError.
func New(code Code, message string) *AppError {
	return &AppError{Code: code, Message: message}
}

// Wrap wraps an existing error with context.
func Wrap(err error, code Code, message string) *AppError {
	return &AppError{Code: code, Message: message, Err: err}
}

// IsCode checks if an error chain contains a specific code.
func IsCode(err error, code Code) bool {
	var ae *AppError
	if errors.As(err, &ae) {
		return ae.Code == code
	}
	return false
}
