package controllers_quality_control

import (
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/controllers"
	"github.com/gofiber/fiber/v3"
)

type qualityControllController struct {
	jwtService            controllers.IJWTService
	qualityControlService controllers.IQualityControlService
}

func AddQualityControlControllerRoutes(
	api fiber.Router,
	jwtService controllers.IJWTService,
	qualityControlService controllers.IQualityControlService,
) {
	controller := qualityControllController{
		jwtService:            jwtService,
		qualityControlService: qualityControlService,
	}

	api.Post("/quality-controls", controller.Create)
	api.Patch("/quality-controls", controller.Update)
	api.Get("/quality-controls/order/:order_id", controller.GetByOrderId)
}
