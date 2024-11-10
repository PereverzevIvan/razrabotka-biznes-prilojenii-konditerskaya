package repos_mysql_purchased_component

import (
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models/logic_errors"
)

func (repo *purchasedComponentRepo) DeleteEmpty(purchased_component_id int) error {
	result := repo.conn.
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
