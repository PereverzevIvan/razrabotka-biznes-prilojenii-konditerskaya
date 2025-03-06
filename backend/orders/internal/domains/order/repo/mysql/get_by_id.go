package order_repo_mysql

import "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/models"

func (repo *OrderRepo) GetByID(id int) (*models.Order, error) {
	var order models.Order

	err := repo.db.First(&order, id).Error
	if err != nil {
		return nil, err
	}

	return &order, nil
}
