package order_usecase

import "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/models"

func (service *OrderUsecase) GetByID(id int) (order *models.Order, err error) {
	return service.orderRepo.GetByID(id)
}
