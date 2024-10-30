package main

import (
	"fmt"

	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/gateway/pkg/config_loader"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/hello/configs"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/hello/internal/controllers"
)

// @title Ручка приветствия
// @version 1.0
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:7777
// @BasePath /api
// @securityDefinitions.apikey ApiKeyAuth
// @in cookie
// @name access-token
const (
	CONFIG_PATH_PARAM_NAME = "config_path"
)

func main() {
	var cfg configs.Config

	config_loader.MustLoadFromCmd(CONFIG_PATH_PARAM_NAME, &cfg)
	// cfg := configloader.MustLoad()
	fmt.Println(cfg)

	controllers.Init(cfg.ConfigServer.Port)
	// conn := service.NewStorage(cfg.ConfigDatabase)
	// fmt.Println(conn)

	// app := fiber.New()

	// app.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"*", "http://localhost:5043", "http://localhost:5173"},
	// 	AllowCredentials: true, // Разрешение отправки кук
	// }))

	// // controllers.InitControllers(app, conn.Conn, &cfg.ConfigJWT)

	// log.Info(app.Listen(fmt.Sprintf(":%d", cfg.ConfigServer.Port)))
}
