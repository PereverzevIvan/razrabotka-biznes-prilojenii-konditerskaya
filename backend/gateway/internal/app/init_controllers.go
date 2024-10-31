package app

import "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/gateway/internal/controllers"

func (app *App) initControllers() error {

	// Proxy до микросервисов
	root := app.fiberApp.Group("/")
	controllers.AddGatewayRoutes(root)

	// api авторизации
	api := app.fiberApp.Group("/api")

	controllers.AddAuthControllerRoutes(
		api,
		app.serviceProvider.UserService(),
	)

	return nil
}
