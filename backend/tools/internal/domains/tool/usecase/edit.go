package tool_usecase

import (
	tool_params "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/domains/tool/params"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/models"
)

func (u *ToolUsecase) Edit(tool_id int, params *tool_params.ToolEditParams) (*models.Tool, error) {
	return u.toolRepo.Edit(tool_id, params)
}
