package tool_controller

import (
	tool_params "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/domains/tool/params"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

func (controller *toolController) Edit(ctx fiber.Ctx) error {
	tool_id := fiber.Params[int](ctx, "id")
	if tool_id <= 0 {
		return ctx.Status(fiber.StatusBadRequest).SendString("invalid tool id")
	}

	params := &tool_params.ToolEditParams{}
	err := ctx.Bind().Body(&params)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	tool, err := controller.toolUsecase.Edit(tool_id, params)
	if err != nil {
		log.Error(err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.Status(fiber.StatusCreated).JSON(tool)
}
