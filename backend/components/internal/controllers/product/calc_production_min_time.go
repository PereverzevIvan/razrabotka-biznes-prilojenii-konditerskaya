package controllers_product

import (
	"errors"

	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models/logic_errors"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

func (controller *productController) ProductionMinTime(ctx fiber.Ctx) error {
	product_id := fiber.Params[int](ctx, "product_id")
	if product_id <= 0 {
		return ctx.Status(fiber.StatusBadRequest).SendString("invalid product id")
	}

	product, err := controller.productService.GetByIDWithRecipe(product_id)
	if err != nil {
		log.Error(err)
		switch err {
		case logic_errors.ErrCycleDetectedInProductRecipe:
			return ctx.Status(fiber.StatusConflict).JSON(fiber.Map{
				"error":   logic_errors.ErrCycleDetectedInProductRecipe.Error(),
				"product": product,
			})
		}
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	result, err := controller.productService.CalcProductionMinTime(product)
	if err != nil {
		log.Error(err)
		switch {
		case errors.Is(err, logic_errors.ErrNoAvailableTools):
			return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
		case errors.Is(err, logic_errors.ErrNoRecipeOperationsInProduct):
			return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(result)
}
