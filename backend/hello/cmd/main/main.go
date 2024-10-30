package main

import (
	"fmt"
	"net/http"

	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/gateway/pkg/config_loader"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/hello/configs"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/hello/internal/controllers"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/hello/internal/middlewares"
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
	fmt.Println(cfg)

	mux := http.NewServeMux()

	controllers.Init(mux)

	wrappedMux := middlewares.IpWhitelistMiddleware(cfg.AllowedOrigins, mux)

	err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.ConfigServer.Port), wrappedMux)
	if err != nil {
		panic(err)
	}

}
