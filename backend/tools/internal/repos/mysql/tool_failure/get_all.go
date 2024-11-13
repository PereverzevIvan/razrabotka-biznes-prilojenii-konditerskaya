package repos_mysql_tool_failure

import "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/models"

func (repo *toolFailureRepo) GetAll() ([]models.ToolFailure, error) {
	var tool_failures []models.ToolFailure

	err := repo.conn.
		Find(&tool_failures).
		Error
	if err != nil {
		return nil, err
	}

	return tool_failures, nil
}
