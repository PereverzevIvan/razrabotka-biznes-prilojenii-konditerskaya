package repos_mysql_order

import "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/models"

func (repo *orderRepo) Create(order *models.Order) (*models.Order, error) {
	err := repo.conn.Create(order).Error
	if err != nil {
		return nil, err
	}

	return order, nil
}
