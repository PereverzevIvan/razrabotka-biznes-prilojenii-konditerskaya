package repos_mysql_tool_type

import "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/models"

func (repo *toolTypeRepo) GetAll() ([]models.ToolType, error) {
	var tool_types []models.ToolType

	err := repo.conn.
		Find(&tool_types).
		Error
	if err != nil {
		return nil, err
	}

	return tool_types, nil
}
