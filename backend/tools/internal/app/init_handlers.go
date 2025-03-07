package app

import (
	"log"
	"net"

	tool_failure_handler "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/domains/tool_failure/handler"
	tool_type_handler "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/domains/tool_type/handler"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/storage"
	"google.golang.org/grpc"
)

type RegisterHandlerFn func(grpc.ServiceRegistrar, *storage.Storage)

func (app *App) initHandlers() error {
	lis, err := net.Listen("tcp", "0.0.0.0:50041")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	srv := grpc.NewServer()

	registerHandlerFns := []RegisterHandlerFn{
		tool_type_handler.RegisterHandler,
		tool_failure_handler.RegisterHandler,
	}

	for _, fn := range registerHandlerFns {
		fn(srv, app.Storage)
	}

	// TODO: вынести в runApp и запускать в goroutine
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	} else {
		log.Printf("server listening at %v", lis.Addr())
	}
	return nil
}
