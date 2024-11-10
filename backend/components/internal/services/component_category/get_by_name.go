package services_component_category

import "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models"

func (service *componentCategoryService) GetByName(name string) (*models.ComponentCategory, error) {
	// WARNING: check code below on errors
	if component_category, ok := service.componentCategoryMapCache[name]; ok {
		return &component_category, nil
	}

	component_category, err := service.componentCategoryRepo.GetByName(name)
	if err != nil {
		return nil, err
	}
	if component_category == nil {
		return nil, models.ErrNotFound
	}

	service.componentCategoryMapCache[name] = *component_category

	return component_category, nil
}
