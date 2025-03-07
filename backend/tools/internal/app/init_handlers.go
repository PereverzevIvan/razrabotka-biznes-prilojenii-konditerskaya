package app

import (
	tool_handler "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/domains/tool/handler"
	tool_failure_handler "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/domains/tool_failure/handler"
	tool_type_handler "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/domains/tool_type/handler"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/storage"
	"google.golang.org/grpc"
)

type RegisterHandlerFn func(grpc.ServiceRegistrar, *storage.Storage)

func (app *App) initGrpcHandlers() error {
	app.srv = grpc.NewServer()

	registerHandlerFns := []RegisterHandlerFn{
		tool_handler.RegisterHandler,
		tool_type_handler.RegisterHandler,
		tool_failure_handler.RegisterHandler,
	}

	for _, fn := range registerHandlerFns {
		fn(app.srv, app.Storage)
	}

	return nil
}
