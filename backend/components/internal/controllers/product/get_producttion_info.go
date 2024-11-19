package controllers_product

import (
	"errors"

	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models/logic_errors"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

func (controller *productController) GetProductionInfo(ctx fiber.Ctx) error {
	product_id := fiber.Params[int](ctx, "product_id")
	if product_id <= 0 {
		return ctx.Status(fiber.StatusBadRequest).SendString("invalid product id")
	}

	count_needed_components, err := controller.productService.CountNeededRecipeComponents(product_id)
	if err != nil {
		log.Error(err)
		switch {
		case errors.Is(err, logic_errors.ErrCycleDetectedInProductRecipe):
			return ctx.Status(fiber.StatusConflict).JSON(fiber.Map{
				"error":   logic_errors.ErrCycleDetectedInProductRecipe.Error(),
				"product": count_needed_components,
			})
		}
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	total_price,
		purchased_components_prices,
		to_purchase_components_prices,
		err := controller.productService.CalcProductionPrice(product_id)
	if err != nil {
		log.Error(err)
		switch {
		case errors.Is(err, logic_errors.ErrNoSupplierForComponent):
			return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	components_to_purchase_ids := make([]int, 0, len(to_purchase_components_prices))
	for component_id := range to_purchase_components_prices {
		components_to_purchase_ids = append(components_to_purchase_ids, component_id)
	}

	max_deliver_time, err := controller.productService.CalcComponentsMaxDeliveryTime(components_to_purchase_ids)
	if err != nil {
		log.Error(err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"max_deliver_time":              max_deliver_time,
		"total_price":                   total_price,
		"count_needed_components":       count_needed_components,
		"purchased_components_prices":   purchased_components_prices,
		"to_purchase_components_prices": to_purchase_components_prices,
	})
}
