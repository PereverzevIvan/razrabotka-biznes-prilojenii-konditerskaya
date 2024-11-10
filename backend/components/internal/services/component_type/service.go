package services_component_type

import (
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/services"
)

type componentTypeService struct {
	componentTypeRepo services.IComponentTypeRepo
}

func NewComponentTypeService(componentTypeRepo services.IComponentTypeRepo) *componentTypeService {
	return &componentTypeService{
		componentTypeRepo: componentTypeRepo,
	}
}
