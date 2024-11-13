package services_order

import (
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/services"
)

type orderService struct {
	orderRepo services.IOrderRepo
}

func NewOrderService(repo services.IOrderRepo) *orderService {
	service := &orderService{
		orderRepo: repo,
	}

	return service
}
