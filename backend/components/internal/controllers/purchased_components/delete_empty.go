package controllers_purchased_component

import (
	controllers_utils "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/controllers/utils"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models/logic_errors"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

func (controller *purchasedComponentController) DeleteEmpty(ctx fiber.Ctx) error {
	component_category := controllers_utils.GetComponentCategory(controller.componentCategoryService, ctx)
	if component_category == nil {
		return nil
	}

	purchased_component_id := fiber.Params[int](ctx, "purchased_component_id")
	if purchased_component_id <= 0 {
		return ctx.Status(fiber.StatusBadRequest).SendString("invalid purchased component id")
	}

	err := controller.purchasedComponentService.DeleteEmpty(purchased_component_id)
	if err != nil {
		log.Error(err)
		switch err {
		case logic_errors.ErrDeleteNonEmptyQuantity:
			return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.SendStatus(fiber.StatusOK)
}
