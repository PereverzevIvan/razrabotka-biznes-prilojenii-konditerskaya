package purchased_component_repo_mysql

import (
	purchased_component_params "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/domains/purchased_component/params"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models"
	repos_mysql_scopes "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/repos/mysql/scopes"
	"gorm.io/gorm"
)

func (r *PurchasedComponentRepo) GetAll(
	component_category_id int,
	params *purchased_component_params.GetAllParams,
) ([]models.PurchasedComponent, error) {
	var purchased_components []models.PurchasedComponent

	err := r.db.
		Scopes(
			scopeGetAllParams(component_category_id, params),
			repos_mysql_scopes.ScopePaginateParams(params.PaginateParams),
		).
		Find(&purchased_components).
		Error

	if err != nil {
		return nil, err
	}

	return purchased_components, nil
}

func scopeGetAllParams(
	component_category_id int,
	params *purchased_component_params.GetAllParams,
) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		db = db.
			InnerJoins("Component").
			InnerJoins("Component.ComponentType").
			InnerJoins("Component.ComponentType.ComponentCategory").
			Where("Component__ComponentType__ComponentCategory.id = ?", component_category_id)

		if params == nil {
			return db
		}

		if params.ComponentName != nil {
			db = db.Where("Component.name LIKE ?", "%"+*params.ComponentName+"%")
		}

		if params.ComponentArticle != nil {
			db = db.Where("Component.article LIKE ?", "%"+*params.ComponentArticle+"%")
		}

		if params.ShelfLifeFrom != nil {
			db = db.Where("shelf_life >= ?", params.ShelfLifeFrom)
		}

		if params.ShelfLifeTo != nil {
			db = db.Where("shelf_life <= ?", params.ShelfLifeTo)
		}

		switch params.GetSort() {
		case purchased_component_params.KGetAllSortByQuantity:
			db = db.Order("quantity")
		case purchased_component_params.KGetAllSortByShelfLife:
			db = db.Order("shelf_life")
		}

		return db
	}
}

func scopeGetAllParamsAggregate(
	component_category_id int,
	params *purchased_component_params.GetAllParams,
) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		db = db.
			Table(models.PurchasedComponent{}.TableName()).
			Joins("INNER JOIN components ON components.id = purchased_components.component_id").
			Joins("INNER JOIN component_types ON component_types.id = components.component_type_id").
			Joins("INNER JOIN component_categories ON component_categories.id = component_types.component_category_id").
			Where("component_categories.id = ?", component_category_id)

		if params == nil {
			return db
		}

		if params.ComponentName != nil {
			db = db.Where("components.name LIKE ?", "%"+*params.ComponentName+"%")
		}

		if params.ComponentArticle != nil {
			db = db.Where("components.article LIKE ?", "%"+*params.ComponentArticle+"%")
		}

		if params.ShelfLifeFrom != nil {
			db = db.Where("purchased_components.shelf_life >= ?", params.ShelfLifeFrom)
		}

		if params.ShelfLifeTo != nil {
			db = db.Where("purchased_components.shelf_life <= ?", params.ShelfLifeTo)
		}

		return db
	}
}
