package purchased_component_repo_mysql

import (
	purchased_component_params "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/domains/purchased_component/params"
)

func (r *PurchasedComponentRepo) GetAllTotalCount(
	component_category_id int,
	params *purchased_component_params.GetAllParams,
) (int64, error) {
	var total_count *int64

	err := r.db.
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
