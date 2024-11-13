package services_tool_failure

import (
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/models"
	params_tool_failure "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/models/params/tool_failure"
)

func (service *toolFailureService) Create(params *params_tool_failure.CreateParams) (*models.ToolFailure, error) {
	return service.toolFailureRepo.Create(params)
}
