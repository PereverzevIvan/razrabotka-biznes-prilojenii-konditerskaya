package tool_failure_deps

import (
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/configs"
	tool_failure_repo_mysql "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/domains/tool_failure/repo/mysql"
	tool_failure_usecase "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/domains/tool_failure/usecase"
	"gorm.io/gorm"
)

type Deps struct {
	db        *gorm.DB
	JwtConfig *configs.JWTConfig

	toolFailureRepo    *tool_failure_repo_mysql.ToolFailureRepo
	toolFailureUsecase *tool_failure_usecase.ToolFailureUsecase
}

func NewDeps(db *gorm.DB, jwtConfig *configs.JWTConfig) *Deps {
	return &Deps{
		db:        db,
		JwtConfig: jwtConfig,
	}
}

func (d *Deps) ToolFailureRepo() *tool_failure_repo_mysql.ToolFailureRepo {
	if d.toolFailureRepo == nil {
		d.toolFailureRepo = tool_failure_repo_mysql.NewToolFailureRepo(d.db)
	}
	return d.toolFailureRepo
}

func (d *Deps) ToolFailureUsecase() *tool_failure_usecase.ToolFailureUsecase {
	if d.toolFailureUsecase == nil {
		d.toolFailureUsecase = tool_failure_usecase.NewToolFailureUsecase(d.ToolFailureRepo())
	}
	return d.toolFailureUsecase
}
