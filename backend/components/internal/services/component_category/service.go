package services_component_category

import (
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/services"
)

type componentCategoryService struct {
	componentCategoryRepo     services.IComponentCategoryRepo
	componentCategoryMapCache map[string]models.ComponentCategory
}

func NewComponentCategoryService(repo services.IComponentCategoryRepo) *componentCategoryService {
	service := &componentCategoryService{
		componentCategoryRepo:     repo,
		componentCategoryMapCache: make(map[string]models.ComponentCategory),
	}

	service.GetAll()
	return service
}
