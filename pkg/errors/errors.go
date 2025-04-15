package errors

import (
	"errors"
	"fmt"
)

// common errors
var (
	ErrInvalidArgument      = New("invalid argument")
	ErrInvalidType          = New("invalid type")
	ErrInvalidValue         = New("invalid value")
	ErrInvalidConfig        = New("invalid config")
	ErrInvalidFormat        = New("invalid format")
	ErrInvalidState         = New("invalid state")
	ErrInvalidNoLogger      = New("invalid no logger")
	ErrInvalidStatus        = New("invalid status")
	ErrInvalidFileStatus    = New("invalid file not exist")
	ErrInvalidStatusCode    = New("invalid status code")
	ErrInvalidStatusMessage = New("invalid status message")
	ErrInvalidStatusReason  = New("invalid status reason")
	ErrInvalidStatusDetails = New("invalid status details")
)

// New returns a new error with the given message.
func New(text string) error {
	return errors.New(text)
}

func ErrorNew(format string, a ...any) error {
	return fmt.Errorf(format, a...)
}

// Specify returns a new error with the given message and cause.
