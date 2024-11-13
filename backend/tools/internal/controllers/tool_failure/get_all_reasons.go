package controllers_tool_failure

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

func (controller *toolFailureController) GetAllReasons(ctx fiber.Ctx) error {
	tool_failure_reasons, err := controller.toolFailureService.GetAllReasons()
	if err != nil {
		log.Error(err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.Status(fiber.StatusOK).JSON(tool_failure_reasons)
}
