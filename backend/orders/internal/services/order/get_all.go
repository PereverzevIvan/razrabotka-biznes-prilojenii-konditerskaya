package services_order

import "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/models"

func (service *orderService) GetAll() ([]models.Order, error) {
	return service.orderRepo.GetAll()
}
