package main

import (
	"fmt"

	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/gateway/configs"
	configloader "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/gateway/pkg/config_loader"
	configloader_utils "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/gateway/pkg/config_loader/utils"
)

// @title Разработка бизнес-приложений - лаба 3
// @version 1.0
// @description Это API лабораторной работы 3 по дисциплине "Разработка бизнес-приложений".
// @description Тема проекта - кондитерская.
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

	configloader.MustLoadFromCmd(
		configloader_utils.FetchCmdParamValue(CONFIG_PATH_PARAM_NAME),
		&cfg)
	fmt.Println(cfg)

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
