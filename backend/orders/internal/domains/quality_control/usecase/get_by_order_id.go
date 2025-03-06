package quality_control_usecase

import "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/models"

func (u *QualityControlUsecase) GetByOrderID(orderID int) ([]models.QualityControl, error) {
	return u.qualityControlRepo.GetByOrderID(orderID)
}
