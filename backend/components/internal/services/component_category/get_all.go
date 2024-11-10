package services_component_category

import "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models"

func (service *componentCategoryService) GetAll() ([]models.ComponentCategory, error) {
	component_categories, err := service.componentCategoryRepo.GetAll()
	if err != nil {
		return nil, err
	}

	for _, component_category := range component_categories {
		if cached_component_category, ok := service.componentCategoryMapCache[component_category.Name]; !ok || cached_component_category.ID != component_category.ID {
			service.componentCategoryMapCache[component_category.Name] = component_category
		}
	}

	return component_categories, nil
}
