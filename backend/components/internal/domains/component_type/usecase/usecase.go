package component_type_usecase

import (
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models"
)

type IComponentTypeRepo interface {
	GetAll(component_category_id int) ([]models.ComponentType, error)
}

type ComponentTypeUsecase struct {
	componentTypeRepo IComponentTypeRepo
}

func NewComponentTypeUsecase(componentTypeRepo IComponentTypeRepo) *ComponentTypeUsecase {
	return &ComponentTypeUsecase{
		componentTypeRepo: componentTypeRepo,
	}
}
