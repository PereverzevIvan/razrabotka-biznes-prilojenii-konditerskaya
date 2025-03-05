package tool_deps

import (
	tool_repo_mysql "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/domains/tool/repo/mysql"
	tool_usecase "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/domains/tool/usecase"
	"gorm.io/gorm"
)

type Deps struct {
	db *gorm.DB

	toolRepo    *tool_repo_mysql.ToolRepo
	toolUsecase *tool_usecase.ToolUsecase
}

func NewDeps(db *gorm.DB) *Deps {
	return &Deps{
		db: db,
	}
}

func (d *Deps) ToolRepo() *tool_repo_mysql.ToolRepo {
	if d.toolRepo == nil {
		d.toolRepo = tool_repo_mysql.NewToolRepo(d.db)
	}
	return d.toolRepo
}

func (d *Deps) ToolUsecase() *tool_usecase.ToolUsecase {
	if d.toolUsecase == nil {
		d.toolUsecase = tool_usecase.NewToolUsecase(d.ToolRepo())
	}
	return d.toolUsecase
}
