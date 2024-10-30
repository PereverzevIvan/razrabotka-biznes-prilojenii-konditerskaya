package main

import (
	"fmt"
	"net/http"

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
	fmt.Println(cfg)

	controllers.Init()

	http.HandleFunc("/*", func(w http.ResponseWriter, r *http.Request) {
		var ans_text string = fmt.Sprintf(
			"requested url %s not found; path is %s",
			r.URL,
			r.URL.Path,
		)

		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(ans_text))
	})

	err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.ConfigServer.Port), nil)
	if err != nil {
		panic(err)
	}

}
