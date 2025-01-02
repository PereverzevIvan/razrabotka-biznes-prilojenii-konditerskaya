package logic_errors

import "errors"

var (
	ErrInvalidUpdate error = errors.New("invalid update")
)
