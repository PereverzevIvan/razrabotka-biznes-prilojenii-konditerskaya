package component_type_usecase

import "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models"

func (u *ComponentTypeUsecase) GetAll(component_category_id int) ([]models.ComponentType, error) {
	return u.componentTypeRepo.GetAll(component_category_id)
}
