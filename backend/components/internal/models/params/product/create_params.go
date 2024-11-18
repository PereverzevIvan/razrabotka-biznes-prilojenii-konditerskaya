package params_product

import "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models"

type CreateParams struct {
	Name          string `json:"name"`
	IsSemiproduct bool   `json:"is_semiproduct"`
	Sizes         string `json:"sizes"`

	RecipeSemiProducts []models.RecipeSemiproduct `json:"semiproducts"`
	RecipeComponents   []models.RecipeComponents  `json:"recipe_components"`
	RecipeOperations   []models.RecipeOperation   `json:"recipe_operations"`
}

func (p *CreateParams) Validate() []string {
	errs := make([]string, 0)

	if p.Name == "" {
		errs = append(errs, "name is required")
	}

	return errs
}
