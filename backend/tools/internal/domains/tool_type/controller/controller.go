package tool_type_controller

import (
	tool_type_deps "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/domains/tool_type/deps"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/models"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/storage"
	"github.com/gofiber/fiber/v3"
)

type IToolTypeUsecase interface {
	GetAll() ([]models.ToolType, error)
}

type toolTypeController struct {
	toolTypeUsecase IToolTypeUsecase
}

func AddRoutes(
	api fiber.Router,
	storage *storage.Storage,
) {
	deps := tool_type_deps.NewDeps(storage.DB)

	controller := &toolTypeController{
		toolTypeUsecase: deps.ToolTypeUsecase(),
	}

	apiToolTyps := api.Group("/tool-types")

	apiToolTyps.Get("/", controller.GetAll)
}
