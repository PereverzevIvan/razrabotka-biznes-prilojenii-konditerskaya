package app

import controllers_order "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/controllers/order"

func (app *App) initControllers() error {
	// api
	api := app.fiberApp.Group("/api")

	controllers_order.AddOrderControllerRoutes(
		api,
		app.serviceProvider.JWTService(),
		app.serviceProvider.OrderService(),
	)

	return nil
}
