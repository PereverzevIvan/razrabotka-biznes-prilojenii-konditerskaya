package services

import (
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/models"
	params_tool "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/models/params/tool"
	params_tool_failure "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/models/params/tool_failure"
)

type IToolRepo interface {
	GetAll(params *params_tool.GetAllParams) ([]models.Tool, error)
	GetByID(tool_id int) (*models.Tool, error)
	Add(tool *models.Tool) (*models.Tool, error)
	Edit(tool_id int, params *params_tool.ToolEditParams) (*models.Tool, error)
	Delete(tool_id int) error
}

type IToolTypeRepo interface {
	GetAll() ([]models.ToolType, error)
}

type IToolFailureRepo interface {
	Create(params *params_tool_failure.CreateParams) (*models.ToolFailure, error)
	AddFixedAt(params *params_tool_failure.AddFixedAtParams) error
}
