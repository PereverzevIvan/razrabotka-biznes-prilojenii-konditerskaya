package app

import (
	"fmt"

	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/gateway/configs"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/gateway/pkg/config_loader"
	"github.com/gofiber/fiber/v3"
)

type App struct {
	config          *configs.Config
	Storage         *Storage
	serviceProvider *ServiceProvider
	fiberApp        *fiber.App
}

func NewApp() (*App, error) {
	app := &App{}

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
		app.initServiceProvider,
		app.initControllers,
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
