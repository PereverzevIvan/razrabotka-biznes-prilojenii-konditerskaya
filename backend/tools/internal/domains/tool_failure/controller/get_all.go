package tool_failure_controller

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

func (controller *toolFailureController) GetAll(ctx fiber.Ctx) error {
	tool_failures, err := controller.toolFailureUsecase.GetAll()
	if err != nil {
		log.Error(err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.Status(fiber.StatusOK).JSON(tool_failures)
}
