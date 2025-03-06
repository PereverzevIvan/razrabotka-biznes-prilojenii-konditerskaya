package tool_type_deps

import (
	tool_type_repo_mysql "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/domains/tool_type/repo/mysql"
	tool_type_usecase "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/domains/tool_type/usecase"
	"gorm.io/gorm"
)

type Deps struct {
	db *gorm.DB

	toolTypeRepo    *tool_type_repo_mysql.ToolTypeRepo
	toolTypeUsecase *tool_type_usecase.ToolTypeUsecase
}

func NewDeps(db *gorm.DB) *Deps {
	return &Deps{
		db: db,
	}
}

func (d *Deps) ToolTypeRepo() *tool_type_repo_mysql.ToolTypeRepo {
	if d.toolTypeRepo == nil {
		d.toolTypeRepo = tool_type_repo_mysql.NewToolTypeRepo(d.db)
	}
	return d.toolTypeRepo
}

func (d *Deps) ToolTypeUsecase() *tool_type_usecase.ToolTypeUsecase {
	if d.toolTypeUsecase == nil {
		d.toolTypeUsecase = tool_type_usecase.NewToolTypeUsecase(d.ToolTypeRepo())
	}
	return d.toolTypeUsecase
}
