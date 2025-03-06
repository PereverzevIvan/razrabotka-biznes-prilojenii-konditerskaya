package order_usecase

import (
	"time"

	order_params "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/domains/order/params"
	quality_control_params "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/domains/quality_control/params"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/models"
)

type IUserRepo interface {
	GetByID(id int) (*models.User, error)
}

type IOrderRepo interface {
	GetAll(params *order_params.RepoGetAllParams) ([]models.Order, error)
	GetByID(id int) (*models.Order, error)
	Create(order *models.Order) (*models.Order, error)
	CountForDate(date time.Time) (int64, error)
	Update(order *models.Order) (*models.Order, error)
}

type IComponentRepo interface {
	ConsumeProductComponents(product_id int) error
}

type IQualityControlRepo interface {
	Create(quality_control *models.QualityControl) (*models.QualityControl, error)
	Update(params *quality_control_params.UpdateParams) (*models.QualityControl, error)
	GetByOrderID(order_id int) ([]models.QualityControl, error)
	IsOrderValid(order_id int) (bool, error)
}

type OrderUsecase struct {
	userRepo           IUserRepo
	orderRepo          IOrderRepo
	componentRepo      IComponentRepo
	qualityControlRepo IQualityControlRepo
}

func NewOrderUsecase(
	userRepo IUserRepo,
	orderRepo IOrderRepo,
	componentRepo IComponentRepo,
	qualityControlRepo IQualityControlRepo,
) *OrderUsecase {

	service := &OrderUsecase{
		userRepo:           userRepo,
		orderRepo:          orderRepo,
		componentRepo:      componentRepo,
		qualityControlRepo: qualityControlRepo,
	}

	return service
}
