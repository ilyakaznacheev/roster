package database

import "errors"

// NotFoundError not found in database
type NotFoundError struct {
	Err  error
	Text string
}

func NewNotFoundError(err error) *NotFoundError {
	return &NotFoundError{
		Err:  err,
		Text: "not found in database",
	}
}

func (e *NotFoundError) Unwrap() error {
	return e.Err
}

func (e *NotFoundError) Error() string {
	return e.Text
}

// ErrNotFound not found
var ErrNotFound *NotFoundError

// ErrExists data already exists
var ErrExists = errors.New("already exists")
