package services_product

import (
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models"
	params_product "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models/params/product"
)

func (service *productService) Create(params *params_product.CreateParams) (*models.Product, error) {
	product := &models.Product{
		Name:          params.Name,
		IsSemiproduct: params.IsSemiproduct,
		Sizes:         params.Sizes,
	}

	for _, component := range params.RecipeComponents {
		product.RecipeComponents = append(product.RecipeComponents, &models.RecipeComponents{
			ComponentID: component.ComponentID,
			Count:       component.Count,
		})
	}

	for _, semiproduct := range params.RecipeSemiProducts {
		product.RecipeSemiProducts = append(product.RecipeSemiProducts, &models.RecipeSemiproduct{
			SemiproductID: semiproduct.SemiproductID,
			Count:         semiproduct.Count,
		})
	}

	for i, operation := range params.RecipeOperations {
		product.RecipeOperations = append(product.RecipeOperations, &models.RecipeOperation{
			ToolTypeID:      operation.ToolTypeID,
			Desctiption:     operation.Desctiption,
			DurationMinutes: operation.DurationMinutes,
			OrderIdx:        i + 1,
		})
	}

	return service.productRepo.Create(product)
}