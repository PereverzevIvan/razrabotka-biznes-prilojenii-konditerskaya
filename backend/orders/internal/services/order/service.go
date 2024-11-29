package services_order

import (
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/services"
)

type orderService struct {
	userRepo  services.IUserRepo
	orderRepo services.IOrderRepo
}

func NewOrderService(
	userRepo services.IUserRepo,
	orderRepo services.IOrderRepo,
) *orderService {

	service := &orderService{
		userRepo:  userRepo,
		orderRepo: orderRepo,
	}

	return service
}
