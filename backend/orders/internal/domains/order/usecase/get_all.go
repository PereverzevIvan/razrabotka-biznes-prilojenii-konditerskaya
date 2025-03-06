package order_usecase

import (
	"fmt"

	order_params "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/domains/order/params"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/models"
)

func (service *OrderUsecase) GetAll(params *order_params.GetAllParams) ([]models.Order, error) {

	switch params.UserRole {
	case models.KRoleCustomer.Name():
		return service.orderRepo.GetAll(&order_params.RepoGetAllParams{
			CustomerID: params.UserId,
		})
	case models.KRoleClientManager.Name():
		return service.orderRepo.GetAll(&order_params.RepoGetAllParams{
			StatusIDs:   []int{int(models.KOrderStatusNew)},
			ORManagerID: params.UserId,
		})
	case models.KRolePurchaseManager.Name():
		return service.orderRepo.GetAll(&order_params.RepoGetAllParams{
			StatusIDs: []int{int(models.KOrderStatusPurchase)},
		})
	case models.KRoleMaster.Name():
		return service.orderRepo.GetAll(&order_params.RepoGetAllParams{
			StatusIDs: []int{
				int(models.KOrderStatusProduction),
				int(models.KOrderStatusControl),
			},
		})
	case models.KRoleDirector.Name():
		return service.orderRepo.GetAll(&order_params.RepoGetAllParams{})
	}

	return nil, fmt.Errorf("unkown user role: %s", params.UserRole)
}
