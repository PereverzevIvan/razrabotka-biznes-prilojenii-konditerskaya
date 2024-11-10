package services

import (
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models"
	params_purchased_component "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models/params/purchased_component"
)

type IComponentTypeRepo interface {
	GetAll(component_category_id int) ([]models.ComponentType, error)
}

type IComponentCategoryRepo interface {
	GetAll() ([]models.ComponentCategory, error)
	GetByName(name string) (*models.ComponentCategory, error)
}

type IPurchasedComponentRepo interface {
	GetAll(component_category_id int, params *params_purchased_component.GetAllParams) ([]models.PurchasedComponent, error)
	GetAllTotalRows(component_category_id int, params *params_purchased_component.GetAllParams) (int64, error)
	GetAllTotalCount(component_category_id int, params *params_purchased_component.GetAllParams) (int64, error)
	GetAllTotalPrice(component_category_id int, params *params_purchased_component.GetAllParams) (float64, error)

	DeleteEmpty(purchased_component_id int) error

	Create(purchased_component *models.PurchasedComponent) error
}
