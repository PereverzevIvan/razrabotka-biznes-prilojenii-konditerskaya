package controllers_purchased_component

import (
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/controllers"
	"github.com/gofiber/fiber/v3"
)

type componentController struct {
	componentCategoryService  controllers.IComponentCategoryService
	purchasedComponentService controllers.IPurchasedComponentService
}

func AddPurchasedComponentControllerRoutes(
	api fiber.Router,
	componentCategoryService controllers.IComponentCategoryService,
	purchasedComponentService controllers.IPurchasedComponentService,
) {
	controller := componentController{
		componentCategoryService:  componentCategoryService,
		purchasedComponentService: purchasedComponentService,
	}

	api.Get(":category_name/purchased/", controller.GetAll)
	api.Delete(":category_name/purchased/:purchased_component_id", controller.DeleteEmpty)
	// api.Delete(":category_name/purchased/:component_id", controller.Delete)
}
