package controllers_product

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
)

func (controller *productController) GetImage(ctx fiber.Ctx) error {
	product_id := fiber.Params[int](ctx, "product_id")
	if product_id <= 0 {
		return ctx.Status(fiber.StatusBadRequest).SendString("invalid product id")
	}

	// image, err := controller.productService.GetImage(product_id)
	// if err != nil {
	// 	log.Error(err)
	// 	return ctx.SendStatus(fiber.StatusInternalServerError)
	// }

	return ctx.Status(fiber.StatusOK).
		SendFile(fmt.Sprintf(
			"static/images/products/%v.png",
			product_id,
		),
		)
}
