package logic_errors

import (
	"errors"
)

var (
	ErrCycleDetectedInProductRecipe     error = errors.New("cycle detected in product recipe")
	ErrNotEnoughComponentsToMakeProduct error = errors.New("not enough components to make product")
)
