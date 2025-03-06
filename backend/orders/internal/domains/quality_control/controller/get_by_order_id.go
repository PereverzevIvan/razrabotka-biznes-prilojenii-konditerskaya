package quality_control_controller

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

func (controller *qualityControllController) GetByOrderId(ctx fiber.Ctx) error {

	order_id := fiber.Params[int](ctx, "order_id")
	if order_id <= 0 {
		return ctx.Status(fiber.StatusBadRequest).SendString("invalid order id")
	}

	quality_controls, err := controller.qualityControlUsecase.GetByOrderID(order_id)
	if err != nil {
		log.Error(err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.Status(fiber.StatusOK).JSON(quality_controls)
}
