package repos_mysql_tool

import (
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/models"
	params_tool "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/models/params/tool"
)

func (repo *toolRepo) Edit(tool_id int, params *params_tool.ToolEditParams) (*models.Tool, error) {

	err := repo.conn.
		Model(&models.Tool{}).
		Where("id = ?", tool_id).
		Updates(&params).
		Error
	if err != nil {
		return nil, err
	}

	return repo.GetByID(tool_id)
}
