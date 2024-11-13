package controllers_order

import (
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/controllers"
	"github.com/gofiber/fiber/v3"
)

type orderController struct {
	orderService controllers.IOrderService
}

func AddOrderControllerRoutes(
	api fiber.Router,
	orderService controllers.IOrderService,
) {
	controller := orderController{
		orderService: orderService,
	}

	api.Get("orders/", controller.GetAll)
}
