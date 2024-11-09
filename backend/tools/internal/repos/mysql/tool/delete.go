package repos_mysql_tool

import "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/models"

func (repo *toolRepo) Delete(tool_id int) error {

	err := repo.conn.
		Where("id = ?", tool_id).
		Delete(&models.Tool{}).
		Error
	if err != nil {
		return err
	}

	return nil
}
