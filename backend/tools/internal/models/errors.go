package models

type Error string

func (e Error) Error() string {
	return string(e)
}

const (
	ErrNotFound = Error("not found error")
	ErrUnique   = Error("unique constraint error")
	ErrFK       = Error("foreign key error")
)
