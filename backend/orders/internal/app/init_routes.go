package app

import (
	order_controller "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/domains/order/controller"
	quality_control_controller "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/domains/quality_control/controller"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/storage"
	"github.com/gofiber/fiber/v3"
)

type AddRouteFn func(fiber.Router, *storage.Storage)

func (app *App) initControllers() error {
	// api
	api := app.fiberApp.Group("/api")

	addRoutes := []AddRouteFn{
		order_controller.AddRoutes,
		quality_control_controller.AddRoutes,
	}

	for _, fn := range addRoutes {
		fn(api, app.Storage)
	}

	return nil
}
