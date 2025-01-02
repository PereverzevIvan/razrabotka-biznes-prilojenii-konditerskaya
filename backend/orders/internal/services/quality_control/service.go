package services_quality_control

import "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/services"

type qualityControlService struct {
	qualityControlRepo services.IQualityControlRepo
}

func NewQualityControlService(
	qualityControlRepo services.IQualityControlRepo,
) *qualityControlService {
	return &qualityControlService{
		qualityControlRepo: qualityControlRepo,
	}
}
