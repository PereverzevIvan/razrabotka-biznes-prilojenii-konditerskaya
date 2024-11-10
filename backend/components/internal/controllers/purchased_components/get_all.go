package controllers_purchased_component

import (
	controllers_utils "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/controllers/utils"
	params_purchased_component "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models/params/purchased_component"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

func (controller *purchasedComponentController) GetAll(ctx fiber.Ctx) error {
	component_category := controllers_utils.GetComponentCategory(controller.componentCategoryService, ctx)
	if component_category == nil {
		return nil
	}

	params := &params_purchased_component.GetAllParams{}
	err := ctx.Bind().Query(params)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	if errs := params.Validate(); len(errs) != 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"errors": errs})
	}

	purchased_components_result, err := controller.purchasedComponentService.GetAll(
		component_category.ID,
		params,
	)
	if err != nil {
		log.Error(err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.JSON(purchased_components_result)
}
