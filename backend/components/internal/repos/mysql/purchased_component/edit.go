package repos_mysql_purchased_component

import "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models"

func (repo *purchasedComponentRepo) Edit(purchased_component *models.PurchasedComponent) error {
	result := repo.conn.
		Model(&purchased_component).
		Updates(&purchased_component)

	if result.Error != nil {
		return result.Error
	}

	// if result.RowsAffected == 0 {
	// 	return models.ErrFK
	// }

	return nil
}
