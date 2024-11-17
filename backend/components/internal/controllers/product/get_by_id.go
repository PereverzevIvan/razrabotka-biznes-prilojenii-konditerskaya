package controllers_product

import (
	services_product "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/services/product"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

func (controller *productController) GetByID(ctx fiber.Ctx) error {

	product_id := fiber.Params[int](ctx, "product_id")
	if product_id <= 0 {
		return ctx.Status(fiber.StatusBadRequest).SendString("invalid product id")
	}

	product, err := controller.productService.GetByIDWithRecipe(product_id)
	if err != nil {
		switch err {
		case services_product.ErrCycleDetectedInProductRecipe:
			return ctx.Status(fiber.StatusConflict).JSON(fiber.Map{
				"error":   services_product.ErrCycleDetectedInProductRecipe.Error(),
				"product": product,
			})
		}
		log.Error(err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.Status(fiber.StatusOK).JSON(product)
}
