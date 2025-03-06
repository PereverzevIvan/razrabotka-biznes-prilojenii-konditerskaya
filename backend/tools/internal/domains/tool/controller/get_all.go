package tool_controller

import (
	tool_params "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/domains/tool/params"
	"github.com/gofiber/fiber/v3"
)

func (controller *toolController) GetAll(ctx fiber.Ctx) error {
	params := &tool_params.GetAllParams{}

	err := ctx.Bind().Query(&params)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	tools, err := controller.toolUsecase.GetAll(params)
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.Status(fiber.StatusOK).JSON(tools)
}
