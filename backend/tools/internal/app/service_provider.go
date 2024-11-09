package app

import (
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/configs"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/controllers"
	repos_mysql_tool "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/repos/mysql/tool"
	repos_mysql_tool_type "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/repos/mysql/tool_type"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/services"
	services_jwt "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/services/jwt"
	services_tool "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/services/tool"
	services_tool_type "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/services/tool_type"

	"gorm.io/gorm"
)

type ServiceProvider struct {
	db *gorm.DB

	jwtConfig  *configs.JWTConfig
	jwtService controllers.IJWTService

	toolService controllers.IToolService
	toolRepo    services.IToolRepo

	toolTypeService controllers.IToolTypeService
	toolTypeRepo    services.IToolTypeRepo
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

func (s *ServiceProvider) ToolService() controllers.IToolService {
	if s.toolService == nil {
		s.toolService = services_tool.NewToolService(s.ToolRepo())
	}
	return s.toolService
}

func (s *ServiceProvider) ToolRepo() services.IToolRepo {
	if s.toolRepo == nil {
		s.toolRepo = repos_mysql_tool.NewToolRepo(s.db)
	}
	return s.toolRepo
}

func (s *ServiceProvider) ToolTypeService() controllers.IToolTypeService {
	if s.toolTypeService == nil {
		s.toolTypeService = services_tool_type.NewToolTypeService(s.ToolTypeRepo())
	}
	return s.toolTypeService
}

func (s *ServiceProvider) ToolTypeRepo() services.IToolTypeRepo {
	if s.toolTypeRepo == nil {
		s.toolTypeRepo = repos_mysql_tool_type.NewToolTypeRepo(s.db)
	}
	return s.toolTypeRepo
}
