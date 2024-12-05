package repos_mysql_order

import (
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/models"
	params_order "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/models/params/order"
	"gorm.io/gorm"
)

func (repo *orderRepo) GetAll(params *params_order.RepoGetAllParams) ([]models.Order, error) {
	var orders []models.Order

	err := repo.conn.
		Scopes(getAllParams(params)).
		Find(&orders).Error
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func getAllParams(params *params_order.RepoGetAllParams) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if params == nil {
			return db
		}

		if params.CustomerID > 0 {
			db = db.Where("customer_id = ?", params.CustomerID)
		}

		if params.ManagerID > 0 {
			db = db.Where("manager_id = ?", params.ManagerID)
		}

		if len(params.StatusIDs) > 0 {
			db = db.Where("status_id IN ?", params.StatusIDs)
		}

		if params.ORManagerID > 0 {
			db = db.Or("manager_id = ?", params.ORManagerID)
		}

		return db
	}
}
