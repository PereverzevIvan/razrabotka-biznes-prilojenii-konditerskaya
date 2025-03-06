package product_params

import "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models"

type UpdateParams struct {
	ID                 int                        `json:"id"`
	Name               string                     `json:"name"`
	IsSemiproduct      bool                       `json:"is_semiproduct"`
	Sizes              string                     `json:"sizes"`
	RecipeSemiProducts []models.RecipeSemiproduct `json:"semiproducts"`
	RecipeComponents   []models.RecipeComponents  `json:"recipe_components"`
	RecipeOperations   []models.RecipeOperation   `json:"recipe_operations"`
}

func (params *UpdateParams) Validate() []string {
	errs := make([]string, 0)

	if params.ID <= 0 {
		errs = append(errs, "id is required")
	}

	return errs
}
