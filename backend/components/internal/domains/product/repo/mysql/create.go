package product_repo_mysql

import "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models"

func (r *ProductRepo) Create(product *models.Product) (*models.Product, error) {

	err := r.db.Create(product).Error

	if err != nil {
		return nil, err
	}

	return product, nil
}
