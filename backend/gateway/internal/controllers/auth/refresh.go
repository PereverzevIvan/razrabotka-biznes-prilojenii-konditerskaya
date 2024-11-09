package controllers_auth

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

func (controller *AuthController) Refresh(ctx fiber.Ctx) error {
	user_id, err := controller.jwtService.GetUserIDFromRefreshTokenCookie(ctx)
	if err != nil {
		log.Error(err)
		return ctx.SendStatus(fiber.StatusUnauthorized)
	}

	user, err := controller.userService.GetByID(user_id)
	if err != nil {
		log.Error(err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	access_token, refresh_token, err := controller.jwtService.GenerateAccessAndRefreshTokens(user)
	if err != nil {
		log.Error(err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	err = controller.setAccessRefreshTokens(ctx, access_token, refresh_token)
	if err != nil {
		log.Error(err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.SendStatus(fiber.StatusOK)
}
