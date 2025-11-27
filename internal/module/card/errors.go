package card

import "errors"

var (
	ErrNotFound          = errors.New("not found")
	ErrValidation        = errors.New("validation error")
	ErrInvalidInput      = errors.New("invalid input")
	ErrInvalidDateFormat = errors.New("wrong date format")
	ErrAlreadyExists     = errors.New("resource already exists")
)
