package controllers_auth

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

func (controller *AuthController) GetUserById(ctx fiber.Ctx) error {
	user_id, err := controller.jwtService.GetUserIDFromAccessTokenCookie(ctx)
	if err != nil {
		log.Error(err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	requested_user_id := fiber.Params[int](ctx, "id")
	if requested_user_id <= 0 {
		requested_user_id = user_id
	}

	if user_id != requested_user_id {
		return ctx.SendStatus(fiber.StatusForbidden)
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
