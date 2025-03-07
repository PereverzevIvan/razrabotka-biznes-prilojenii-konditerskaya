package product_controller

import (
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models/logic_errors"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

func (c *productController) GetByID(ctx fiber.Ctx) error {

	product_id := fiber.Params[int](ctx, "product_id")
	if product_id <= 0 {
		return ctx.Status(fiber.StatusBadRequest).SendString("invalid product id")
	}

	product, err := c.productUsecase.GetByIDWithRecipe(ctx.Context(), product_id)
	if err != nil {
		switch err {
		case logic_errors.ErrCycleDetectedInProductRecipe:
			return ctx.Status(fiber.StatusConflict).JSON(fiber.Map{
				"error":   logic_errors.ErrCycleDetectedInProductRecipe.Error(),
				"product": product,
			})
		}
		log.Error(err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.Status(fiber.StatusOK).JSON(product)
}
