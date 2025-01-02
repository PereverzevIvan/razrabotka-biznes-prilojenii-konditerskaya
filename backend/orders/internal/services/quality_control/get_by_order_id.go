package services_quality_control

import "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/models"

func (service *qualityControlService) GetByOrderID(orderID int) ([]models.QualityControl, error) {
	return service.qualityControlRepo.GetByOrderID(orderID)
}
