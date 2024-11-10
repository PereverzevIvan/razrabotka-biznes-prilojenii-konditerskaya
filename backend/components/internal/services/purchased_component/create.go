package services_purchased_component

import "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models"

func (service *purchasedComponentService) Create(purchased_component *models.PurchasedComponent) error {
	purchased_component.ID = 0

	return service.purchasedComponentRepo.Create(purchased_component)
}
