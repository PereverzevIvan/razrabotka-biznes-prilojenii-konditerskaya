package tool_failure_controller

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

func (controller *toolFailureController) GetAllReasons(ctx fiber.Ctx) error {
	tool_failure_reasons, err := controller.toolFailureUsecase.GetAllReasons()
	if err != nil {
		log.Error(err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.Status(fiber.StatusOK).JSON(tool_failure_reasons)
}
