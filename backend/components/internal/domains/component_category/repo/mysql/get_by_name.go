package component_category_repo_mysql

import "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models"

func (r *ComponentCategoryRepo) GetByName(name string) (*models.ComponentCategory, error) {
	var component_category models.ComponentCategory

	err := r.db.Where("name = ?", name).First(&component_category).Error
	if err != nil {
		return nil, err
	}

	return &component_category, nil
}
