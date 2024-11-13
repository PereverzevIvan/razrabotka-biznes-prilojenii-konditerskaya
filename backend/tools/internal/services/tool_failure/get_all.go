package services_tool_failure

import "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/models"

func (service *toolFailureService) GetAll() ([]models.ToolFailure, error) {
	return service.toolFailureRepo.GetAll()
}
