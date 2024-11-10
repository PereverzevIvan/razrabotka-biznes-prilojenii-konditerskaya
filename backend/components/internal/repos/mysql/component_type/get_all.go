package repos_mysql_component_type

import "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models"

func (repo *componentTypeRepo) GetAll(component_category_id int) ([]models.ComponentType, error) {
	var component_types []models.ComponentType

	err := repo.conn.
		Where("component_category_id = ?", component_category_id).
		Find(&component_types).
		Error
	if err != nil {
		return nil, err
	}

	return component_types, nil
}
