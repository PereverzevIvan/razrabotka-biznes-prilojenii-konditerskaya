package repos_mysql_order

import "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/models"

func (repo *orderRepo) GetByID(id int) (*models.Order, error) {
	var order models.Order

	err := repo.conn.First(&order, id).Error
	if err != nil {
		return nil, err
	}

	return &order, nil
}
