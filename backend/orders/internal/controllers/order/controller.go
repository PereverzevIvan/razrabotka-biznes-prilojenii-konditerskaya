package controllers_order

import (
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/controllers"
	"github.com/gofiber/fiber/v3"
)

type orderController struct {
	jwtService   controllers.IJWTService
	orderService controllers.IOrderService
}

func AddOrderControllerRoutes(
	api fiber.Router,
	jwtService controllers.IJWTService,
	orderService controllers.IOrderService,
) {
	controller := orderController{
		jwtService:   jwtService,
		orderService: orderService,
	}

	api.Get("orders/", controller.GetAll)
	api.Post("orders/", controller.Create)
}
