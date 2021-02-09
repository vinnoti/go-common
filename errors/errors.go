package errors

import (
	"errors"
)

var (
	ErrDataNotFound           = errors.New("record not found")
	ErrSave                   = errors.New("failed to save record")
	ErrValidation             = errors.New("can't validate or bad request")
	ErrInvalidCredentials     = errors.New("invalid credentials. please re-check")
	ErrClassNotAvailable      = errors.New("class code not available")
	ErrUnauthorized           = errors.New("you don't have permission to access this resource")
	ErrInvalidClassCode       = errors.New("class code should be 6 character alphanumeric")
	ErrEmailAlreadyRegistered = errors.New("this email address is already registered")
	ErrEmailAlreadyVerified   = errors.New("this email address is already verified")
	ErrUserNotRegistered      = errors.New("this email address is not yet registered")
)
