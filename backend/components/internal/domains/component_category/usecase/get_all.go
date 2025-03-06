package component_category_usecase

import "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models"

func (u *ComponentCategoryUsecase) GetAll() ([]models.ComponentCategory, error) {
	component_categories, err := u.componentCategoryRepo.GetAll()
	if err != nil {
		return nil, err
	}

	for _, component_category := range component_categories {
		if cached_component_category, ok := u.componentCategoryMapCache[component_category.Name]; !ok || cached_component_category.ID != component_category.ID {
			u.componentCategoryMapCache[component_category.Name] = component_category
		}
	}

	return component_categories, nil
}
