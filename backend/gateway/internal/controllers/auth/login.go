package controllers_auth

import (
	params_auth "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/gateway/models/params/auth"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

func (controller *AuthController) Login(ctx fiber.Ctx) error {
	var params params_auth.LoginParams
	err := ctx.Bind().Body(&params)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	// Получаем пользователя и проверяем пароль
	user, err := controller.userService.GetByEmail(params.Email)
	if err != nil {
		log.Error(err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}
	if user == nil {
		return ctx.Status(fiber.StatusNotFound).SendString("email or password is incorrect")
	}

	is_password_correct := controller.userService.IsPasswordCorrect(user, params.Password)
	if !is_password_correct {
		return ctx.Status(fiber.StatusNotFound).SendString("email or password is incorrect")
	}

	// Генерируем jwt токены
	access_token, refresh_token, err := controller.jwtService.GenerateAccessAndRefreshTokens(user)
	if err != nil {
		log.Error(err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	err = controller.setAccessRefreshTokens(ctx, access_token, refresh_token)
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.Status(fiber.StatusOK).SendString("Login successful")
}
