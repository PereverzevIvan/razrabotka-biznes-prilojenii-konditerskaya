package component_type_deps

import (
	component_category_repo_mysql "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/domains/component_category/repo/mysql"
	component_category_usecase "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/domains/component_category/usecase"
	component_type_repo_mysql "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/domains/component_type/repo/mysql"
	component_type_usecase "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/domains/component_type/usecase"
	"gorm.io/gorm"
)

type Deps struct {
	db *gorm.DB

	componentTypeRepo    *component_type_repo_mysql.ComponentTypeRepo
	componentTypeUsecase *component_type_usecase.ComponentTypeUsecase

	componentCategoryRepo    *component_category_repo_mysql.ComponentCategoryRepo
	componentCategoryUsecase *component_category_usecase.ComponentCategoryUsecase
}

func NewDeps(db *gorm.DB) *Deps {
	return &Deps{
		db: db,
	}
}

func (d *Deps) ComponentTypeRepo() *component_type_repo_mysql.ComponentTypeRepo {
	if d.componentTypeRepo == nil {
		d.componentTypeRepo = component_type_repo_mysql.NewComponentTypeRepo(d.db)
	}
	return d.componentTypeRepo
}
func (d *Deps) ComponentTypeUsecase() *component_type_usecase.ComponentTypeUsecase {
	if d.componentTypeUsecase == nil {
		d.componentTypeUsecase = component_type_usecase.NewComponentTypeUsecase(d.ComponentTypeRepo())
	}
	return d.componentTypeUsecase
}

func (d *Deps) ComponentCategoryRepo() *component_category_repo_mysql.ComponentCategoryRepo {
	if d.componentCategoryRepo == nil {
		d.componentCategoryRepo = component_category_repo_mysql.NewComponentCategoryRepo(d.db)
	}
	return d.componentCategoryRepo
}
func (d *Deps) ComponentCategoryUsecase() *component_category_usecase.ComponentCategoryUsecase {
	if d.componentCategoryUsecase == nil {
		d.componentCategoryUsecase = component_category_usecase.NewComponentCategoryUsecase(d.ComponentCategoryRepo())
	}
	return d.componentCategoryUsecase
}
