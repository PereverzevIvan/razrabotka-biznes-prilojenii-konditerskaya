package quality_control_usecase

import (
	quality_control_params "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/domains/quality_control/params"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/models"
)

func (u *QualityControlUsecase) Update(
	params *quality_control_params.UpdateParams,
) (*models.QualityControl, error) {
	return u.qualityControlRepo.Update(params)
}
