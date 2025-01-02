package services_quality_control

import (
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/models"
	params_quality_control "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/models/params/quality_control"
)

func (service *qualityControlService) Update(
	params *params_quality_control.UpdateParams,
) (*models.QualityControl, error) {
	return service.qualityControlRepo.Update(params)
}
