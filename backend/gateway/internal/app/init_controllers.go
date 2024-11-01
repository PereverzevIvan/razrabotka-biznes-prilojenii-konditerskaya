package app

import (
	controllers_auth "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/gateway/internal/controllers/auth"
	controllers_gateway "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/gateway/internal/controllers/gateway"
)

func (app *App) initControllers() error {

	// Proxy до микросервисов
	root := app.fiberApp.Group("/")
	controllers_gateway.AddGatewayRoutes(
		root,
		app.serviceProvider.AuthMiddleware(),
	)

	// api авторизации
	api := app.fiberApp.Group("/api")

	controllers_auth.AddAuthControllerRoutes(
		api,
		app.serviceProvider.UserService(),
		app.serviceProvider.JWTService(),
		app.serviceProvider.AuthMiddleware(),
	)

	return nil
}
