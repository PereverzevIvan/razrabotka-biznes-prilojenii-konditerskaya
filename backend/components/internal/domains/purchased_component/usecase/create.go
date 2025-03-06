package purchased_component_usecase

import "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models"

func (u *PurchasedComponentUsecase) Create(purchased_component *models.PurchasedComponent) error {
	purchased_component.ID = 0

	return u.purchasedComponentRepo.Create(purchased_component)
}
