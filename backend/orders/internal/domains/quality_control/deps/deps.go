package quality_control_deps

import (
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/configs"
	jwt_service "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/domains/jwt/service"
	quality_control_repo_mysql "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/domains/quality_control/repo/mysql"
	quality_control_usecase "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/domains/quality_control/usecase"
	"gorm.io/gorm"
)

type Deps struct {
	db        *gorm.DB
	jwtConfig *configs.JWTConfig

	qualityControlRepo    *quality_control_repo_mysql.QualityControlRepo
	qualityControlUsecase *quality_control_usecase.QualityControlUsecase

	jwtService *jwt_service.JWTService
}

func NewDeps(db *gorm.DB, jwtConfig *configs.JWTConfig) *Deps {
	return &Deps{
		db:        db,
		jwtConfig: jwtConfig,
	}
}

func (d *Deps) QualityControlRepo() *quality_control_repo_mysql.QualityControlRepo {
	if d.qualityControlRepo == nil {
		d.qualityControlRepo = quality_control_repo_mysql.NewQualityControlRepo(d.db)
	}
	return d.qualityControlRepo
}

func (d *Deps) QualityControlUsecase() *quality_control_usecase.QualityControlUsecase {
	if d.qualityControlUsecase == nil {
		d.qualityControlUsecase = quality_control_usecase.NewQualityControlService(d.QualityControlRepo())
	}
	return d.qualityControlUsecase
}

func (d *Deps) JwtService() *jwt_service.JWTService {
	if d.jwtService == nil {
		d.jwtService = jwt_service.NewJWTService(d.jwtConfig)
	}
	return d.jwtService
}
