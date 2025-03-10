package main

import (
	"context"
	"log"

	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/gateway/internal/app"
)

// @title Разработка бизнес-приложений - лаба 3
// @version 1.0
// @description Это API лабораторной работы 3 по дисциплине "Разработка бизнес-приложений".
// @description Тема проекта - кондитерская.
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3000
// @BasePath /api
// @securityDefinitions.apikey ApiKeyAuth
// @in cookie
// @name access_token
func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	App, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalln(err)
		return
	}

	err = App.Run()
	if err != nil {
		log.Fatalln(err)
		return
	}
}
