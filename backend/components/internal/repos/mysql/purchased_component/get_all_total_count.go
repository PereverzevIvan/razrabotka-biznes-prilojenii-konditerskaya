package repos_mysql_purchased_component

import (
	params_purchased_component "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models/params/purchased_component"
)

func (repo *purchasedComponentRepo) GetAllTotalCount(
	component_category_id int,
	params *params_purchased_component.GetAllParams,
) (int64, error) {
	var total_count *int64

	err := repo.conn.
		Scopes(
			scopeGetAllParamsAggregate(component_category_id, params),
		).
		Select("SUM(quantity) as total_count").
		Row().Scan(&total_count)

	if err != nil {
		return 0, err
	}
	if total_count == nil {
		return 0, nil
	}

	return *total_count, nil
}
