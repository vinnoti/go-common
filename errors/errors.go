package errors

import "fmt"

type DataNotFoundError struct {
	Entity string
	Key    string
	Err    error
}

func (e *DataNotFoundError) Error() string {
	return fmt.Sprintf("Data %s not found with key %s: %s", e.Entity, e.Key, e.Err)
}

type SaveError struct {
	Entity string
	Err    error
}

func (e *SaveError) Error() string {
	return fmt.Sprintf("Failed to save %s: %s", e.Entity, e.Err)
}

type ValidationError struct {
	Entity string
	Err    error
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("Unable to validate %s: %s", e.Entity, e.Err)
}

type InvalidCredentialsError struct{}

func (e *InvalidCredentialsError) Error() string {
	return fmt.Sprintf("Invalid credentials. Please re-check.")
}

type ClassCodeNotAvailableError struct {
	Code string
}

func (e *ClassCodeNotAvailableError) Error() string {
	return fmt.Sprintf("Class code %s is not available.", e.Code)
}

type UnauthorizedError struct {
	Action string
	Entity string
}

func (e *UnauthorizedError) Error() string {
	return fmt.Sprintf("You don't have permission to %s %s.", e.Action, e.Entity)
}

type InvalidClassCode struct{}

func (e *InvalidClassCode) Error() string {
	return fmt.Sprintf("Class code should be 6 character alphanumeric.")
}
