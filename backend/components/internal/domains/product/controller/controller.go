package product_controller

import (
	"context"
	"mime/multipart"

	product_deps "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/domains/product/deps"
	product_params "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/domains/product/params"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/storage"
	"github.com/gofiber/fiber/v3"
)

type IProductUsecase interface {
	GetAll() ([]models.Product, error)
	Create(params *product_params.CreateParams) (*models.Product, error)
	Update(params *product_params.UpdateParams) (*models.Product, error)
	// GetByID(id int) (*models.Product, error)
	GetByIDWithRecipe(ctx context.Context, id int) (*models.Product, error)

	CountNeededRecipeComponents(id int) (map[int]int, error)
	MakeProduct(id int) error

	CalcProductionPrice(product_id int) (
		total_price float64,
		purchased_components_prices map[int]float64,
		to_purchase_components_prices map[int]float64,
		err error,
	)

	CalcComponentsMaxDeliveryTime(component_ids []int) (int, error)

	CalcProductionMinTime(ctx context.Context, product *models.Product) (map[int][]models.ToolWorkInterval, error)

	SaveImage(ctx fiber.Ctx, id int, image *multipart.FileHeader) (string, error)
	// GetImage(product_id int) (string, error)
}

type productController struct {
	productUsecase IProductUsecase
}

func AddRoutes(
	api fiber.Router,
	storage *storage.Storage,
) {
	deps := product_deps.NewDeps(storage.DB, storage.ToolClient, storage.ToolTypeClient)

	controller := productController{
		productUsecase: deps.ProductUsecase(),
	}

	apiProducts := api.Group("/products")

	apiProducts.Route("").
		Get(controller.GetAll).
		Post(controller.Create).
		Patch(controller.Update)

	apiProducts.Get("/:product_id", controller.GetByID)
	apiProducts.Get("/:product_id/production-info", controller.GetProductionInfo)
	apiProducts.Get("/:product_id/production-min-time", controller.ProductionMinTime)
	apiProducts.Post("/:product_id/make", controller.MakeProduct)
	apiProducts.Put("/:product_id/image", controller.SaveImage)
	apiProducts.Get("/:product_id/image", controller.GetImage)
}
