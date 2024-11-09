package services_tool

import "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/models"

func (service *toolService) GetByID(tool_id int) (*models.Tool, error) {
	return service.toolRepo.GetByID(tool_id)
}
