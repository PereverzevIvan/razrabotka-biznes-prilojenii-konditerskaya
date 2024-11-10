package services_purchased_component

import "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models"

func (service *purchasedComponentService) Edit(purchased_component *models.PurchasedComponent) error {
	return service.purchasedComponentRepo.Edit(purchased_component)
}
