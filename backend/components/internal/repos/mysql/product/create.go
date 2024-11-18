package repos_mysql_product

import "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models"

func (repo *productRepo) Create(product *models.Product) (*models.Product, error) {

	err := repo.conn.Create(product).Error

	if err != nil {
		return nil, err
	}

	return product, nil
}
