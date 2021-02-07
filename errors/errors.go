package errors

import "fmt"

type DataNotFoundError struct {
	entity string
	key    string
	err    error
}

func (e *DataNotFoundError) Error() string {
	return fmt.Sprintf("Data %s not found with key %s: %s", e.entity, e.key, e.err)
}

type SaveError struct {
	entity string
	err    error
}

func (e *SaveError) Error() string {
	return fmt.Sprintf("Failed to save %s: %s", e.entity, e.err)
}

type ValidationError struct {
	entity string
	err    error
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("Unable to validate %s: %s", e.entity, e.err)
}

type InvalidCredentialsError struct{}

func (e *InvalidCredentialsError) Error() string {
	return fmt.Sprintf("Invalid credentials. Please re-check.")
}

type ClassCodeNotAvailableError struct {
	code string
}

func (e *ClassCodeNotAvailableError) Error() string {
	return fmt.Sprintf("Class code %s is not available.", e.code)
}

type UnauthorizedError struct {
	action string
	entity string
}

func (e *UnauthorizedError) Error() string {
	return fmt.Sprintf("You don't have permission to %s %s.", e.action, e.entity)
}

type InvalidClassCode struct{}

func (e *InvalidClassCode) Error() string {
	return fmt.Sprintf("Class code should be 6 character alphanumeric.")
}
