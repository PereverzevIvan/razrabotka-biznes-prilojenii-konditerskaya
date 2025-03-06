package controllers_utils

import (
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/controllers"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

func GetComponentCategory(
	componentCategoryService controllers.IComponentCategoryUsecase,
	ctx fiber.Ctx,
) *models.ComponentCategory {
	category_name := ctx.Params("category_name")

	component_category, err := componentCategoryService.GetByName(category_name)
	if err != nil {
		log.Error(err)
		ctx.SendStatus(fiber.StatusInternalServerError)
		return nil
	}
	if component_category == nil {
		ctx.Status(fiber.StatusNotFound).SendString("component category '" + category_name + "' not found")
		return nil
	}

	return component_category
}
