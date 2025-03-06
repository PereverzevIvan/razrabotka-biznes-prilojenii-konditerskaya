package component_category_usecase

import "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models"

func (u *ComponentCategoryUsecase) GetByName(name string) (*models.ComponentCategory, error) {
	// WARNING: check code below on errors
	if component_category, ok := u.componentCategoryMapCache[name]; ok {
		return &component_category, nil
	}

	component_category, err := u.componentCategoryRepo.GetByName(name)
	if err != nil {
		return nil, err
	}
	if component_category == nil {
		return nil, models.ErrNotFound
	}

	u.componentCategoryMapCache[name] = *component_category

	return component_category, nil
}
