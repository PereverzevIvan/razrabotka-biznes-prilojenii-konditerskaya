package controllers_quality_control

import (
	params_quality_control "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/models/params/quality_control"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

func (controller *qualityControllController) Create(ctx fiber.Ctx) error {

	user_id, err := controller.jwtService.GetUserIDFromAccessTokenCookie(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	var params = &params_quality_control.CreateParams{}
	if err := ctx.Bind().Body(params); err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	params.MasterID = user_id

	if errs := params.Validate(); len(errs) != 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"errors": errs})
	}

	quality_control, err := controller.qualityControlService.Create(params)
	if err != nil {
		log.Error(err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.Status(fiber.StatusCreated).JSON(quality_control)
}
