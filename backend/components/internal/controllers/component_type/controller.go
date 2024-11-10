package controllers_component_type

import (
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/controllers"
	"github.com/gofiber/fiber/v3"
)

type componentTypeController struct {
	componentCategoriesMap   map[string]int
	componentCategoryService controllers.IComponentCategoryService
	componentTypeService     controllers.IComponentTypeService
}

func AddComponentTypeControllerRoutes(
	api fiber.Router,
	componentCategoryService controllers.IComponentCategoryService,
	componentTypeService controllers.IComponentTypeService,
) {
	controller := componentTypeController{
		componentCategoryService: componentCategoryService,
		componentTypeService:     componentTypeService,
	}

	api.Get(":category_name/types", controller.GetAll)
}
