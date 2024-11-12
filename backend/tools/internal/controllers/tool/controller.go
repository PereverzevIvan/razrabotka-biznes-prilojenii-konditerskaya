package controllers_tool

import (
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/controllers"
	"github.com/gofiber/fiber/v3"
)

type toolController struct {
	toolService controllers.IToolService
}

func AddToolControllerRoutes(
	api fiber.Router,
	toolService controllers.IToolService,
) {

	controller := &toolController{
		toolService: toolService,
	}

	api.Get("/tools", controller.GetAll)
	api.Get("/tools/:id", controller.GetByID)
	api.Post("/tools/add", controller.Add)
	api.Delete("/tools/:id", controller.Delete)
	api.Patch("/tools/:id", controller.Edit)
}