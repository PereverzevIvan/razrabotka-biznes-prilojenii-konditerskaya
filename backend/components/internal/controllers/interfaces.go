package controllers

import (
	"mime/multipart"

	product_params "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/domains/product/params"
	purchased_component_params "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/domains/purchased_component/params"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models"
	results_purchased_component "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models/results/purchased_component"
	"github.com/gofiber/fiber/v3"
)

type IJWTService interface {
	GetRoleFromAccessTokenCookie(ctx fiber.Ctx) (string, error)
	GetUserIDFromAccessTokenCookie(ctx fiber.Ctx) (int, error)
}

type IComponentCategoryUsecase interface {
	GetByName(name string) (*models.ComponentCategory, error)
}

type IComponentTypeUsecase interface {
	GetAll(component_category_id int) ([]models.ComponentType, error)
}

type IComponentService interface {
	GetAll() ([]models.Component, error)
	GetByID(component_id int) (*models.Component, error)
}

type IPurchasedComponentService interface {
	GetAll(component_category_id int, params *purchased_component_params.GetAllParams) (*results_purchased_component.GetAllResults, error)
	DeleteEmpty(purchased_component_id int) error
	Create(purchased_component *models.PurchasedComponent) error
	Edit(purchased_component *models.PurchasedComponent) error
}

type IProductUsecase interface {
	GetAll() ([]models.Product, error)
	Create(params *product_params.CreateParams) (*models.Product, error)
	Update(params *product_params.UpdateParams) (*models.Product, error)
	// GetByID(id int) (*models.Product, error)
	GetByIDWithRecipe(id int) (*models.Product, error)

	CountNeededRecipeComponents(id int) (map[int]int, error)
	MakeProduct(id int) error

	CalcProductionPrice(product_id int) (
		total_price float64,
		purchased_components_prices map[int]float64,
		to_purchase_components_prices map[int]float64,
		err error,
	)

	CalcComponentsMaxDeliveryTime(component_ids []int) (int, error)

	CalcProductionMinTime(product *models.Product) (map[int][]models.ToolWorkInterval, error)

	SaveImage(ctx fiber.Ctx, id int, image *multipart.FileHeader) (string, error)
	// GetImage(product_id int) (string, error)
}
