package middlewares_auth

import (
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/gateway/models"
	"github.com/gofiber/fiber/v3"
)

func (middleware *authMiddleware) Admin(ctx fiber.Ctx) error {
	role, err := middleware.jwtService.GetRoleFromAccessTokenCookie(ctx)
	if err != nil {
		return ctx.SendStatus(fiber.StatusUnauthorized)
	}

	if role != models.KRoleClientManager.Name() {
		return ctx.SendStatus(fiber.StatusUnauthorized)
	}

	return ctx.Next()
}
