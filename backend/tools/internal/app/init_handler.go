package app

import (
	"log"
	"net"

	tool_type_handler "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/domains/tool_type/handler"
	"google.golang.org/grpc"
)

func (app *App) initHandlers() error {
	lis, err := net.Listen("tcp", "0.0.0.0:50041")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	srv := grpc.NewServer()

	tool_type_handler.NewHandler(srv, app.Storage)

	if err := srv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	} else {
		log.Printf("server listening at %v", lis.Addr())
	}
	return nil
}
