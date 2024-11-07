package controllers_utils

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
)

func SendBadRequestParamsResponse(ctx fiber.Ctx, err error) error {
	return ctx.Status(fiber.StatusBadRequest).SendString(
		fmt.Sprintf("Invalid params: %v", err))
}

func SendBadRequestResponseWithMessage(ctx fiber.Ctx, message string) error {
	return ctx.Status(fiber.StatusBadRequest).SendString(message)
}

func SendResponseMessages(ctx fiber.Ctx, messages []string) error {
	var answer struct {
		Messages []string `json:"messages"`
	}

	answer.Messages = messages

	return ctx.JSON(answer)
}
