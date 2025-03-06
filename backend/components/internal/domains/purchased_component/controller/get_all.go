package purchased_component_controller

import (
	controllers_utils "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/controllers/utils"
	purchased_component_params "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/domains/purchased_component/params"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

func (c *purchasedComponentController) GetAll(ctx fiber.Ctx) error {
	component_category := controllers_utils.GetComponentCategory(c.componentCategoryService, ctx)
	if component_category == nil {
		return nil
	}

	params := &purchased_component_params.GetAllParams{}
	err := ctx.Bind().Query(params)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	if errs := params.Validate(); len(errs) != 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"errors": errs})
	}

	purchased_components_result, err := c.purchasedComponentService.GetAll(
		component_category.ID,
		params,
	)
	if err != nil {
		log.Error(err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.JSON(purchased_components_result)
}
