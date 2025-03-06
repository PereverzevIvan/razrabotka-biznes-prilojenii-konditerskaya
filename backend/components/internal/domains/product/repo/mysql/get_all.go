package product_repo_mysql

import "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models"

func (r *ProductRepo) GetAll() ([]models.Product, error) {
	var products []models.Product

	err := r.db.Find(&products).Error
	if err != nil {
		return nil, err
	}

	return products, nil
}
