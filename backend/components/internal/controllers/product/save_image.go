package controllers_product

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

func (controller *productController) SaveImage(ctx fiber.Ctx) error {
	product_id := fiber.Params[int](ctx, "product_id")
	if product_id <= 0 {
		return ctx.Status(fiber.StatusBadRequest).SendString("invalid product id")
	}

	file, err := ctx.FormFile("image")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	if file == nil {
		return ctx.Status(fiber.StatusBadRequest).SendString("No file uploaded")
	}

	res, err := controller.productService.SaveImage(ctx, product_id, file)
	if err != nil {
		log.Error(err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	if res == "" {
		return ctx.Status(fiber.StatusInternalServerError).SendString("failed to save image")
	}

	return ctx.Status(fiber.StatusOK).SendString(res)
}
