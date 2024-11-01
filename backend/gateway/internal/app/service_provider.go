package app

import (
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/gateway/configs"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/gateway/internal/controllers"
	middlewares_auth "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/gateway/internal/middlewares/auth"
	repo_mysql_user "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/gateway/internal/repositories/repo_mysql/user"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/gateway/internal/services"

	services_jwt "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/gateway/internal/services/jwt"
	services_user "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/gateway/internal/services/user"

	"gorm.io/gorm"
)

type ServiceProvider struct {
	db *gorm.DB

	jwtConfig  *configs.JWTConfig
	jwtService controllers.IJWTService

	userService controllers.IUserService
	userRepo    services.IUserRepo
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

func (s *ServiceProvider) AuthMiddleware() controllers.IAuthMiddleware {
	return middlewares_auth.NewAuthMiddleware(s.JWTService())
}

func (s *ServiceProvider) UserService() controllers.IUserService {
	if s.userService == nil {
		s.userService = services_user.NewUserService(s.UserRepo())
	}
	return s.userService
}

func (s *ServiceProvider) UserRepo() services.IUserRepo {
	if s.userRepo == nil {
		s.userRepo = repo_mysql_user.NewUserRepo(s.db)
	}
	return s.userRepo
}
