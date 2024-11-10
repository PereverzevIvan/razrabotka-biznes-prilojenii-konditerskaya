package repos_mysql_tool

import (
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/models"
	params_tool "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/models/params/tool"
	"gorm.io/gorm"
)

func (repo *toolRepo) GetAll(params *params_tool.GetAllParams) ([]models.Tool, error) {

	var tools []models.Tool

	err := repo.conn.
		Scopes(
			repo.scopeGetAllParams(params),
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

func (repo *toolRepo) scopeGetAllParams(params *params_tool.GetAllParams) func(db *gorm.DB) *gorm.DB {
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

		if params.Sort != nil {
			switch *params.Sort {
			case params_tool.KSortByToolType:
				db = db.Order(params_tool.KSortByToolType)
			case params_tool.KSortByDegreeOfWear:
				db = db.Order(params_tool.KSortByDegreeOfWear)
			case params_tool.KSortByName:
				db = db.Order(params_tool.KSortByName)
			case params_tool.KSortByPurchaseDate:
				db = db.Order(params_tool.KSortByPurchaseDate)
			default:
				db = db.Order(params_tool.KDefaultSort)
			}
		}
		return db
	}
}
