package controllers_tool

import (
	params_tool "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/models/params/tool"
	"github.com/gofiber/fiber/v3"
)

func (controller *toolController) GetAll(ctx fiber.Ctx) error {
	params := &params_tool.GetAllParams{}

	err := ctx.Bind().Query(&params)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	tools, err := controller.toolService.GetAll(params)
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.Status(fiber.StatusOK).JSON(tools)
}
