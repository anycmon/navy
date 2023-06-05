package nerrors

import "errors"

var (
	// ErrInvalidInput is returned when the input is invalid
	ErrInvalidInput = errors.New("invalid input")

	// ErrNotFound is returned when entity not found
	ErrNotFound = errors.New("not found")
)
