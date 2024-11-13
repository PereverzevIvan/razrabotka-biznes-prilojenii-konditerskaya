package repos_mysql_tool_failure

import "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/models"

func (repo *toolFailureRepo) GetAllReasons() ([]models.ToolFailureReason, error) {
	var tool_failure_reasons []models.ToolFailureReason

	err := repo.conn.
		Find(&tool_failure_reasons).
		Error
	if err != nil {
		return nil, err
	}

	return tool_failure_reasons, nil
}
