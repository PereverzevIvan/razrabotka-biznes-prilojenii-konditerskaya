package repos_mysql_component_category

import "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models"

func (repo *componentCategoryRepo) GetAll() ([]models.ComponentCategory, error) {
	var component_categories []models.ComponentCategory

	err := repo.conn.Find(&component_categories).Error
	if err != nil {
		return nil, err
	}

	return component_categories, nil
}
