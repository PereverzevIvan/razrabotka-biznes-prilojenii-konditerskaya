package tool_failure_usecase

import (
	tool_failure_params "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/domains/tool_failure/params"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/models"
)

func (service *ToolFailureUsecase) Create(params *tool_failure_params.CreateParams) (*models.ToolFailure, error) {
	return service.toolFailureRepo.Create(params)
}
