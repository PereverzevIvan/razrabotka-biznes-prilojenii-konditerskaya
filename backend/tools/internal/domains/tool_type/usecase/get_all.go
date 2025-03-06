package tool_type_usecase

import "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/models"

func (u *ToolTypeUsecase) GetAll() ([]models.ToolType, error) {
	return u.toolTypeRepo.GetAll()
}
