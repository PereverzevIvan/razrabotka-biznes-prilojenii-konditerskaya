package quality_control_usecase

import (
	quality_control_params "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/domains/quality_control/params"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/models"
)

type IQualityControlRepo interface {
	Create(quality_control *models.QualityControl) (*models.QualityControl, error)
	Update(params *quality_control_params.UpdateParams) (*models.QualityControl, error)
	GetByOrderID(order_id int) ([]models.QualityControl, error)
	IsOrderValid(order_id int) (bool, error)
}

type QualityControlUsecase struct {
	qualityControlRepo IQualityControlRepo
}

func NewQualityControlService(
	qualityControlRepo IQualityControlRepo,
) *QualityControlUsecase {
	return &QualityControlUsecase{
		qualityControlRepo: qualityControlRepo,
	}
}
