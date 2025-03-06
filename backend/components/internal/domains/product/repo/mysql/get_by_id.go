package product_repo_mysql

import (
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models"
)

func (r *ProductRepo) GetByID(id int) (*models.Product, error) {
	var product models.Product

	product.RecipeComponents = make([]*models.RecipeComponents, 0)
	product.RecipeSemiProducts = make([]*models.RecipeSemiproduct, 0)

	err := r.db.
		Debug().
		Model(&models.Product{}).
		Preload("RecipeComponents").
		Preload("RecipeSemiProducts").
		Preload("RecipeOperations").
		Preload("RecipeOperations.ToolType").
		Where("id = ?", id).
		First(&product).
		Error
	if err != nil {
		return nil, err
	}

	return &product, nil
}
