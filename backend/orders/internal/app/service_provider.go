package app

import (
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/configs"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/controllers"
	repos_mysql_order "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/repos/mysql/order"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/services"
	services_jwt "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/services/jwt"
	services_order "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/services/order"

	"gorm.io/gorm"
)

type ServiceProvider struct {
	db *gorm.DB

	jwtConfig  *configs.JWTConfig
	jwtService controllers.IJWTService

	orderService controllers.IOrderService
	orderRepo    services.IOrderRepo

	// ...
}

func newServiceProvider(db *gorm.DB, jwtConfig *configs.JWTConfig) *ServiceProvider {
	return &ServiceProvider{
		db: db,

		jwtConfig: jwtConfig,
	}
}

func (s *ServiceProvider) JWTService() controllers.IJWTService {
	if s.jwtService == nil {
		s.jwtService = services_jwt.NewJWTService(s.jwtConfig)
	}

	return s.jwtService
}

func (s *ServiceProvider) OrderService() controllers.IOrderService {
	if s.orderService == nil {
		s.orderService = services_order.NewOrderService(s.OrderRepo())
	}
	return s.orderService
}

func (s *ServiceProvider) OrderRepo() services.IOrderRepo {
	if s.orderRepo == nil {
		s.orderRepo = repos_mysql_order.NewOrderRepo(s.db)
	}
	return s.orderRepo
}
