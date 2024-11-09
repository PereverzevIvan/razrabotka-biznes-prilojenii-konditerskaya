package app

import (
	controllers_tool "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/controllers/tool"
	controllers_tool_type "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/controllers/tool_type"
)

func (app *App) initControllers() error {

	// // Proxy до микросервисов
	// root := app.fiberApp.Group("/")
	// controllers_gateway.AddGatewayRoutes(
	// 	root,
	// 	app.serviceProvider.AuthMiddleware(),
	// )

	// api
	api := app.fiberApp.Group("/api")

	controllers_tool.AddToolControllerRoutes(
		api,
		app.serviceProvider.ToolService(),
	)

	controllers_tool_type.AddToolTypeControllerRoutes(
		api,
		app.serviceProvider.ToolTypeService(),
	)

	return nil
}
