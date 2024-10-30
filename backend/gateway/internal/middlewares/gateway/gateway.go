package middlewares_gateway

import (
	"fmt"

	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/gateway/configs"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/gateway/pkg/config_loader"
)

const (
	ROUTES_PATH_PARAM_NAME = "routes_path"
)

func InitGatewayRoutes() {
	var cfg configs.RoutesConfig
	config_loader.MustLoad(configs.RoutesPath, &cfg)

	fmt.Printf("Config gateway: %+v\n", cfg)
}

// func IpWhitelistMiddleware(allowedOrigins []configs.AllowedOrigin, next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		remoteIP, port, err := net.SplitHostPort(r.RemoteAddr)
// 		if err != nil {
// 			http.Error(w, "Invalid request", http.StatusBadRequest)
// 			return
// 		}

// 		for _, allowedOrigin := range allowedOrigins {
// 			if allowedOrigin.Host == remoteIP && allowedOrigin.Port == port {
// 				next.ServeHTTP(w, r)
// 				return
// 			}
// 		}

// 		http.Error(w, "Forbidden", http.StatusForbidden)
// 	})
// }
