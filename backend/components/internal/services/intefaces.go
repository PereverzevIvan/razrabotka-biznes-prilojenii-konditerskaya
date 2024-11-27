package services

import (
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models"
	params_purchased_component "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models/params/purchased_component"
	params_tool "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models/params/tool"
)

type IComponentTypeRepo interface {
	GetAll(component_category_id int) ([]models.ComponentType, error)
}

type IComponentCategoryRepo interface {
	GetAll() ([]models.ComponentCategory, error)
	GetByName(name string) (*models.ComponentCategory, error)
}

type IComponentRepo interface {
	// GetAll(component_category_id int) ([]models.Component, error)
	GetByID(id int) (*models.Component, error)
}

type IPurchasedComponentRepo interface {
	GetAll(component_category_id int, params *params_purchased_component.GetAllParams) ([]models.PurchasedComponent, error)
	GetAllTotalRows(component_category_id int, params *params_purchased_component.GetAllParams) (int64, error)
	GetAllTotalCount(component_category_id int, params *params_purchased_component.GetAllParams) (int64, error)
	GetAllTotalPrice(component_category_id int, params *params_purchased_component.GetAllParams) (float64, error)

	DeleteEmpty(purchased_component_id int) error

	Create(purchased_component *models.PurchasedComponent) error

	Edit(purchased_component *models.PurchasedComponent) error

	CountRemainingComponents(component_id int) (int64, error)
	DeductComponents(component_id int, deduct_count int) error
	CalcPriceOfRequiredCount(component_id int, count int) (float64, error)
}
type ISupplierComponentRepo interface {
	FastestDeliveryComponentPrice(component_id int) (float64, error)
	FastestDeliveryComponentTime(component_id int) (int, error)
}

type IToolTypeRepo interface {
	GetByID(id int) (*models.ToolType, error)
}

type IToolRepo interface {
	GetAll(params *params_tool.GetAllParams) ([]models.Tool, error)
}

type IProductRepo interface {
	GetAll() ([]models.Product, error)
	Create(product *models.Product) (*models.Product, error)
	GetByID(id int) (*models.Product, error)
}
