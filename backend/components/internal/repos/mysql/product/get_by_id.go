package repos_mysql_product

import "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models"

func (repo *productRepo) GetByID(id int) (*models.Product, error) {
	var product models.Product

	err := repo.conn.
		Debug().
		Model(&models.Product{}).
		Where("id = ?", id).
		First(&product).
		Error
	if err != nil {
		return nil, err
	}

	err = repo.conn.
		Debug().
		Model(&models.RecipeComponents{}).
		Joins("Component").
		Where("product_id = ?", id).
		Find(&product.RecipeComponents).
		Error
	if err != nil {
		return nil, err
	}

	return &product, nil
}
