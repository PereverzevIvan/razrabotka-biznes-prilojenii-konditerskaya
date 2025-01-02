package app

import (
	controllers_order "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/controllers/order"
	controllers_quality_control "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/controllers/quality_control"
)

func (app *App) initControllers() error {
	// api
	api := app.fiberApp.Group("/api")

	controllers_order.AddOrderControllerRoutes(
		api,
		app.serviceProvider.JWTService(),
		app.serviceProvider.OrderService(),
	)

	controllers_quality_control.AddQualityControlControllerRoutes(
		api,
		app.serviceProvider.JWTService(),
		app.serviceProvider.QualityControlService(),
	)

	return nil
}
