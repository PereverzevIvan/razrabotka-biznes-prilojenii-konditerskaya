package app

import (
	tool_controller "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/domains/tool/controller"
	tool_failure_controller "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/domains/tool_failure/controller"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/storage"
	"github.com/gofiber/fiber/v3"
)

type AddRouteFn func(fiber.Router, *storage.Storage)

func (app *App) initRoutes() error {

	addRoutes := []AddRouteFn{
		tool_controller.AddRoutes,
		tool_failure_controller.AddRoutes,
	}

	for _, fn := range addRoutes {
		fn(app.fiberApp, app.Storage)
	}

	return nil
}
