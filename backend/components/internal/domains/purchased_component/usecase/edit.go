package purchased_component_usecase

import "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models"

func (u *PurchasedComponentUsecase) Edit(purchased_component *models.PurchasedComponent) error {
	return u.purchasedComponentRepo.Edit(purchased_component)
}
