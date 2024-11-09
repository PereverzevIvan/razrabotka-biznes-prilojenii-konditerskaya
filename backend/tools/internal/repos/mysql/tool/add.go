package repos_mysql_tool

import "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/models"

func (repo *toolRepo) Add(tool *models.Tool) (*models.Tool, error) {

	err := repo.conn.
		Create(&tool).
		Error

	if err != nil {
		return nil, err
	}

	return tool, nil
}
