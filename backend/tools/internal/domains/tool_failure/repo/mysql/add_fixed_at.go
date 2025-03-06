package tool_failure_repo_mysql

import (
	tool_failure_params "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/domains/tool_failure/params"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/models"
)

func (repo *ToolFailureRepo) AddFixedAt(params *tool_failure_params.AddFixedAtParams) error {

	tool_failure := models.ToolFailure{
		ID:      params.ToolFailureID,
		FixedAt: params.FixedAt,
	}

	err := repo.conn.
		Updates(tool_failure).
		Where("id = ?", params.ToolFailureID).
		Error

	if err != nil {
		return err
	}

	return nil
}
