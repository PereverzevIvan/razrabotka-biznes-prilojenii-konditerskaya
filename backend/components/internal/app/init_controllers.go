package app

import (
	controllers_component_type "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/controllers/component_type"
	controllers_product "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/controllers/product"
	controllers_purchased_component "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/controllers/purchased_components"
)

func (app *App) initControllers() error {
	// api
	api := app.fiberApp.Group("/api")

	controllers_component_type.AddComponentTypeControllerRoutes(
		api,
		app.serviceProvider.ComponentCategoryService(),
		app.serviceProvider.ComponentTypeService(),
	)

	controllers_purchased_component.AddPurchasedComponentControllerRoutes(
		api,
		app.serviceProvider.ComponentCategoryService(),
		app.serviceProvider.PurchasedComponentService(),
	)

	controllers_product.AddProductControllerRoutes(
		api,
		app.serviceProvider.ProductService(),
	)
	return nil
}
