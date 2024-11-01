package controllers_auth

import (
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/gateway/internal/controllers"
	"github.com/gofiber/fiber/v3"
)

type AuthController struct {
	userService controllers.IUserService
	jwtService  controllers.IJWTService
}

func AddAuthControllerRoutes(
	api fiber.Router,
	userService controllers.IUserService,
	jwtService controllers.IJWTService,
	authMiddleware controllers.IAuthMiddleware,
) {
	controller := &AuthController{
		userService: userService,
		jwtService:  jwtService,
	}

	api.Post("/register", controller.Register)
	api.Post("/login", controller.Login)
	api.Get("/logout", controller.Logout)
	api.Get("/refresh", controller.Refresh, authMiddleware.Authorized)
	api.Get("/user/:id", controller.GetUserById, authMiddleware.Authorized)
	// api
	// controller.AddLoginRoute()
}
