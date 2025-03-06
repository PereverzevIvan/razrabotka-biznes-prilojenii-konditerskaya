package quality_control_controller

import (
	quality_control_params "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/domains/quality_control/params"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

func (controller *qualityControllController) Update(ctx fiber.Ctx) error {

	var params = &quality_control_params.UpdateParams{}
	err := ctx.Bind().Body(params)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	if errs := params.Validate(); len(errs) != 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"errors": errs})
	}

	quality_control, err := controller.qualityControlUsecase.Update(params)
	if err != nil {
		log.Error(err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.Status(fiber.StatusCreated).JSON(quality_control)
}
