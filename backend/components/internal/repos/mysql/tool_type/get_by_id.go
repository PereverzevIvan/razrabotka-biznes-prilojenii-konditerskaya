package repos_mysql_tool_type

import "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models"

func (repo *toolTypeRepo) GetByID(id int) (*models.ToolType, error) {
	var tool_type *models.ToolType

	err := repo.conn.
		Where("id = ?", id).
		First(&tool_type).
		Error
	if err != nil {
		return nil, err
	}

	return tool_type, nil
}
