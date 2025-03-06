package product_repo_mysql

import "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models"

func (r *ProductRepo) Update(product *models.Product) error {
	return r.db.Save(product).Error
}
