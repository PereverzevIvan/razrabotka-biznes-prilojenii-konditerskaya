package controllers_tool_type

import (
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/controllers"
	"github.com/gofiber/fiber/v3"
)

type toolTypeController struct {
	toolTypeService controllers.IToolTypeService
}

func AddToolTypeControllerRoutes(
	api fiber.Router,
	toolTypeService controllers.IToolTypeService,
) {

	controller := &toolTypeController{
		toolTypeService: toolTypeService,
	}

	api.Get("/tool-types", controller.GetAll)
}
