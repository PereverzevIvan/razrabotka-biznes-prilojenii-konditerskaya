package middlewares_auth

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"github.com/golang-jwt/jwt"
)

func (middleware *authMiddleware) Authorized(ctx fiber.Ctx) error {
	access_token := ctx.Cookies("access_token")
	log.Info("access_token: ", access_token)
	if access_token == "" {
		return ctx.SendStatus(fiber.StatusUnauthorized)
	}

	token, err := middleware.jwtService.ValidateToken(access_token)
	if err != nil {
		return ctx.SendStatus(fiber.StatusUnauthorized)
	}

	user_id := token.Claims.(jwt.MapClaims)["id"].(float64)
	ctx.Locals("user_id", int(user_id))

	role := token.Claims.(jwt.MapClaims)["role"].(string)
	ctx.Locals("role", role)

	return ctx.Next()
}
