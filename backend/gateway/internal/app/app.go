package app

import (
	"context"
	"fmt"
	"net/http"

	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/gateway/configs"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/gateway/pkg/config_loader"
	"github.com/gofiber/fiber/v3"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

type App struct {
	ctx             context.Context
	config          *configs.Config
	Storage         *Storage
	serviceProvider *ServiceProvider
	fiberApp        *fiber.App
	mux             *runtime.ServeMux
}

func NewApp(ctx context.Context) (*App, error) {
	app := &App{
		ctx: ctx,
	}

	app.fiberApp = fiber.New()

	err := app.initDependencies()
	if err != nil {
		return nil, err
	}

	return app, nil
}

func (app *App) Run() error {

	wg := make(chan error, 2)
	go func() {
		// Start HTTP server (and proxy calls to gRPC server endpoint)
		wg <- http.ListenAndServe(":8081", app.mux)
	}()
	go func() {
		// Start HTTP fiber gateway
		wg <- app.fiberApp.Listen(fmt.Sprintf(":%d", app.config.ServerConfig.Port))
	}()

	for i := 0; i < 2; i++ {
		err := <-wg
		if err != nil {
			return err
		}
	}

	return nil
}

func (app *App) initDependencies() error {
	inits := []func() error{
		app.initConfigs,
		app.initStorage,
		app.initServiceProvider,
		app.initControllers,
		app.initGrpcGateway,
	}

	for _, init := range inits {
		if err := init(); err != nil {
			return err
		}
	}

	return nil
}

func (app *App) initConfigs() error {
	err := configs.LoadCmdParams()
	if err != nil {
		return err
	}

	cfg := configs.Config{}
	err = config_loader.Load(configs.ConfigPath, &cfg)
	if err != nil {
		return fmt.Errorf("initConfigs: %w", err)
	}

	app.config = &cfg

	fmt.Println("server config: ", app.config)

	return nil
}

func (app *App) initStorage() error {
	storage, err := NewStorage(&app.config.DBConfig)
	if err != nil {
		return err
	}
	app.Storage = storage
	return nil
}

func (app *App) initServiceProvider() error {
	app.serviceProvider = newServiceProvider(
		app.Storage.Conn,
		&app.config.JWTConfig,
	)

	return nil
}
