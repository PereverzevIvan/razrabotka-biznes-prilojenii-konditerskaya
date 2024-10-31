package app

import (
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/gateway/internal/controllers"
	mysql_repo "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/gateway/internal/repositories/mysql"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/gateway/internal/services"
	"gorm.io/gorm"
)

type ServiceProvider struct {
	db          *gorm.DB
	userService controllers.IUserService
	userRepo    services.IUserRepo
	// ...
}

func newServiceProvider(db *gorm.DB) *ServiceProvider {
	return &ServiceProvider{
		db: db,
	}
}

func (s *ServiceProvider) UserService() controllers.IUserService {
	if s.userService == nil {
		s.userService = services.NewUserService(s.UserRepo())
	}
	return s.userService
}

func (s *ServiceProvider) UserRepo() services.IUserRepo {
	if s.userRepo == nil {
		s.userRepo = mysql_repo.NewUserRepo(s.db)
	}
	return s.userRepo
}
