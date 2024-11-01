package controllers_auth

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

func (controller *AuthController) GetUserById(ctx fiber.Ctx) error {
	requested_user_id := fiber.Params[int](ctx, "id")
	if requested_user_id <= 0 {
		return ctx.Status(fiber.StatusBadRequest).SendString("invalid user id")
	}

	user, err := controller.userService.GetByID(requested_user_id)
	if err != nil {
		log.Error(err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}
	if user == nil {
		return ctx.SendStatus(fiber.StatusNotFound)
	}

	return ctx.JSON(user)
}
