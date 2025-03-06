package order_repo_mysql

import "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/models"

func (repo *OrderRepo) Create(order *models.Order) (*models.Order, error) {
	err := repo.db.Create(order).Error
	if err != nil {
		return nil, err
	}

	return order, nil
}
