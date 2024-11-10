package logic_errors

import "fmt"

var (
	ErrDeleteNonEmptyQuantity = fmt.Errorf("cannot delete non-empty quantity")
)
