package controllers

import (
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/models"
	params_tool "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/models/params/tool"
	params_tool_failure "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/models/params/tool_failure"
	"github.com/gofiber/fiber/v3"
)

type IJWTService interface {
	GetRoleFromAccessTokenCookie(ctx fiber.Ctx) (string, error)
	GetUserIDFromAccessTokenCookie(ctx fiber.Ctx) (int, error)
}

type IToolService interface {
	GetAll(params *params_tool.GetAllParams) ([]models.Tool, error)
	GetByID(tool_id int) (*models.Tool, error)
	Add(params *params_tool.ToolAddParams) (*models.Tool, error)
	Edit(tool_id int, params *params_tool.ToolEditParams) (*models.Tool, error)
	Delete(tool_id int) error
}

type IToolTypeService interface {
	GetAll() ([]models.ToolType, error)
}

type IToolFailureReasonService interface {
	GetAll() ([]models.ToolFailureReason, error)
}

type IToolFailureService interface {
	Create(params *params_tool_failure.CreateParams) (*models.ToolFailure, error)
	AddFixedAt(params *params_tool_failure.AddFixedAtParams) error
	GetAll() ([]models.ToolFailure, error)
	GetAllReasons() ([]models.ToolFailureReason, error)
}
