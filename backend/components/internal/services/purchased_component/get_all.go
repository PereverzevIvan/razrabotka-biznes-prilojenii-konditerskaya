package services_purchased_component

import results_purchased_component "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models/results/purchased_component"

func (service *purchasedComponentService) GetAll() (*results_purchased_component.GetAllResults, error) {
	return service.purchasedComponentRepo.GetAll()
}
