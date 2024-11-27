package logic_errors

import (
	"errors"
)

var (
	ErrCycleDetectedInProductRecipe     error = errors.New("cycle detected in product recipe")
	ErrNotEnoughComponentsToMakeProduct error = errors.New("not enough components to make product")
	ErrNoRecipeOperationsInProduct      error = errors.New("no recipe operations in product")
	ErrNoAvailableTools                 error = errors.New("no available tools")
)
