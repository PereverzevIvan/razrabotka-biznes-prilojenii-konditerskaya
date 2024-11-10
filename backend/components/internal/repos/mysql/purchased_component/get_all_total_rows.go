package repos_mysql_purchased_component

import (
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models"
	params_purchased_component "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models/params/purchased_component"
)

func (repo *purchasedComponentRepo) GetAllTotalRows(
	component_category_id int,
	params *params_purchased_component.GetAllParams,
) (int64, error) {
	var total_rows int64

	err := repo.conn.
		Model(&models.PurchasedComponent{}).
		Scopes(
			scopeGetAllParams(component_category_id, params),
		).
		Count(&total_rows).
		Error

	if err != nil {
		return 0, err
	}

	return total_rows, nil
}
