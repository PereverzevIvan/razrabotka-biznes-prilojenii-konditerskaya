package tool_usecase

import (
	tool_params "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/domains/tool/params"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/models"
)

func (u *ToolUsecase) Add(params *tool_params.ToolAddParams) (*models.Tool, error) {

	new_tool := &models.Tool{
		ToolTypeID:     params.ToolType,
		SupplierID:     params.SupplierID,
		DegreeOfWearID: params.DegreeOfWearID,
		Name:           params.Name,
		Description:    params.Description,
	}

	return u.toolRepo.Add(new_tool)
}
