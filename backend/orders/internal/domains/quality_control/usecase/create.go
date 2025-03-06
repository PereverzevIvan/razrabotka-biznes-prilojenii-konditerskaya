package quality_control_usecase

import (
	quality_control_params "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/domains/quality_control/params"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/models"
)

func (u *QualityControlUsecase) Create(
	params *quality_control_params.CreateParams,
) (*models.QualityControl, error) {
	quality_control := &models.QualityControl{
		MasterID:    params.MasterID,
		OrderID:     params.OrderID,
		IsSatisfies: params.IsSatisfies,
		Comment:     params.Comment,
		Parameter:   &params.Parameter,
	}

	return u.qualityControlRepo.Create(quality_control)
}
