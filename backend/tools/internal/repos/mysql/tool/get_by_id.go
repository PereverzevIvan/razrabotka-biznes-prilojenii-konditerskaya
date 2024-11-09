package repos_mysql_tool

import "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/models"

func (repo *toolRepo) GetByID(tool_id int) (*models.Tool, error) {
	var tool *models.Tool

	err := repo.conn.
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
