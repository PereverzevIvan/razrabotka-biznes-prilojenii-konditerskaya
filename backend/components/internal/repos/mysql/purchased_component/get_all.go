package repos_mysql_purchased_component

import (
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models"
	params_purchased_component "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models/params/purchased_component"
	repos_mysql_scopes "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/repos/mysql/scopes"
	"gorm.io/gorm"
)

func (repo *purchasedComponentRepo) GetAll(
	component_category_id int,
	params *params_purchased_component.GetAllParams,
) ([]models.PurchasedComponent, error) {
	var purchased_components []models.PurchasedComponent

	err := repo.conn.
		Scopes(
			scopeGetAllParams(params),
			repos_mysql_scopes.ScopePaginateParams(params.PaginateParams),
		).
		InnerJoins("Component").
		InnerJoins("Component.ComponentType").
		InnerJoins("Component.ComponentType.ComponentCategory").
		Find(&purchased_components).Error

	if err != nil {
		return nil, err
	}

	return purchased_components, nil
}

func scopeGetAllParams(params *params_purchased_component.GetAllParams) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if params == nil {
			return db
		}

		if params.ComponentName != nil {
			db = db.Where("Component.Name LIKE ?", "%"+*params.ComponentName+"%")
		}

		if params.ComponentArticle != nil {
			db = db.Where("Component.Article LIKE ?", "%"+*params.ComponentArticle+"%")
		}

		if params.ShelfLifeFrom != nil {
			db = db.Where("Component.ShelfLife >= ?", params.ShelfLifeFrom)
		}

		if params.ShelfLifeTo != nil {
			db = db.Where("Component.ShelfLife <= ?", params.ShelfLifeTo)
		}

		switch params.GetSort() {
		// case params_purchased_component.KSortByComponentName:
		// 	db = db.Order("params_purchased_component.KSortByComponentName")
		// case params_purchased_component.KSortByComponentArticle:
		// 	db = db.Order("params_purchased_component.KSortByComponentArticle")
		case params_purchased_component.KGetAllSortByQuantity:
			db = db.Order("Quantity")
		case params_purchased_component.KGetAllSortByShelfLife:
			db = db.Order("ShelfLife")
		}

		return db
	}
}
