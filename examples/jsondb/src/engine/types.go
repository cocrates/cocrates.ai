package engine

import "errors"

// ErrorType represents the type of error that occurred.
type ErrorType int

const (
	// ErrFileNotFound indicates the file does not exist.
	ErrFileNotFound ErrorType = iota
	// ErrInvalidPath indicates an invalid field path.
	ErrInvalidPath
	// ErrFieldNotFound indicates the field does not exist.
	ErrFieldNotFound
	// ErrTypeMismatch indicates a type mismatch.
	ErrTypeMismatch
	// ErrPermissionDenied indicates a permission error.
	ErrPermissionDenied
	// ErrStorageFailure indicates a storage backend failure.
	ErrStorageFailure
)

// JsonDBError represents a structured error from jsondb.
type JsonDBError struct {
	Type    ErrorType
	Message string
	Path    string
	Cause   error
}

func (e *JsonDBError) Error() string {
	if e.Cause != nil {
		return e.Message + ": " + e.Cause.Error()
	}
	return e.Message
}

func (e *JsonDBError) Unwrap() error {
	return e.Cause
}

// Common errors
var (
	ErrNotOpen = errors.New("engine not open")
)
