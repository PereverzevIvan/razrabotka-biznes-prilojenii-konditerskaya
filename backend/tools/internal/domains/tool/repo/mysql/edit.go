package tool_repo_mysql

import (
	tool_params "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/domains/tool/params"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/models"
)

func (r *ToolRepo) Edit(tool_id int, params *tool_params.ToolEditParams) (*models.Tool, error) {

	err := r.db.
		Model(&models.Tool{}).
		Where("id = ?", tool_id).
		Updates(&params).
		Error
	if err != nil {
		return nil, err
	}

	return r.GetByID(tool_id)
}
