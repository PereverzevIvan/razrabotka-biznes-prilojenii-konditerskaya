package controllers

import "github.com/gofiber/fiber/v3"

type AuthController struct {
	userService IUserService
}

func AddAuthControllerRoutes(
	api fiber.Router,
	userService IUserService,
) {
	controller := &AuthController{
		userService: userService,
	}

	api.Get("/user/:id", controller.GetUserById)
	// api
	// controller.AddLoginRoute()
}

func (c *AuthController) GetUserById(ctx fiber.Ctx) error {
	requested_user_id := fiber.Params[int](ctx, "id")
	if requested_user_id <= 0 {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	user, err := c.userService.GetByID(1)
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}
	if user == nil {
		return ctx.SendStatus(fiber.StatusNotFound)
	}

	return ctx.JSON(user)
}
