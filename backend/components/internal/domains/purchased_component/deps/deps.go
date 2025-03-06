package purchased_component_deps

import (
	component_category_repo_mysql "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/domains/component_category/repo/mysql"
	component_category_usecase "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/domains/component_category/usecase"
	purchased_component_repo_mysql "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/domains/purchased_component/repo/mysql"
	purchased_component_usecase "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/domains/purchased_component/usecase"
	"gorm.io/gorm"
)

type Deps struct {
	db *gorm.DB

	purchasedComponentRepo    *purchased_component_repo_mysql.PurchasedComponentRepo
	purchasedComponentUsecase *purchased_component_usecase.PurchasedComponentUsecase

	componentCategoryRepo    *component_category_repo_mysql.ComponentCategoryRepo
	componentCategoryUsecase *component_category_usecase.ComponentCategoryUsecase
}

func NewDeps(db *gorm.DB) *Deps {
	return &Deps{
		db: db,
	}
}

func (d *Deps) PurchasedComponentRepo() *purchased_component_repo_mysql.PurchasedComponentRepo {
	if d.purchasedComponentRepo == nil {
		d.purchasedComponentRepo = purchased_component_repo_mysql.NewPurchasedComponentRepo(d.db)
	}
	return d.purchasedComponentRepo
}
func (d *Deps) PurchasedComponentUsecase() *purchased_component_usecase.PurchasedComponentUsecase {
	if d.purchasedComponentUsecase == nil {
		d.purchasedComponentUsecase = purchased_component_usecase.NewPurchasedComponentUsecase(d.PurchasedComponentRepo())
	}
	return d.purchasedComponentUsecase
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
