package services_product

import "errors"

var (
	ErrCycleDetectedInProductRecipe error = errors.New("cycle detected in product recipe")
	// ErrProductNotFound = models.ErrProductNotFound

)
