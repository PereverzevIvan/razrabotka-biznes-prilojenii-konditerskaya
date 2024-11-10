package services_component_type

import "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models"

func (service *componentTypeService) GetAll(component_category_id int) ([]models.ComponentType, error) {
	return service.componentTypeRepo.GetAll(component_category_id)
}
