package tool_type_repo_mysql

import "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/models"

func (r *ToolTypeRepo) GetAll() ([]models.ToolType, error) {
	var tool_types []models.ToolType

	err := r.db.
		Find(&tool_types).
		Error
	if err != nil {
		return nil, err
	}

	return tool_types, nil
}
