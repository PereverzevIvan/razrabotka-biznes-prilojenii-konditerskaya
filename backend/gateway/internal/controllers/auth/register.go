package controllers_auth

import (
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/gateway/models"
	params_auth "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/gateway/models/params/auth"
	"github.com/gofiber/fiber/v3"
)

func (controller *AuthController) Register(ctx fiber.Ctx) error {
	var params params_auth.RegisterParams
	err := ctx.Bind().Body(&params)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	user, err := controller.userService.Create(&params)
	if err != nil {
		if err.Error() == models.ErrUnique.Error() {
			return ctx.Status(fiber.StatusConflict).SendString("email already exists")
		}
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	// Генерируем jwt токены
	access_token, refresh_token, err := controller.jwtService.GenerateAccessAndRefreshTokens(user)
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	// Отправляем токены
	err = controller.setAccessRefreshTokens(ctx, access_token, refresh_token)
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.SendStatus(fiber.StatusOK)
}
