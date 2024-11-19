package controllers_product

import (
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/controllers"
	"github.com/gofiber/fiber/v3"
)

type productController struct {
	productService controllers.IProductService
}

func AddProductControllerRoutes(
	api fiber.Router,
	productService controllers.IProductService,
) {
	controller := productController{
		productService: productService,
	}

	api.Get("/products", controller.GetAll)
	api.Get("/products/:product_id", controller.GetByID)
	api.Get("/products/:product_id/production-info", controller.GetProductionInfo)
	api.Post("/products/:product_id/make", controller.MakeProduct)
}
