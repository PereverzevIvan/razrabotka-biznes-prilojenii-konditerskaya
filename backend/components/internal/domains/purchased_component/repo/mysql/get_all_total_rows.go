package purchased_component_repo_mysql

import (
	purchased_component_params "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/domains/purchased_component/params"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models"
)

func (r *PurchasedComponentRepo) GetAllTotalRows(
	component_category_id int,
	params *purchased_component_params.GetAllParams,
) (int64, error) {
	var total_rows int64

	err := r.db.
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
