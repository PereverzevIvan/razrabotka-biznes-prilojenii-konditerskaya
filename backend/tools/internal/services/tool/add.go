package services_tool

import (
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/models"
	params_tool "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/models/params/tool"
)

func (service *toolService) Add(params *params_tool.ToolAddParams) (*models.Tool, error) {

	new_tool := &models.Tool{
		ToolTypeID:     params.ToolType,
		SupplierID:     params.SupplierID,
		DegreeOfWearID: params.DegreeOfWearID,
		Name:           params.Name,
		Description:    params.Description,
	}

	return service.toolRepo.Add(new_tool)
}
