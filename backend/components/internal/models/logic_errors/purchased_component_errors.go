package logic_errors

import "errors"

var (
	ErrDeleteNonEmptyQuantity               = errors.New("cannot delete non-empty quantity")
	ErrNotEnoughPurchasedComponentsToDeduct = errors.New("not enough purchased components to deduct")
)
