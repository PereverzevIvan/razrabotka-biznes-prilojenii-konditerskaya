package tool_failure_usecase

import "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/models"

func (service *ToolFailureUsecase) GetAll() ([]models.ToolFailure, error) {
	return service.toolFailureRepo.GetAll()
}
