package controllers_auth

import (
	"time"

	"github.com/gofiber/fiber/v3"
)

func (controller *AuthController) setAccessRefreshTokens(
	ctx fiber.Ctx,
	access_token string,
	refresh_token string,
) error {
	// Отправляем токены
	ctx.Cookie(&fiber.Cookie{
		Name:    "access-token",
		Value:   access_token,
		Expires: time.Now().Add(controller.jwtService.AccessTokenExpiration()),
		// HTTPOnly: true,
	})
	ctx.Cookie(&fiber.Cookie{
		Name:     "refresh-token",
		Value:    refresh_token,
		Expires:  time.Now().Add(controller.jwtService.RefreshTokenExpiration()),
		HTTPOnly: true,
	})

	return nil
}
