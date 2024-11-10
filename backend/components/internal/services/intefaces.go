package services

import (
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models"
	results_purchased_component "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models/results/purchased_component"
)

type IComponentTypeRepo interface {
	GetAll(component_category_id int) ([]models.ComponentType, error)
}

type IComponentCategoryRepo interface {
	GetAll() ([]models.ComponentCategory, error)
	GetByName(name string) (*models.ComponentCategory, error)
}

type IPurchasedComponentRepo interface {
	GetAll() (*results_purchased_component.GetAllResults, error)
}
