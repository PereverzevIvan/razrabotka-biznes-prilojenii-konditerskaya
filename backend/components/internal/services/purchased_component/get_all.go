package services_purchased_component

import (
	params_purchased_component "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models/params/purchased_component"
	results_purchased_component "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models/results/purchased_component"
)

func (service *purchasedComponentService) GetAll(
	component_category_id int,
	params *params_purchased_component.GetAllParams,
) (*results_purchased_component.GetAllResults, error) {
	result := &results_purchased_component.GetAllResults{}
	var err error

	result.Data, err = service.purchasedComponentRepo.GetAll(component_category_id, params)
	if err != nil {
		return nil, err
	}

	result.TotalRows, err = service.purchasedComponentRepo.GetAllTotalRows(component_category_id, params)
	if err != nil {
		return nil, err
	}

	result.TotalCount, err = service.purchasedComponentRepo.GetAllTotalCount(component_category_id, params)
	if err != nil {
		return nil, err
	}

	result.TotalPrice, err = service.purchasedComponentRepo.GetAllTotalPrice(component_category_id, params)
	if err != nil {
		return nil, err
	}

	return result, nil
}
