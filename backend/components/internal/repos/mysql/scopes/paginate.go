package repos_mysql_scopes

import (
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models/params"
	"gorm.io/gorm"
)

const DEAFAULT_LIMIT = 10

func ScopePaginateParams(params *params.PaginateParams) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if params == nil {
			return db.Offset(0).Limit(DEAFAULT_LIMIT)
		}

		limit := params.Limit
		if limit <= 0 {
			limit = DEAFAULT_LIMIT
		}

		offset := 0
		if params.Page > 1 {
			offset = (params.Page - 1) * params.Limit
		}

		return db.Offset(offset).Limit(limit)
	}
}
