package services_order

import (
	"fmt"

	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/models"
	params_order "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/models/params/order"
)

func (service *orderService) GetAll(params *params_order.GetAllParams) ([]models.Order, error) {

	switch params.UserRole {
	case models.KRoleCustomer.Name():
		return service.orderRepo.GetAll(&params_order.RepoGetAllParams{
			CustomerID: params.UserId,
		})
	case models.KRoleClientManager.Name():
		return service.orderRepo.GetAll(&params_order.RepoGetAllParams{
			StatusIDs:   []int{int(models.KOrderStatusNew)},
			ORManagerID: params.UserId,
		})
	case models.KRolePurchaseManager.Name():
		return service.orderRepo.GetAll(&params_order.RepoGetAllParams{
			StatusIDs: []int{int(models.KOrderStatusPurchase)},
		})
	case models.KRoleMaster.Name():
		return service.orderRepo.GetAll(&params_order.RepoGetAllParams{
			StatusIDs: []int{
				int(models.KOrderStatusProduction),
				int(models.KOrderStatusControl),
			},
		})
	case models.KRoleDirector.Name():
		return service.orderRepo.GetAll(&params_order.RepoGetAllParams{})
	}

	return nil, fmt.Errorf("unkown user role: %s", params.UserRole)
}
