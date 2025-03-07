package app

import (
	"context"
	"flag"

	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/proto/pkg/api/tool"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/proto/pkg/api/tool_failure"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/proto/pkg/api/tool_type"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	// command-line options:
	// gRPC server endpoint
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "golang_tools:50041", "gRPC server endpoint")
)

type RegisterServiceHandlerFromEndpointFn func(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error

type RegisterService struct {
	Fn RegisterServiceHandlerFromEndpointFn
}

func (app *App) initGrpcGateway() error {
	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	app.mux = runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	err := tool_type.RegisterToolTypeServiceHandlerFromEndpoint(app.ctx, app.mux, *grpcServerEndpoint, opts)
	if err != nil {
		return err
	}
	err = tool.RegisterToolServiceHandlerFromEndpoint(app.ctx, app.mux, *grpcServerEndpoint, opts)
	if err != nil {
		return err
	}
	err = tool_failure.RegisterToolFailureServiceHandlerFromEndpoint(app.ctx, app.mux, *grpcServerEndpoint, opts)
	if err != nil {
		return err
	}

	return nil
}
