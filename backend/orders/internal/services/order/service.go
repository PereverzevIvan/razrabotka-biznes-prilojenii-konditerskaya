package services_order

import (
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/services"
)

type orderService struct {
	userRepo            services.IUserRepo
	orderRepo           services.IOrderRepo
	componentRepo       services.IComponentRepo
	qualityControlRepoo services.IQualityControlRepo
}

func NewOrderService(
	userRepo services.IUserRepo,
	orderRepo services.IOrderRepo,
	compoentRepo services.IComponentRepo,
	qualityControlRepo services.IQualityControlRepo,
) *orderService {

	service := &orderService{
		userRepo:            userRepo,
		orderRepo:           orderRepo,
		componentRepo:       compoentRepo,
		qualityControlRepoo: qualityControlRepo,
	}

	return service
}
