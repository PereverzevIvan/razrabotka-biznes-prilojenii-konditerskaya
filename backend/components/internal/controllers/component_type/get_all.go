package controllers_component_type

import (
	controllers_utils "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/controllers/utils"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

func (controller *componentTypeController) GetAll(ctx fiber.Ctx) error {
	component_category := controllers_utils.GetComponentCategory(controller.componentCategoryService, ctx)
	if component_category == nil {
		return nil
	}

	component_types, err := controller.componentTypeService.GetAll(component_category.ID)
	if err != nil {
		log.Error(err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.JSON(component_types)
}
