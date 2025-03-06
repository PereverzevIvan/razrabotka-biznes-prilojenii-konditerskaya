package component_category_usecase

import (
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models"
)

type IComponentCategoryRepo interface {
	GetAll() ([]models.ComponentCategory, error)
	GetByName(name string) (*models.ComponentCategory, error)
}

type ComponentCategoryUsecase struct {
	componentCategoryRepo     IComponentCategoryRepo
	componentCategoryMapCache map[string]models.ComponentCategory
}

func NewComponentCategoryUsecase(componentCategoryRepo IComponentCategoryRepo) *ComponentCategoryUsecase {
	usecase := &ComponentCategoryUsecase{
		componentCategoryRepo:     componentCategoryRepo,
		componentCategoryMapCache: make(map[string]models.ComponentCategory),
	}

	usecase.GetAll()
	return usecase
}
