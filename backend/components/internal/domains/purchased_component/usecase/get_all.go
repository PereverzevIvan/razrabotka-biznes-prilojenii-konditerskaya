package purchased_component_usecase

import (
	purchased_component_params "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/domains/purchased_component/params"
	results_purchased_component "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models/results/purchased_component"
)

func (u *PurchasedComponentUsecase) GetAll(
	component_category_id int,
	params *purchased_component_params.GetAllParams,
) (*results_purchased_component.GetAllResults, error) {
	result := &results_purchased_component.GetAllResults{}
	var err error

	result.Data, err = u.purchasedComponentRepo.GetAll(component_category_id, params)
	if err != nil {
		return nil, err
	}

	result.TotalRows, err = u.purchasedComponentRepo.GetAllTotalRows(component_category_id, params)
	if err != nil {
		return nil, err
	}

	result.TotalCount, err = u.purchasedComponentRepo.GetAllTotalCount(component_category_id, params)
	if err != nil {
		return nil, err
	}

	result.TotalPrice, err = u.purchasedComponentRepo.GetAllTotalPrice(component_category_id, params)
	if err != nil {
		return nil, err
	}

	return result, nil
}
