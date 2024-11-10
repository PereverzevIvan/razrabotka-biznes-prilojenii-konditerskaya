package app

import (
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/configs"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/controllers"
	repos_mysql_component_category "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/repos/mysql/component_category"
	repos_mysql_component_type "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/repos/mysql/component_type"
	repos_mysql_purchased_component "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/repos/mysql/purchased_component"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/services"
	services_component_category "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/services/component_category"
	services_component_type "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/services/component_type"
	services_jwt "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/services/jwt"
	services_purchased_component "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/services/purchased_component"

	"gorm.io/gorm"
)

type ServiceProvider struct {
	db *gorm.DB

	jwtConfig  *configs.JWTConfig
	jwtService controllers.IJWTService

	componentCategoryService controllers.IComponentCategoryService
	componentCategoryRepo    services.IComponentCategoryRepo

	componentTypeService controllers.IComponentTypeService
	componentTypeRepo    services.IComponentTypeRepo

	purchasedComponentService controllers.IPurchasedComponentService
	purchasedComponentRepo    services.IPurchasedComponentRepo
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

func (s *ServiceProvider) ComponentCategoryService() controllers.IComponentCategoryService {
	if s.componentCategoryService == nil {
		s.componentCategoryService = services_component_category.NewComponentCategoryService(s.ComponentCategoryRepo())
	}
	return s.componentCategoryService
}

func (s *ServiceProvider) ComponentCategoryRepo() services.IComponentCategoryRepo {
	if s.componentCategoryRepo == nil {
		s.componentCategoryRepo = repos_mysql_component_category.NewComponentCategoryRepo(s.db)
	}
	return s.componentCategoryRepo
}

func (s *ServiceProvider) ComponentTypeService() controllers.IComponentTypeService {
	if s.componentTypeService == nil {
		s.componentTypeService = services_component_type.NewComponentTypeService(s.ComponentTypeRepo())
	}
	return s.componentTypeService
}

func (s *ServiceProvider) ComponentTypeRepo() services.IComponentTypeRepo {
	if s.componentTypeRepo == nil {
		s.componentTypeRepo = repos_mysql_component_type.NewComponentTypeRepo(s.db)
	}
	return s.componentTypeRepo
}

func (s *ServiceProvider) PurchasedComponentService() controllers.IPurchasedComponentService {
	if s.purchasedComponentService == nil {
		s.purchasedComponentService = services_purchased_component.NewPurchasedComponentService(s.PurchasedComponentRepo())
	}
	return s.purchasedComponentService
}

func (s *ServiceProvider) PurchasedComponentRepo() services.IPurchasedComponentRepo {
	if s.purchasedComponentRepo == nil {
		s.purchasedComponentRepo = repos_mysql_purchased_component.NewPurchasedComponentService(s.db)
	}
	return s.purchasedComponentRepo
}
