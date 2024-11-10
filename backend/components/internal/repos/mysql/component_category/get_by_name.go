package repos_mysql_component_category

import "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models"

func (repo *componentCategoryRepo) GetByName(name string) (*models.ComponentCategory, error) {
	var component_category models.ComponentCategory

	err := repo.conn.Where("name = ?", name).First(&component_category).Error
	if err != nil {
		return nil, err
	}

	return &component_category, nil
}
