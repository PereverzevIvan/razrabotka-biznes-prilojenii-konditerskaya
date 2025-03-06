package order_repo_mysql

import (
	order_params "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/domains/order/params"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/models"
	"gorm.io/gorm"
)

func (repo *OrderRepo) GetAll(params *order_params.RepoGetAllParams) ([]models.Order, error) {
	var orders []models.Order

	err := repo.db.
		Scopes(getAllParams(params)).
		Find(&orders).Error
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func getAllParams(params *order_params.RepoGetAllParams) func(*gorm.DB) *gorm.DB {
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
