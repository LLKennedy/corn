package corn

import "fmt"

const errorStart = "Corn CLI: "

// Error is a corn error
type Error struct {
	Message string
	Err     error
}

// ErrorMessage creates a new error with just a message
func ErrorMessage(message string) *Error {
	return &Error{
		Message: message,
	}
}

// ErrorError creates a new error with just an error
func ErrorError(err error) *Error {
	return &Error{
		Err: err,
	}
}

// ErrorBoth creates a new error with both a message and an error
func ErrorBoth(message string, err error) *Error {
	return &Error{
		Message: message,
		Err:     err,
	}
}

// Error returns the error string
func (e *Error) Error() string {
	errString := ""
	switch {
	case e == nil:
		// Totally nil error
		errString = errorStart + "nil error"
	case e.Message == "" && e.Err == nil:
		// Nothing set
		errString = errorStart + "empty error"
	case e.Err == nil:
		// Only message
		errString = errorStart + e.Message
	case e.Message == "":
		// Only error
		errString = errorStart + e.Err.Error()
	default:
		// Both are set
		errString = errorStart + fmt.Sprintf("Message: %s, Error: %v", e.Message, e.Err)
	}
	return errString
}

// Unwrap returns the wrapped error
func (e *Error) Unwrap() error {
	if e == nil {
		return nil
	}
	return e.Err
}
