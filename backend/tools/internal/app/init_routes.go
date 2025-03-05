package app

import (
	tool_controller "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/domains/tool/controller"
	tool_type_controller "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/domains/tool_type/controller"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/storage"
	"github.com/gofiber/fiber/v3"
)

type AddRouteFn func(fiber.Router, *storage.Storage)

func (app *App) initRoutes() error {

	addRoutes := []AddRouteFn{
		tool_controller.AddRoutes,
		tool_type_controller.AddRoutes,
	}

	for _, fn := range addRoutes {
		fn(app.fiberApp, app.Storage)
	}

	// // Proxy до микросервисов
	// root := app.fiberApp.Group("/")
	// controllers_gateway.AddGatewayRoutes(
	// 	root,
	// 	app.serviceProvider.AuthMiddleware(),
	// )

	// api
	// api := app.fiberApp.Group("/api")

	// controllers_tool.AddToolControllerRoutes(
	// 	api,
	// 	app.serviceProvider.ToolService(),
	// )

	// controllers_tool_type.AddToolTypeControllerRoutes(
	// 	api,
	// 	app.serviceProvider.ToolTypeService(),
	// )

	// controllers_tool_failure.AddToolFailureControllerRoutes(
	// 	api,
	// 	app.serviceProvider.JWTService(),
	// 	// app.serviceProvider.ToolFailureReasonService(),
	// 	app.serviceProvider.ToolFailureService(),
	// )

	return nil
}
