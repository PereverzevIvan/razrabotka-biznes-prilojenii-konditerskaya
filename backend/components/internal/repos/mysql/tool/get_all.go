package repos_mysql_tool

import (
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models"
	params_tool "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models/params/tool"
	"gorm.io/gorm"
)

func (repo *repoTools) GetAll(params *params_tool.GetAllParams) ([]models.Tool, error) {

	var tools []models.Tool

	err := repo.conn.
		Scopes(
			scopeGetAllParams(params),
		).
		Find(&tools).
		Error
	if err != nil {
		return nil, err
	}

	return tools, nil
}

func scopeGetAllParams(params *params_tool.GetAllParams) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if params == nil {
			return db
		}

		if params.ToolTypeID > 0 {
			db = db.Where("type_id = ?", params.ToolTypeID)
		}

		if params.OnlyServiceable {
			db = db.Where(
				"degree_of_wear_id = ? OR degree_of_wear_id = ?",
				models.KDegreeOfWearNew,
				models.KDegreeOfWearWornOut,
			)
		}

		return db
	}
}
