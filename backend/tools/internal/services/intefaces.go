package services

import (
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/models"
	params_tool "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/models/params/tool"
)

type IToolRepo interface {
	GetAll(params *params_tool.ToolGetAllParams) ([]models.Tool, error)
	GetByID(tool_id int) (*models.Tool, error)
	Add(tool *models.Tool) (*models.Tool, error)
	Edit(tool_id int, params *params_tool.ToolEditParams) (*models.Tool, error)
	Delete(tool_id int) error
}

type IToolTypeRepo interface {
	GetAll() ([]models.ToolType, error)
}
