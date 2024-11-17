package controllers_product

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

func (controller *productController) GetByID(ctx fiber.Ctx) error {

	product_id := fiber.Params[int](ctx, "product_id")
	if product_id <= 0 {
		return ctx.Status(fiber.StatusBadRequest).SendString("invalid product id")
	}

	product, err := controller.productService.GetByID(product_id)
	if err != nil {
		log.Error(err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.Status(fiber.StatusOK).JSON(product)
}
