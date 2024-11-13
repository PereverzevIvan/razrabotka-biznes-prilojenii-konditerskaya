package controllers_order

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

func (controller *orderController) GetAll(ctx fiber.Ctx) error {
	orders, err := controller.orderService.GetAll()
	if err != nil {
		log.Error(err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.Status(fiber.StatusOK).JSON(orders)
}
