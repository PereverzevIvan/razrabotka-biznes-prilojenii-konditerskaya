package component_type_repo_mysql

import "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models"

func (r *ComponentTypeRepo) GetAll(component_category_id int) ([]models.ComponentType, error) {
	var component_types []models.ComponentType

	err := r.db.
		Where("component_category_id = ?", component_category_id).
		Find(&component_types).
		Error
	if err != nil {
		return nil, err
	}

	return component_types, nil
}
