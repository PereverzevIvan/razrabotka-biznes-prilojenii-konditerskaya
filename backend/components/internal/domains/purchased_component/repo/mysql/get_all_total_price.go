package purchased_component_repo_mysql

import (
	purchased_component_params "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/domains/purchased_component/params"
)

func (r *PurchasedComponentRepo) GetAllTotalPrice(
	component_category_id int,
	params *purchased_component_params.GetAllParams,
) (float64, error) {
	var total_price *float64

	err := r.db.
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
