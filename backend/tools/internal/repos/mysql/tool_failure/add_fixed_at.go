package repos_mysql_tool_failure

import (
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/models"
	params_tool_failure "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/models/params/tool_failure"
)

func (repo *toolFailureRepo) AddFixedAt(params *params_tool_failure.AddFixedAtParams) error {

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
