package services_tool_type

import "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/models"

func (s *toolTypeService) GetAll() ([]models.ToolType, error) {
	return s.toolTypeRepo.GetAll()
}
