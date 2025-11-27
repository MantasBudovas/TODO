package card

import "errors"

var (
	ErrAlreadyExists   = errors.New("card already exists")
	ErrInvalidPriority = errors.New("invalid priority: must be Low, Medium, or High")
)
