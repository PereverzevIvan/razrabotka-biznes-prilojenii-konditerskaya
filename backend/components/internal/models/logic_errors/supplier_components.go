package logic_errors

import "errors"

var (
	ErrNoSupplierForComponent error = errors.New("no supplier for component")
)
