package tool_failure_repo_mysql

import "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/models"

func (repo *ToolFailureRepo) GetAllReasons() ([]models.ToolFailureReason, error) {
	var tool_failure_reasons []models.ToolFailureReason

	err := repo.conn.
		Find(&tool_failure_reasons).
		Error
	if err != nil {
		return nil, err
	}

	return tool_failure_reasons, nil
}
