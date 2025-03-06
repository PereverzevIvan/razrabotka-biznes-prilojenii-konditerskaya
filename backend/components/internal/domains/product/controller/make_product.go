package product_controller

import (
	"errors"

	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models/logic_errors"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

func (c *productController) MakeProduct(ctx fiber.Ctx) error {

	product_id := fiber.Params[int](ctx, "product_id")
	if product_id <= 0 {
		return ctx.Status(fiber.StatusBadRequest).SendString("invalid product id")
	}

	res, err := c.productUsecase.CountNeededRecipeComponents(product_id)
	if err != nil {
		log.Error(err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	err = c.productUsecase.MakeProduct(product_id)
	if err != nil {
		log.Error(err)
		// log.Info("check: ", errors.As(err, &logic_errors.ErrCycleDetectedInProductRecipe))
		switch {
		case errors.Is(err, logic_errors.ErrNotEnoughComponentsToMakeProduct):
			return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.Status(fiber.StatusOK).JSON(res)
}
