package order_usecase

import (
	"fmt"
	"time"

	order_params "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/domains/order/params"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/models"
)

func (service *OrderUsecase) Create(params *order_params.CreateParams) (order *models.Order, err error) {

	order = &models.Order{
		ProductID: params.ProductID,

		CustomerID: params.CustomerID,
		ManagerID:  params.ManagerID,

		Name:        params.Name,
		Description: params.Description,
		Cost:        params.Cost,

		CreatedAt: time.Now(),

		PlannedCompletionAt: params.PlannedCompletionAt,
	}

	switch params.AuthorRole {
	case models.KRoleClientManager.Name():
		order.ManagerID = &params.AuthorId
		order.StatusID = int(models.KOrderStatusCompilingRecipe)
	default:
		order.StatusID = int(models.KOrderStatusNew)
	}

	order.Number, err = service.generateOrderNumber(order.CustomerID, order.CreatedAt)
	if err != nil {
		return nil, err
	}

	return service.orderRepo.Create(order)
}

func (service *OrderUsecase) generateOrderNumber(customer_id int, time_now time.Time) (string, error) {
	var create_date int = int(time_now.Day())
	var create_month int = int(time_now.Month())
	var create_year int = time_now.Year()

	customer, err := service.userRepo.GetByID(customer_id)
	if err != nil {
		return "", err
	}
	if customer == nil {
		return "", fmt.Errorf("customer not found")
	}

	var customer_first_name_first_letter rune = '_'
	if len(customer.FirstName) > 0 {
		customer_first_name_first_letter = []rune(customer.FirstName)[0]
	}

	var customer_last_name_first_letter rune = '_'
	if len(customer.LastName) > 0 {
		customer_last_name_first_letter = []rune(customer.LastName)[0]
	}

	count_today_orders, err := service.orderRepo.CountForDate(time_now)
	if err != nil {
		return "", err
	}

	order_idx := count_today_orders%99 + 1

	return fmt.Sprintf("%02d%02d%d%c%c%02d",
		create_date,
		create_month,
		create_year,
		customer_last_name_first_letter,
		customer_first_name_first_letter,
		order_idx,
	), nil
}
