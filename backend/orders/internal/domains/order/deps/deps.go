package order_deps

import (
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/configs"
	component_repo_rest "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/domains/component/repo/rest"
	jwt_service "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/domains/jwt/service"
	order_repo_mysql "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/domains/order/repo/mysql"
	order_usecase "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/domains/order/usecase"
	quality_control_repo_mysql "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/domains/quality_control/repo/mysql"
	user_repo_mysql "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/domains/user/repo/mysql"
	"gorm.io/gorm"
)

type Deps struct {
	db                *gorm.DB
	jwtConfig         *configs.JWTConfig
	componenetBaseUrl string

	jwtService *jwt_service.JWTService

	userRepo           *user_repo_mysql.UserRepo
	orderRepo          *order_repo_mysql.OrderRepo
	componentRepo      *component_repo_rest.ComponentRepo
	qualityControlRepo *quality_control_repo_mysql.QualityControlRepo

	orderUsecase *order_usecase.OrderUsecase
}

func NewDeps(db *gorm.DB, jwtConfig *configs.JWTConfig, componenetBaseUrl string) *Deps {
	return &Deps{
		db:                db,
		jwtConfig:         jwtConfig,
		componenetBaseUrl: componenetBaseUrl,
	}
}

func (d *Deps) JwtService() *jwt_service.JWTService {
	if d.jwtService == nil {
		d.jwtService = jwt_service.NewJWTService(d.jwtConfig)
	}
	return d.jwtService
}

func (d *Deps) UserRepo() *user_repo_mysql.UserRepo {
	if d.userRepo == nil {
		d.userRepo = user_repo_mysql.NewUserRepo(d.db)
	}
	return d.userRepo
}

func (d *Deps) OrderRepo() *order_repo_mysql.OrderRepo {
	if d.orderRepo == nil {
		d.orderRepo = order_repo_mysql.NewOrderRepo(d.db)
	}
	return d.orderRepo
}

func (d *Deps) ComponentRepo() *component_repo_rest.ComponentRepo {
	if d.componentRepo == nil {
		d.componentRepo = component_repo_rest.NewComponentRepo(d.componenetBaseUrl)
	}
	return d.componentRepo
}

func (d *Deps) QualityControlRepo() *quality_control_repo_mysql.QualityControlRepo {
	if d.qualityControlRepo == nil {
		d.qualityControlRepo = quality_control_repo_mysql.NewQualityControlRepo(d.db)
	}
	return d.qualityControlRepo
}

func (d *Deps) OrderUsecase() *order_usecase.OrderUsecase {
	if d.orderUsecase == nil {
		d.orderUsecase = order_usecase.NewOrderUsecase(
			d.UserRepo(), d.OrderRepo(), d.ComponentRepo(), d.QualityControlRepo())
	}
	return d.orderUsecase
}
