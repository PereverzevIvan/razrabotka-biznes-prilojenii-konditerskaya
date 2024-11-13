package controllers_tool_failure

import (
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/controllers"
	"github.com/gofiber/fiber/v3"
)

type toolFailureController struct {
	jwtService controllers.IJWTService
	// toolFailureReasonService controllers.IToolFailureReasonService
	toolFailureService controllers.IToolFailureService
}

func AddToolFailureControllerRoutes(
	api fiber.Router,
	jwtService controllers.IJWTService,
	// toolFailureReasonService controllers.IToolFailureReasonService,
	toolFailureService controllers.IToolFailureService,
) {

	controller := &toolFailureController{
		jwtService: jwtService,
		// toolFailureReasonService: toolFailureReasonService,
		toolFailureService: toolFailureService,
	}

	api.Get("/tool-failures", controller.GetAll)
	api.Get("/tool-failures/reasons", controller.GetAllReasons)
	api.Post("/tool-failures", controller.Create)
	api.Post("/tool-failures/add-fixed-at", controller.AddFixedAt)
}
