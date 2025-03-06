package tool_usecase

import "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/models"

func (u *ToolUsecase) GetByID(tool_id int) (*models.Tool, error) {
	return u.toolRepo.GetByID(tool_id)
}
