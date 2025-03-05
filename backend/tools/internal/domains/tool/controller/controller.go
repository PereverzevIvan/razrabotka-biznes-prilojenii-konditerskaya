package tool_controller

import (
	tool_deps "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/domains/tool/deps"
	tool_params "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/domains/tool/params"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/models"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/storage"
	"github.com/gofiber/fiber/v3"
)

type IToolUsecase interface {
	GetAll(params *tool_params.GetAllParams) ([]models.Tool, error)
	GetByID(tool_id int) (*models.Tool, error)
	Add(params *tool_params.ToolAddParams) (*models.Tool, error)
	Edit(tool_id int, params *tool_params.ToolEditParams) (*models.Tool, error)
	Delete(tool_id int) error
}

type toolController struct {
	toolUsecase IToolUsecase
}

func AddRoute(
	api fiber.Router,
	storage *storage.Storage,
) {
	deps := tool_deps.NewDeps(storage.DB)

	controller := &toolController{
		toolUsecase: deps.ToolUsecase(),
	}

	apiTools := api.Group("/tools")

	apiTools.Get("/", controller.GetAll)
	apiTools.Post("/", controller.Add)
	apiTools.Get("/:id", controller.GetByID)
	apiTools.Patch("/:id", controller.Edit)
	apiTools.Delete("/:id", controller.Delete)
}
