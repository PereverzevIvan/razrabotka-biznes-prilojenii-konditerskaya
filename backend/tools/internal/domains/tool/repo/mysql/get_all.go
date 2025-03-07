package tool_repo_mysql

import (
	tool_params "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/domains/tool/params"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/models"
	"gorm.io/gorm"
)

func (r *ToolRepo) GetAll(params *tool_params.GetAllParams) ([]models.Tool, error) {

	var tools []models.Tool

	err := r.db.
		Scopes(
			r.scopeGetAllParams(params),
		).
		Joins("ToolType").
		Joins("Supplier").
		Joins("DegreeOfWear").
		Find(&tools).
		Error
	if err != nil {
		return nil, err
	}

	return tools, nil
}

func scopeGetAllParams(params *tool_params.GetAllParams) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if params == nil {
			return db
		}

		return db
	}
}

func (repo *ToolRepo) scopeGetAllParams(params *tool_params.GetAllParams) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {

		if params.Name != nil {
			db = db.Where("name LIKE ?", "%"+*params.Name+"%")
		}

		if params.ToolType != nil {
			db = db.Where("tool_type_id = ?", *params.ToolType)
		}

		if params.DegreeOfWear != nil {
			db = db.Where("degree_of_wear_id = ?", *params.DegreeOfWear)
		}

		if params.SupplierID != nil {
			db = db.Where("supplier_id = ?", *params.SupplierID)
		}

		if params.OnlyServiceable {
			db = db.Where(
				"degree_of_wear_id = ? OR degree_of_wear_id = ?",
				models.KDegreeOfWearNew,
				models.KDegreeOfWearWornOut,
			)
		}

		if params.Sort != nil {
			switch *params.Sort {
			case tool_params.KSortByToolType:
				db = db.Order(tool_params.KSortByToolType)
			case tool_params.KSortByDegreeOfWear:
				db = db.Order(tool_params.KSortByDegreeOfWear)
			case tool_params.KSortByName:
				db = db.Order(tool_params.KSortByName)
			case tool_params.KSortByPurchaseDate:
				db = db.Order(tool_params.KSortByPurchaseDate)
			default:
				db = db.Order(tool_params.KDefaultSort)
			}
		}
		return db
	}
}
