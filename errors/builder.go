package errors

import "errors"

// New returns an error with the provided message
func New(msg string) error {
	return errors.New(msg)
}
