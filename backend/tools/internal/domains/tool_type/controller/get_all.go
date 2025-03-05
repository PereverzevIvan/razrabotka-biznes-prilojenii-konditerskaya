package tool_type_controller

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

func (c *toolTypeController) GetAll(ctx fiber.Ctx) error {
	tool_types, err := c.toolTypeUsecase.GetAll()
	if err != nil {
		log.Error(err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.Status(fiber.StatusOK).JSON(tool_types)
}
