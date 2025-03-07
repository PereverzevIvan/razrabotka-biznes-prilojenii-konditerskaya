package app

import (
	"fmt"
	"log"
	"net"

	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/gateway/pkg/config_loader"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/configs"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/storage"
	"google.golang.org/grpc"
	// "github.com/gofiber/fiber/v3/log"
)

type App struct {
	config  *configs.Config
	Storage *storage.Storage

	lis net.Listener
	srv *grpc.Server
}

func NewApp() (*App, error) {
	app := &App{}

	err := app.initDependencies()
	if err != nil {
		return nil, err
	}

	return app, nil
}

func (app *App) Run() error {

	err := app.srv.Serve(app.lis)
	if err != nil {
		return fmt.Errorf("failed to serve: %v", err)
	}

	return nil
}

func (app *App) initDependencies() error {
	inits := []func() error{
		app.initConfigs,
		app.initStorage,
		app.initGrpcHandlers,
		app.initNetListener,
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
		return err
	}

	app.config = &cfg

	log.Println("server config: ", app.config)

	return nil
}

func (app *App) initStorage() error {
	storage, err := storage.NewStorage(&app.config.DBConfig, &app.config.JWTConfig)
	if err != nil {
		return err
	}
	app.Storage = storage
	return nil
}
