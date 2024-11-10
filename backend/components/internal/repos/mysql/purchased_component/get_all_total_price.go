package repos_mysql_purchased_component

import (
	params_purchased_component "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models/params/purchased_component"
)

func (repo *purchasedComponentRepo) GetAllTotalPrice(
	component_category_id int,
	params *params_purchased_component.GetAllParams,
) (float64, error) {
	var total_price *float64

	err := repo.conn.
		Scopes(
			scopeGetAllParamsAggregate(component_category_id, params),
		).
		Select("SUM(quantity * purchase_price) as total_price").
		Row().Scan(&total_price)

	if err != nil {
		return 0, err
	}
	if total_price == nil {
		return 0, nil
	}

	return *total_price, nil
}
