package services_purchased_component

import "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/services"

type purchasedComponentService struct {
	purchasedComponentRepo services.IPurchasedComponentRepo
}

func NewPurchasedComponentService(purchasedComponentRepo services.IPurchasedComponentRepo) *purchasedComponentService {
	return &purchasedComponentService{
		purchasedComponentRepo: purchasedComponentRepo,
	}
}
