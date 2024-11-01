package controllers_auth

import "github.com/gofiber/fiber/v3"

func (controller *AuthController) Logout(ctx fiber.Ctx) error {

	// Очищаем куки
	ctx.ClearCookie("access_token")
	ctx.ClearCookie("refresh_token")

	return ctx.SendStatus(fiber.StatusOK)
}
