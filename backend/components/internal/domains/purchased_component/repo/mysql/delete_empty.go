package purchased_component_repo_mysql

import (
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models/logic_errors"
)

func (r *PurchasedComponentRepo) DeleteEmpty(purchased_component_id int) error {
	result := r.db.
		Where("id = ?", purchased_component_id).
		Where("quantity = 0").
		Delete(&models.PurchasedComponent{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return logic_errors.ErrDeleteNonEmptyQuantity
	}

	return nil
}
