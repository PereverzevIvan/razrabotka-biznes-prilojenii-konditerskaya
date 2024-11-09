package services_tool

import (
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/models"
	params_tool "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/models/params/tool"
)

func (service *toolService) Edit(tool_id int, params *params_tool.ToolEditParams) (*models.Tool, error) {
	return service.toolRepo.Edit(tool_id, params)
}
