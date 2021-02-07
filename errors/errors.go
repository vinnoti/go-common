package errors

import (
	"fmt"
)

type DataNotFoundError struct {
	Entity string
	Key    string
	Err    error
}

func (e *DataNotFoundError) Error() string {
	return fmt.Sprintf("data %s not found with key %s: %s", e.Entity, e.Key, e.Err)
}

type SaveError struct {
	Entity string
	Err    error
}

func (e *SaveError) Error() string {
	return fmt.Sprintf("failed to save %s: %s", e.Entity, e.Err)
}

type ValidationError struct {
	Entity string
	Err    error
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("unable to validate %s: %s", e.Entity, e.Err)
}

type InvalidCredentialsError struct{}

func (e *InvalidCredentialsError) Error() string {
	return fmt.Sprintf("invalid credentials. Please re-check.")
}

type ClassCodeNotAvailableError struct {
	Code string
}

func (e *ClassCodeNotAvailableError) Error() string {
	return fmt.Sprintf("class code %s is not available.", e.Code)
}

type UnauthorizedError struct {
	Action string
	Entity string
}

func (e *UnauthorizedError) Error() string {
	return fmt.Sprintf("you don't have permission to %s %s.", e.Action, e.Entity)
}

type InvalidClassCode struct{}

func (e *InvalidClassCode) Error() string {
	return fmt.Sprintf("class code should be 6 character alphanumeric.")
}

type EmailRegistered struct {
	Email string
}

func (e *EmailRegistered) Error() string {
	return fmt.Sprintf("email %s already registered.", e.Email)
}

type EmailAlreadyVerified struct {
	Email string
}

func (e *EmailAlreadyVerified) Error() string {
	return fmt.Sprintf("email %s already verified.", e.Email)
}
