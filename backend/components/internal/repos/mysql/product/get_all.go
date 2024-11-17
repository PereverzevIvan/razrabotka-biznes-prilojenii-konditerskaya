package repos_mysql_product

import "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models"

func (repo *productRepo) GetAll() ([]models.Product, error) {
	var products []models.Product

	err := repo.conn.Find(&products).Error
	if err != nil {
		return nil, err
	}

	return products, nil
}
