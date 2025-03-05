package tool_failure_controller

import (
	tool_failure_deps "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/domains/tool_failure/deps"
	tool_failure_params "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/domains/tool_failure/params"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/models"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/storage"
	"github.com/gofiber/fiber/v3"
)

type IToolFailureUsecase interface {
	Create(params *tool_failure_params.CreateParams) (*models.ToolFailure, error)
	AddFixedAt(params *tool_failure_params.AddFixedAtParams) error
	GetAll() ([]models.ToolFailure, error)
	GetAllReasons() ([]models.ToolFailureReason, error)
}

type IJWTService interface {
	GetRoleFromAccessTokenCookie(ctx fiber.Ctx) (string, error)
	GetUserIDFromAccessTokenCookie(ctx fiber.Ctx) (int, error)
}

type toolFailureController struct {
	// toolFailureReasonService controllers.IToolFailureReasonService
	toolFailureUsecase IToolFailureUsecase
	jwtService         IJWTService
}

func AddRoutes(
	api fiber.Router,
	storage *storage.Storage,
) {

	deps := tool_failure_deps.NewDeps(storage.DB, storage.JwtConfig)

	controller := &toolFailureController{
		jwtService: deps.JwtService(),
		// toolFailureReasonService: toolFailureReasonService,
		toolFailureUsecase: deps.ToolFailureUsecase(),
	}

	apiTools := api.Group("/tool-failures")

	apiTools.Get("/", controller.GetAll)
	apiTools.Post("/", controller.Create)
	apiTools.Get("/reasons", controller.GetAllReasons)
	apiTools.Post("/add-fixed-at", controller.AddFixedAt)
}
