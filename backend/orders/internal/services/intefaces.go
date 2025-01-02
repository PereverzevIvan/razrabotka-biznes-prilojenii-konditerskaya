package services

import (
	"time"

	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/models"
	params_order "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/models/params/order"
	params_quality_control "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/models/params/quality_control"
)

type IUserRepo interface {
	GetByID(id int) (*models.User, error)
}

type IOrderRepo interface {
	GetAll(params *params_order.RepoGetAllParams) ([]models.Order, error)
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
	Update(params *params_quality_control.UpdateParams) (*models.QualityControl, error)
	GetByOrderID(order_id int) ([]models.QualityControl, error)
	IsOrderValid(order_id int) (bool, error)
}
