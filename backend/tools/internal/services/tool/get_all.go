package services_tool

import (
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/models"
	params_tool "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/models/params/tool"
)

func (service *toolService) GetAll(params *params_tool.ToolGetAllParams) ([]models.Tool, error) {
	return service.toolRepo.GetAll(params)
}
