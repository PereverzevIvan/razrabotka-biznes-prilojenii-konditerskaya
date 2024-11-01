package controllers_gateway

import (
	"time"

	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/gateway/configs"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/gateway/internal/controllers"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/gateway/pkg/config_loader"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"github.com/gofiber/fiber/v3/middleware/proxy"
)

const (
	ROUTES_PATH_PARAM_NAME = "routes_path"
)

func AddGatewayRoutes(root fiber.Router, authMiddleware controllers.IAuthMiddleware) {
	var routes_config configs.RoutesConfig
	config_loader.MustLoad(configs.RoutesPath, &routes_config)

	// fmt.Printf("Config routes: %+v\n", routes_config)

	for _, microservice := range routes_config.Microservices {
		for _, route := range microservice.Routes {

			route_middlewares := parseRouteMiddlewares(
				authMiddleware,
				route.Middlewares,
			)

			addRoutesMethods(
				root,
				microservice.Url,
				route.Path,
				route.Methods,
				route_middlewares,
			)
		}
	}
}

func parseRouteMiddlewares(
	authMiddleware controllers.IAuthMiddleware,
	middlewares []configs.RouteMiddleware,
) []fiber.Handler {
	var route_middlewares []fiber.Handler
	for _, middleware := range middlewares {
		switch middleware {
		case configs.Auth:
			route_middlewares = append(route_middlewares, authMiddleware.Authorized)
		case configs.Admin:
			route_middlewares = append(route_middlewares, authMiddleware.Admin)
		default:
			log.Error("Unknown middleware: %s\n", middleware)
		}
	}
	return route_middlewares
}

func generateForwardRequestFunc(microservice_url string) func(ctx fiber.Ctx) error {
	return func(ctx fiber.Ctx) error {
		if err := proxy.DoTimeout(ctx, microservice_url+ctx.OriginalURL(), time.Second*5); err != nil {
			return err
		}

		// Remove Server header from response
		ctx.Response().Header.Del(fiber.HeaderServer)
		return nil
	}
}

func addRoutesMethods(
	group fiber.Router,
	microservice_url string,
	path string,
	methods []configs.RouteMethods,
	route_middlewares []fiber.Handler,
) {

	for _, method := range methods {
		switch method {
		case configs.GET:
			group.Get(path, generateForwardRequestFunc(microservice_url), route_middlewares...)
		case configs.POST:
			group.Post(path, generateForwardRequestFunc(microservice_url), route_middlewares...)
		case configs.PUT:
			group.Put(path, generateForwardRequestFunc(microservice_url), route_middlewares...)
		case configs.PATCH:
			group.Patch(path, generateForwardRequestFunc(microservice_url), route_middlewares...)
		case configs.DELETE:
			group.Delete(path, generateForwardRequestFunc(microservice_url), route_middlewares...)
		default:
			log.Error("Unknown method: %s\n", method)
		}
	}
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
