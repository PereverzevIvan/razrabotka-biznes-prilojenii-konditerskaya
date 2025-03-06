package tool_repo_mysql

import "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/models"

func (r *ToolRepo) GetByID(tool_id int) (*models.Tool, error) {
	var tool *models.Tool

	err := r.db.
		Joins("ToolType").
		Joins("Supplier").
		Joins("DegreeOfWear").
		Where("id = ?", tool_id).
		First(&tool).
		Error

	if err != nil {
		return nil, err
	}

	return tool, nil
}
