package tool_usecase

import (
	tool_params "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/domains/tool/params"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/models"
)

type IToolRepo interface {
	GetAll(params *tool_params.GetAllParams) ([]models.Tool, error)
	GetByID(tool_id int) (*models.Tool, error)
	Add(tool *models.Tool) (*models.Tool, error)
	Edit(tool_id int, params *tool_params.ToolEditParams) (*models.Tool, error)
	Delete(tool_id int) error
}

type ToolUsecase struct {
	toolRepo IToolRepo
}

func NewToolUsecase(toolRepo IToolRepo) *ToolUsecase {
	return &ToolUsecase{
		toolRepo: toolRepo,
	}
}
