package repos_mysql_component

import "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models"

func (r *componentRepo) GetByID(id int) (*models.Component, error) {
	var component models.Component

	err := r.conn.
		Model(&models.Component{}).
		Where("id = ?", id).
		First(&component).
		Error
	if err != nil {
		return nil, err
	}

	return &component, nil
}
