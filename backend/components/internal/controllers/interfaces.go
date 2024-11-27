package controllers

import (
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models"
	params_product "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models/params/product"
	params_purchased_component "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models/params/purchased_component"
	results_purchased_component "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models/results/purchased_component"
	"github.com/gofiber/fiber/v3"
)

type IJWTService interface {
	GetRoleFromAccessTokenCookie(ctx fiber.Ctx) (string, error)
	GetUserIDFromAccessTokenCookie(ctx fiber.Ctx) (int, error)
}

type IComponentCategoryService interface {
	GetByName(name string) (*models.ComponentCategory, error)
}

type IComponentTypeService interface {
	GetAll(component_category_id int) ([]models.ComponentType, error)
}

type IComponentService interface {
	GetAll() ([]models.Component, error)
	GetByID(component_id int) (*models.Component, error)
}

type IPurchasedComponentService interface {
	GetAll(component_category_id int, params *params_purchased_component.GetAllParams) (*results_purchased_component.GetAllResults, error)
	DeleteEmpty(purchased_component_id int) error
	Create(purchased_component *models.PurchasedComponent) error
	Edit(purchased_component *models.PurchasedComponent) error
}

type IProductService interface {
	GetAll() ([]models.Product, error)
	Create(params *params_product.CreateParams) (*models.Product, error)
	Update(params *params_product.UpdateParams) (*models.Product, error)
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
}
