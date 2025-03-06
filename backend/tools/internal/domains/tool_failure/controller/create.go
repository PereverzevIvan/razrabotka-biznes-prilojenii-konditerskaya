package tool_failure_controller

import (
	tool_failure_params "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/domains/tool_failure/params"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

func (controller *toolFailureController) Create(ctx fiber.Ctx) error {
	master_id, err := controller.jwtService.GetUserIDFromAccessTokenCookie(ctx)
	if err != nil {
		log.Error(err)
		return ctx.Status(fiber.StatusUnauthorized).SendString(err.Error())
	}

	params := &tool_failure_params.CreateParams{
		MasterID: master_id,
	}

	err = ctx.Bind().Body(&params)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	if validate_errs := params.Validate(); len(validate_errs) != 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"errors": validate_errs})
	}

	tool_failure, err := controller.toolFailureUsecase.Create(params)
	if err != nil {
		log.Error(err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.Status(fiber.StatusCreated).JSON(tool_failure)
}
