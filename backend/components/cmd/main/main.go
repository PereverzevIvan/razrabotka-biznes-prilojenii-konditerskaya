package main

import (
	"log"

	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/app"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// @title Разработка бизнес-приложений - лаба 3
// @version 1.0
// @description Микросервис компонентов: ингредиентов, украшений
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3000
// @BasePath /api
// @securityDefinitions.apikey ApiKeyAuth
// @in cookie
// @name access_token
func main() {
	toolConn, err := grpc.NewClient("golang_tools:50041", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer toolConn.Close()

	App, err := app.NewApp(toolConn)
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
