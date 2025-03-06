package tool_repo_mysql

import "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/models"

func (r *ToolRepo) Add(tool *models.Tool) (*models.Tool, error) {

	err := r.db.
		Create(&tool).
		Error

	if err != nil {
		return nil, err
	}

	return tool, nil
}
