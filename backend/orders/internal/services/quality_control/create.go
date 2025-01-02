package services_quality_control

import (
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/models"
	params_quality_control "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/models/params/quality_control"
)

func (service *qualityControlService) Create(
	params *params_quality_control.CreateParams,
) (*models.QualityControl, error) {
	quality_control := &models.QualityControl{
		MasterID:    params.MasterID,
		OrderID:     params.OrderID,
		IsSatisfies: params.IsSatisfies,
		Comment:     params.Comment,
		Parameter:   &params.Parameter,
	}

	return service.qualityControlRepo.Create(quality_control)
}
