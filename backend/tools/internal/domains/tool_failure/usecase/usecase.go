package tool_failure_usecase

import (
	tool_failure_params "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/domains/tool_failure/params"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/models"
)

type IToolFailureRepo interface {
	GetAll() ([]models.ToolFailure, error)
	Create(params *tool_failure_params.CreateParams) (*models.ToolFailure, error)
	AddFixedAt(params *tool_failure_params.AddFixedAtParams) error
	GetAllReasons() ([]models.ToolFailureReason, error)
}

type ToolFailureUsecase struct {
	toolFailureRepo IToolFailureRepo
}

func NewToolFailureUsecase(toolFailureRepo IToolFailureRepo) *ToolFailureUsecase {
	return &ToolFailureUsecase{
		toolFailureRepo: toolFailureRepo,
	}
}
