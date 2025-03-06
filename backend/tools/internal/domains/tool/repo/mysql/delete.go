package tool_repo_mysql

import "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/models"

func (r *ToolRepo) Delete(tool_id int) error {

	err := r.db.
		Where("id = ?", tool_id).
		Delete(&models.Tool{}).
		Error
	if err != nil {
		return err
	}

	return nil
}
