package errors

import "errors"

var (
	ErrNoData         = errors.New("no data")
	ErrInternal       = errors.New("internal server error")
	ErrDuplicateEntry = errors.New("duplicate entry")
)
