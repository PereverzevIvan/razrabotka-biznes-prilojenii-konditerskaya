package middlewares_auth

import (
	"github.com/gofiber/fiber/v3"
)

func (middleware *authMiddleware) Authorized(ctx fiber.Ctx) error {
	access_token := ctx.Cookies("access_token")
	if access_token == "" {
		return ctx.SendStatus(fiber.StatusUnauthorized)
	}

	_, err := middleware.jwtService.ValidateToken(access_token)
	if err != nil {
		return ctx.SendStatus(fiber.StatusUnauthorized)
	}

	return ctx.Next()
}
