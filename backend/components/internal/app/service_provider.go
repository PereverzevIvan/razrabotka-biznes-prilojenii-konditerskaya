package app

import (
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/configs"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/controllers"
	repos_mysql_component "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/repos/mysql/component"
	repos_mysql_component_category "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/repos/mysql/component_category"
	repos_mysql_component_type "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/repos/mysql/component_type"
	repos_mysql_product "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/repos/mysql/product"
	repos_mysql_purchased_component "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/repos/mysql/purchased_component"
	repos_mysql_supplier_component "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/repos/mysql/supplier_component"
	repos_mysql_tool "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/repos/mysql/tool"
	repos_mysql_tool_type "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/repos/mysql/tool_type"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/services"
	services_component_category "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/services/component_category"
	services_component_type "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/services/component_type"
	services_jwt "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/services/jwt"
	services_product "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/services/product"
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

	componentRepo services.IComponentRepo

	purchasedComponentService controllers.IPurchasedComponentService
	purchasedComponentRepo    services.IPurchasedComponentRepo

	supplierComponentRepo services.ISupplierComponentRepo

	toolTypeRepo services.IToolTypeRepo
	toolRepo     services.IToolRepo

	productService controllers.IProductService
	productRepo    services.IProductRepo
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

func (s *ServiceProvider) ComponentRepo() services.IComponentRepo {
	if s.componentRepo == nil {
		s.componentRepo = repos_mysql_component.NewComponentRepo(s.db)
	}
	return s.componentRepo
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

func (s *ServiceProvider) SupplierComponentRepo() services.ISupplierComponentRepo {
	if s.supplierComponentRepo == nil {
		s.supplierComponentRepo = repos_mysql_supplier_component.NewSupplierComponentRepo(s.db)
	}
	return s.supplierComponentRepo
}

func (s *ServiceProvider) ToolTypeRepo() services.IToolTypeRepo {
	if s.toolTypeRepo == nil {
		s.toolTypeRepo = repos_mysql_tool_type.NewToolTypeRepo(s.db)
	}
	return s.toolTypeRepo
}

func (s *ServiceProvider) ToolRepo() services.IToolRepo {
	if s.toolRepo == nil {
		s.toolRepo = repos_mysql_tool.NewToolRepo(s.db)
	}
	return s.toolRepo
}

func (s *ServiceProvider) ProductService() controllers.IProductService {
	if s.productService == nil {
		s.productService = services_product.NewProductService(
			s.ProductRepo(),
			s.ComponentRepo(),
			s.PurchasedComponentRepo(),
			s.SupplierComponentRepo(),
			s.ToolTypeRepo(),
			s.ToolRepo(),
		)
	}
	return s.productService
}

func (s *ServiceProvider) ProductRepo() services.IProductRepo {
	if s.productRepo == nil {
		s.productRepo = repos_mysql_product.NewProductRepo(s.db)
	}
	return s.productRepo
}
