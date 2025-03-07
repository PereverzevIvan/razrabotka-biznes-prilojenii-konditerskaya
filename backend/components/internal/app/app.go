package app

import (
	"fmt"

	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/configs"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/storage"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/gateway/pkg/config_loader"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"google.golang.org/grpc"
)

type App struct {
	config   *configs.Config
	Storage  *storage.Storage
	fiberApp *fiber.App

	toolConn *grpc.ClientConn
}

func NewApp(toolConn *grpc.ClientConn) (*App, error) {
	app := &App{
		toolConn: toolConn,
	}

	app.fiberApp = fiber.New()

	err := app.initDependencies()
	if err != nil {
		return nil, err
	}

	return app, nil
}

func (app *App) Run() error {
	err := app.fiberApp.Listen(fmt.Sprintf(":%d", app.config.ServerConfig.Port))
	if err != nil {
		return err
	}

	return nil
}

func (app *App) initDependencies() error {
	inits := []func() error{
		app.initConfigs,
		app.initStorage,
		app.initRoutes,
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

	err = config_loader.Load(configs.ConfigPath, &app.config)
	if err != nil {
		return err
	}

	log.Info("server config: ", app.config)

	return nil
}

func (app *App) initStorage() error {
	storage, err := storage.NewStorage(&app.config.DBConfig, &app.config.JWTConfig, app.toolConn)
	if err != nil {
		return err
	}
	app.Storage = storage
	return nil
}
