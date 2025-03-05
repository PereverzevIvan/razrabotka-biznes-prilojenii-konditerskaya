package services

import (
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/models"
	params_tool_failure "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/models/params/tool_failure"
)

type IToolTypeRepo interface {
	GetAll() ([]models.ToolType, error)
}

type IToolFailureRepo interface {
	GetAll() ([]models.ToolFailure, error)
	Create(params *params_tool_failure.CreateParams) (*models.ToolFailure, error)
	AddFixedAt(params *params_tool_failure.AddFixedAtParams) error
	GetAllReasons() ([]models.ToolFailureReason, error)
}
