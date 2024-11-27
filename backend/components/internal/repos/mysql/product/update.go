package repos_mysql_product

import "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models"

func (repo *productRepo) Update(product *models.Product) error {
	return repo.conn.Save(product).Error
}
