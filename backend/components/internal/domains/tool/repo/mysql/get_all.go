package tool_repo_mysql

import (
	tool_params "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/domains/tool/params"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models"
	"gorm.io/gorm"
)

func (r *ToolRepo) GetAll(params *tool_params.GetAllParams) ([]models.Tool, error) {

	var tools []models.Tool

	err := r.db.
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

func scopeGetAllParams(params *tool_params.GetAllParams) func(db *gorm.DB) *gorm.DB {
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
