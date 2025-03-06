package component_category_repo_mysql

import "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models"

func (r *ComponentCategoryRepo) GetAll() ([]models.ComponentCategory, error) {
	var component_categories []models.ComponentCategory

	err := r.db.Find(&component_categories).Error
	if err != nil {
		return nil, err
	}

	return component_categories, nil
}
