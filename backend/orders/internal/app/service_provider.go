package app

import (
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/configs"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/controllers"
	repos_mysql_order "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/repos/mysql/order"
	repos_mysql_quality_control "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/repos/mysql/quality_control"
	repos_mysql_user "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/repos/mysql/user"
	repos_rest_component "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/repos/rest/component"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/services"
	services_jwt "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/services/jwt"
	services_order "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/services/order"
	services_quality_control "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/services/quality_control"

	"gorm.io/gorm"
)

type ServiceProvider struct {
	db *gorm.DB

	jwtConfig  *configs.JWTConfig
	jwtService controllers.IJWTService

	userRepo services.IUserRepo

	componentsAPIConfig *configs.ComponentsAPIConfig
	componentRepo       services.IComponentRepo

	qualityControlRepo    services.IQualityControlRepo
	qualityControlService controllers.IQualityControlService

	orderService controllers.IOrderService
	orderRepo    services.IOrderRepo

	// ...
}

func newServiceProvider(
	db *gorm.DB,
	jwtConfig *configs.JWTConfig,
	componentsAPIconfig *configs.ComponentsAPIConfig,
) *ServiceProvider {
	return &ServiceProvider{
		db:                  db,
		jwtConfig:           jwtConfig,
		componentsAPIConfig: componentsAPIconfig,
	}
}

func (s *ServiceProvider) JWTService() controllers.IJWTService {
	if s.jwtService == nil {
		s.jwtService = services_jwt.NewJWTService(s.jwtConfig)
	}

	return s.jwtService
}

func (s *ServiceProvider) UserRepo() services.IUserRepo {
	if s.userRepo == nil {
		s.userRepo = repos_mysql_user.NewUserRepo(s.db)
	}
	return s.userRepo
}

func (s *ServiceProvider) ComponentRepo() services.IComponentRepo {
	if s.componentRepo == nil {
		s.componentRepo = repos_rest_component.NewComponentRepo(
			s.componentsAPIConfig.BaseUrl,
		)
	}
	return s.componentRepo
}

func (s *ServiceProvider) QualityControlRepo() services.IQualityControlRepo {
	if s.qualityControlRepo == nil {
		s.qualityControlRepo = repos_mysql_quality_control.NewQualityControlRepo(s.db)
	}
	return s.qualityControlRepo
}

func (s *ServiceProvider) QualityControlService() controllers.IQualityControlService {
	if s.qualityControlService == nil {
		s.qualityControlService = services_quality_control.NewQualityControlService(
			s.QualityControlRepo(),
		)
	}
	return s.qualityControlService
}

func (s *ServiceProvider) OrderService() controllers.IOrderService {
	if s.orderService == nil {
		s.orderService = services_order.NewOrderService(
			s.UserRepo(),
			s.OrderRepo(),
			s.ComponentRepo(),
			s.QualityControlRepo(),
		)
	}
	return s.orderService
}

func (s *ServiceProvider) OrderRepo() services.IOrderRepo {
	if s.orderRepo == nil {
		s.orderRepo = repos_mysql_order.NewOrderRepo(s.db)
	}
	return s.orderRepo
}
