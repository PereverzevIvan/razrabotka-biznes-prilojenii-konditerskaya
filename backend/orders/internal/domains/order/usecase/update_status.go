package order_usecase

import (
	"errors"
	"fmt"

	order_params "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/domains/order/params"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/models"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/models/logic_errors"
	services_utils "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/services/utils"
)

func (service *OrderUsecase) UpdateStatus(params *order_params.UpdateStatusParams) (order *models.Order, err error) {
	switch params.NewStatusID {
	// case int(KOrderStatusNew):
	// 	return
	case int(models.KOrderStatusCanceled):
		return service.updateToCanceled(params)
	case int(models.KOrderStatusCompilingRecipe):
		return service.updateToCompilingRecipe(params)
	case int(models.KOrderStatusConfirmation):
		return service.updateToConfirmation(params)
	case int(models.KOrderStatusPurchase):
		return service.updateToPurchase(params)
	case int(models.KOrderStatusProduction):
		return service.updateToProduction(params)
	case int(models.KOrderStatusControl):
		return service.updateToControl(params)
	case int(models.KOrderStatusReady):
		return service.updateToReady(params)
	case int(models.KOrderStatusCompleted):
		return service.updateToCompleted(params)
	default:
		return nil, fmt.Errorf("can't update to unkown status with id: %d", params.NewStatusID)
	}
}

func (service *OrderUsecase) updateToCanceled(params *order_params.UpdateStatusParams) (order *models.Order, err error) {
	var (
		isCustomer      bool = (params.UserRole == models.KRoleCustomer.Name())
		isClientManager      = (params.UserRole == models.KRoleClientManager.Name())
	)
	if !isCustomer && !isClientManager {
		return nil, fmt.Errorf("%w: only customer and client manager can cancel order", logic_errors.ErrInvalidUpdate)
	}

	var data *order_params.UpdateStatusToCanceledData
	// Данные для отмены может указать только менеджер по клиентам

	if isClientManager {
		if params.Data == nil {
			return nil, fmt.Errorf("%w: data is required", logic_errors.ErrInvalidUpdate)
		}

		err = services_utils.JsonUnmarshalInterface(params.Data, &data)
		if err != nil {
			return nil, err
		}

		if validate_errs := data.Validate(); len(validate_errs) != 0 {
			for _, validate_err := range validate_errs {
				err = errors.Join(err, errors.New(validate_err))
			}
			return nil, fmt.Errorf("%w: %w", logic_errors.ErrInvalidUpdate, err)
		}
	}

	order, err = service.GetByID(params.OrderID)
	if err != nil {
		return nil, err
	}

	if order.StatusID == int(models.KOrderStatusCanceled) {
		return nil, fmt.Errorf("%w: order is already canceled", logic_errors.ErrInvalidUpdate)
	}
	if order.StatusID > int(models.KOrderStatusConfirmation) {
		return nil, fmt.Errorf("%w: order which was confirmed cannot be canceled", logic_errors.ErrInvalidUpdate)
	}

	switch {
	case isClientManager:
		if order.ManagerID == nil || *order.ManagerID != params.UserID {
			return nil, fmt.Errorf("%w: only manager of this order can cancel it", logic_errors.ErrInvalidUpdate)
		}
	case isCustomer:
		if order.CustomerID != params.UserID {
			return nil, fmt.Errorf("%w: only customer of this order can cancel it", logic_errors.ErrInvalidUpdate)
		}
	}

	if data != nil {
		order.DeclineReason = &data.DeclineReason
	}

	order.StatusID = int(models.KOrderStatusCanceled)
	return service.orderRepo.Update(order)
}

func (service *OrderUsecase) updateToCompilingRecipe(params *order_params.UpdateStatusParams) (order *models.Order, err error) {
	if params.UserRole != models.KRoleClientManager.Name() {
		return nil, fmt.Errorf("%w: only client manager can compile recipe", logic_errors.ErrInvalidUpdate)
	}

	order, err = service.GetByID(params.OrderID)
	if err != nil {
		return nil, err
	}

	if order.StatusID != int(models.KOrderStatusNew) {
		return nil, fmt.Errorf("%w: order status must be new", logic_errors.ErrInvalidUpdate)
	}

	order.ManagerID = &params.UserID

	order.StatusID = int(models.KOrderStatusCompilingRecipe)
	return service.orderRepo.Update(order)
}

func (service *OrderUsecase) updateToConfirmation(params *order_params.UpdateStatusParams) (order *models.Order, err error) {
	if params.UserRole != models.KRoleClientManager.Name() {
		return nil, fmt.Errorf("%w: only client manager can update status to confirmation", logic_errors.ErrInvalidUpdate)
	}

	if params.Data == nil {
		return nil, fmt.Errorf("%w: data is required", logic_errors.ErrInvalidUpdate)
	}

	var data order_params.UpdateStatusToConfirmationData
	err = services_utils.JsonUnmarshalInterface(params.Data, &data)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", logic_errors.ErrInvalidUpdate, err)
	}

	if validate_errs := data.Validate(); len(validate_errs) != 0 {
		for _, validate_err := range validate_errs {
			err = errors.Join(err, errors.New(validate_err))
		}
		return nil, fmt.Errorf("%w: %w", logic_errors.ErrInvalidUpdate, err)
	}

	order, err = service.GetByID(params.OrderID)
	if err != nil {
		return nil, err
	}

	if order.StatusID != int(models.KOrderStatusCompilingRecipe) {
		return nil, fmt.Errorf("%w: order status must be compiling recipe", logic_errors.ErrInvalidUpdate)
	}

	if order.ManagerID == nil || *order.ManagerID != params.UserID {
		return nil, fmt.Errorf("%w: only manager of this order can update status to confirmation", logic_errors.ErrInvalidUpdate)
	}

	order.PlannedCompletionAt = data.PlannedCompletionAt
	order.Cost = data.Cost

	order.StatusID = int(models.KOrderStatusConfirmation)
	return service.orderRepo.Update(order)
}

func (service *OrderUsecase) updateToPurchase(params *order_params.UpdateStatusParams) (order *models.Order, err error) {
	if params.UserRole != models.KRoleCustomer.Name() {
		return nil, fmt.Errorf("%w: only customer can change status from confirmation to purchase", logic_errors.ErrInvalidUpdate)
	}

	order, err = service.GetByID(params.OrderID)
	if err != nil {
		return nil, err
	}

	if order.StatusID != int(models.KOrderStatusConfirmation) {
		return nil, fmt.Errorf("%w: order status must be confirmation", logic_errors.ErrInvalidUpdate)
	}

	if order.CustomerID != params.UserID {
		return nil, fmt.Errorf("%w: only customer of this order can update status from confirmation to purchase", logic_errors.ErrInvalidUpdate)
	}

	order.StatusID = int(models.KOrderStatusPurchase)
	return service.orderRepo.Update(order)
}

func (service *OrderUsecase) updateToProduction(params *order_params.UpdateStatusParams) (order *models.Order, err error) {
	if params.UserRole != models.KRolePurchaseManager.Name() {
		return nil, fmt.Errorf("%w: only purchase manager can update status to production", logic_errors.ErrInvalidUpdate)
	}

	order, err = service.GetByID(params.OrderID)
	if err != nil {
		return nil, err
	}

	if order.StatusID != int(models.KOrderStatusPurchase) {
		return nil, fmt.Errorf("%w: order status must be purchase", logic_errors.ErrInvalidUpdate)
	}

	err = service.componentRepo.ConsumeProductComponents(order.ProductID)
	if err != nil {
		return nil, err
	}

	order.StatusID = int(models.KOrderStatusProduction)
	return service.orderRepo.Update(order)
}

func (service *OrderUsecase) updateToControl(params *order_params.UpdateStatusParams) (order *models.Order, err error) {
	if params.UserRole != models.KRoleMaster.Name() {
		return nil, fmt.Errorf("%w: only master can update status to control", logic_errors.ErrInvalidUpdate)
	}

	order, err = service.GetByID(params.OrderID)
	if err != nil {
		return nil, err
	}

	if order.StatusID != int(models.KOrderStatusProduction) {
		return nil, fmt.Errorf("%w: order status must be production", logic_errors.ErrInvalidUpdate)
	}

	order.StatusID = int(models.KOrderStatusControl)
	return service.orderRepo.Update(order)
}

func (service *OrderUsecase) updateToReady(params *order_params.UpdateStatusParams) (order *models.Order, err error) {
	if params.UserRole != models.KRoleMaster.Name() {
		return nil, fmt.Errorf("%w: only master can update status to ready", logic_errors.ErrInvalidUpdate)
	}

	is_order_valid, err := service.qualityControlRepo.IsOrderValid(params.OrderID)
	if err != nil {
		return nil, err
	}

	if !is_order_valid {
		return nil, fmt.Errorf("%w: not all quality control checks are passed", logic_errors.ErrInvalidUpdate)
	}

	order, err = service.GetByID(params.OrderID)
	if err != nil {
		return nil, err
	}

	if order.StatusID != int(models.KOrderStatusControl) {
		return nil, fmt.Errorf("%w: order status must be control", logic_errors.ErrInvalidUpdate)
	}

	order.StatusID = int(models.KOrderStatusReady)
	return service.orderRepo.Update(order)
}

func (service *OrderUsecase) updateToCompleted(params *order_params.UpdateStatusParams) (order *models.Order, err error) {
	if params.UserRole != models.KRoleClientManager.Name() {
		return nil, fmt.Errorf("%w: only client manager can update status to completed", logic_errors.ErrInvalidUpdate)
	}

	order, err = service.GetByID(params.OrderID)
	if err != nil {
		return nil, err
	}

	if order.StatusID != int(models.KOrderStatusReady) {
		return nil, fmt.Errorf("%w: order status must be ready", logic_errors.ErrInvalidUpdate)
	}

	if order.ManagerID == nil || *order.ManagerID != params.UserID {
		return nil, fmt.Errorf("%w: only manager of this order can update status to completed", logic_errors.ErrInvalidUpdate)
	}

	order.StatusID = int(models.KOrderStatusCompleted)
	return service.orderRepo.Update(order)
}
