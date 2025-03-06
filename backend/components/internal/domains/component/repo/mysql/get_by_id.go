package component_repo_mysql

import "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models"

func (r *ComponentRepo) GetByID(id int) (*models.Component, error) {
	var component models.Component

	err := r.db.
		Model(&models.Component{}).
		Where("id = ?", id).
		First(&component).
		Error
	if err != nil {
		return nil, err
	}

	return &component, nil
}
