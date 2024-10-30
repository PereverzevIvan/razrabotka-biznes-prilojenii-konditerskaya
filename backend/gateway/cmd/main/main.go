package main

import (
	"fmt"

	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/gateway/configs"
	middlewares_gateway "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/gateway/internal/middlewares/gateway"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/gateway/pkg/config_loader"
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
func main() {
	configs.MustLoadCmdParams()

	var cfg configs.Config
	config_loader.MustLoad(configs.ConfigPath, &cfg)
	fmt.Println(cfg)

	// app := fiber.New()

	middlewares_gateway.InitGatewayRoutes()

	// gateway := app.Group("")

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
