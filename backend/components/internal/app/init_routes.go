package app

import (
	component_type_controller "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/domains/component_type/controller"
	product_controller "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/domains/product/controller"
	purchased_component_controller "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/domains/purchased_component/controller"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/storage"
	"github.com/gofiber/fiber/v3"
)

type AddRouteFn func(fiber.Router, *storage.Storage)

func (app *App) initRoutes() error {
	// api
	api := app.fiberApp.Group("/api")

	addRoutes := []AddRouteFn{

		component_type_controller.AddRoutes,
		purchased_component_controller.AddRoutes,
		product_controller.AddRoutes,
	}

	for _, addRoutesFn := range addRoutes {
		addRoutesFn(api, app.Storage)
	}

	// controllers_component_type.AddComponentTypeControllerRoutes(
	// 	api,
	// 	app.serviceProvider.ComponentCategoryService(),
	// 	app.serviceProvider.ComponentTypeService(),
	// )

	// controllers_purchased_component.AddPurchasedComponentControllerRoutes(
	// 	api,
	// 	app.serviceProvider.ComponentCategoryService(),
	// 	app.serviceProvider.PurchasedComponentService(),
	// )

	// controllers_product.AddProductControllerRoutes(
	// 	api,
	// 	app.serviceProvider.ProductService(),
	// )
	return nil
}
