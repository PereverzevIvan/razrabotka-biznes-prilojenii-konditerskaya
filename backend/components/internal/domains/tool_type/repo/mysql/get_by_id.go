package tool_type_repo_mysql

import "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models"

func (r *ToolTypeRepo) GetByID(id int) (*models.ToolType, error) {
	var tool_type *models.ToolType

	err := r.db.
		Where("id = ?", id).
		First(&tool_type).
		Error
	if err != nil {
		return nil, err
	}

	return tool_type, nil
}
