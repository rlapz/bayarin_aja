package my_errors

import "errors"

var (
	ErrNoData         = errors.New("no data")
	ErrInvalid        = errors.New("invalid input")
	ErrUnauthorize    = errors.New("unauthorize")
	ErrInternal       = errors.New("internal server error")
	ErrDuplicateEntry = errors.New("duplicate entry")
)
