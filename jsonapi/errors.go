package jsonapi

import "fmt"

// ErrorType for diff types of errors
type ErrorType string

// List of error types
const (
	UserErrorType ErrorType = "user input error"
	HttpErrorType ErrorType = "http error"
)

// Error custom error message
type Error struct {
	// Error Type
	Type ErrorType

	// HTTP Response Status Code
	Code int

	// Custom Response Error message
	Message string
}

func (e *Error) Error() string {
	if e.Type == HttpErrorType {
		return fmt.Sprintf("%v:: error code: %v error message: %v", HttpErrorType, e.Code, e.Message)
	}
	if e.Type == UserErrorType {
		return fmt.Sprintf("%v:: error message: %v", UserErrorType, e.Message)
	}
	return fmt.Sprintf("%v", e.Message)
}

// Is func is used for comparing two errors
func (e *Error) Is(target error) bool {
	t, ok := target.(*Error)
	if !ok {
		return false
	}
	return (e.Code == t.Code && e.Message == t.Message && e.Type == t.Type)
}
