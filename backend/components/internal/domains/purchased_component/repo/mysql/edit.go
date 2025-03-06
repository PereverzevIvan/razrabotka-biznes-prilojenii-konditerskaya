package purchased_component_repo_mysql

import "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models"

func (r *PurchasedComponentRepo) Edit(purchased_component *models.PurchasedComponent) error {
	result := r.db.
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
