package tool_failure_repo_mysql

import "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/models"

func (repo *ToolFailureRepo) GetAll() ([]models.ToolFailure, error) {
	var tool_failures []models.ToolFailure

	err := repo.conn.
		Find(&tool_failures).
		Error
	if err != nil {
		return nil, err
	}

	return tool_failures, nil
}
