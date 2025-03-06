package quality_control_controller

import (
	quality_control_params "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/domains/quality_control/params"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

func (controller *qualityControllController) Create(ctx fiber.Ctx) error {

	user_id, err := controller.jwtUsecase.GetUserIDFromAccessTokenCookie(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	var params = &quality_control_params.CreateParams{}
	if err := ctx.Bind().Body(params); err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	params.MasterID = user_id

	if errs := params.Validate(); len(errs) != 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"errors": errs})
	}

	quality_control, err := controller.qualityControlUsecase.Create(params)
	if err != nil {
		log.Error(err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.Status(fiber.StatusCreated).JSON(quality_control)
}
