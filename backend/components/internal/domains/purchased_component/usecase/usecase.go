package purchased_component_usecase

import (
	purchased_component_params "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/domains/purchased_component/params"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models"
)

type IPurchasedComponentRepo interface {
	GetAll(component_category_id int, params *purchased_component_params.GetAllParams) ([]models.PurchasedComponent, error)
	GetAllTotalRows(component_category_id int, params *purchased_component_params.GetAllParams) (int64, error)
	GetAllTotalCount(component_category_id int, params *purchased_component_params.GetAllParams) (int64, error)
	GetAllTotalPrice(component_category_id int, params *purchased_component_params.GetAllParams) (float64, error)

	DeleteEmpty(purchased_component_id int) error

	Create(purchased_component *models.PurchasedComponent) error

	Edit(purchased_component *models.PurchasedComponent) error

	CountRemainingComponents(component_id int) (int64, error)
	DeductComponents(component_id int, deduct_count int) error
	CalcPriceOfRequiredCount(component_id int, count int) (float64, error)
}

type PurchasedComponentUsecase struct {
	purchasedComponentRepo IPurchasedComponentRepo
}

func NewPurchasedComponentUsecase(purchasedComponentRepo IPurchasedComponentRepo) *PurchasedComponentUsecase {
	return &PurchasedComponentUsecase{
		purchasedComponentRepo: purchasedComponentRepo,
	}
}
