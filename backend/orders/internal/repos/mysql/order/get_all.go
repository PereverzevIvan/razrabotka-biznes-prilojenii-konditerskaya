package repos_mysql_order

import "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/models"

func (repo *orderRepo) GetAll() ([]models.Order, error) {
	var orders []models.Order

	err := repo.conn.Find(&orders).Error
	if err != nil {
		return nil, err
	}

	return orders, nil
}
