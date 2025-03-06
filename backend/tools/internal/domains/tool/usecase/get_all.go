package tool_usecase

import (
	tool_params "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/domains/tool/params"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/models"
)

func (u *ToolUsecase) GetAll(params *tool_params.GetAllParams) ([]models.Tool, error) {
	return u.toolRepo.GetAll(params)
}
