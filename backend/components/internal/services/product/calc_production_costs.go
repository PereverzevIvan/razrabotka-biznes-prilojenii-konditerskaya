package services_product

import (
	"errors"
	"fmt"

	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models/logic_errors"
)

func (service *productService) CalcProductionPrice(product_id int) (
	total_price float64,
	purchased_components_prices map[int]float64,
	to_purchase_components_prices map[int]float64,
	err error,
) {
	needed_recipe_components_map, err := service.CountNeededRecipeComponents(product_id)
	if err != nil {
		return 0, nil, nil, err
	}

	purchased_components_prices = make(map[int]float64)
	to_purchase_components_prices = make(map[int]float64)

	// Для каждого компонента нужно посчитать стоимость за нужное кол-во
	// Если кол-ва приоретённый компонентов недостаточно, то ищем в поставщиках
	// Если нет поставщика компонета, то возвращаем ошибку
	for component_id, count_needed := range needed_recipe_components_map {

		// Получить цену за компоненты, которые нужны для его сборки
		// 1. Получить количестов имеющихся компонентов
		// 2. Посчитать стоимость
		remaining_count, err := service.purchasedComponentRepo.CountRemainingComponents(component_id)
		if err != nil {
			return 0, nil, nil, err
		}

		remaining_count_to_use := min(count_needed, int(remaining_count))

		remaining_components_price, err := service.purchasedComponentRepo.
			CalcPriceOfRequiredCount(component_id, remaining_count_to_use)
		if err != nil {
			return 0, nil, nil, err
		}

		purchased_components_prices[component_id] = remaining_components_price

		// Уменьшаем кол-во компонентов, которые необходимы для изготовления продута
		count_needed -= remaining_count_to_use
		if count_needed <= 0 {
			continue
		}

		// Если кол-ва приоретённый компонентов недостаточно, то ищем в поставщиках
		// Если нет поставщика компонета, то возвращаем ошибку

		price_per_component, err := service.supplierComponentRepo.FastestDeliveryComponentPrice(component_id)
		if err != nil {
			switch {
			case errors.Is(err, logic_errors.ErrNoSupplierForComponent):
				return 0, nil, nil, fmt.Errorf("%w: %d", err, component_id)
			}
			return 0, nil, nil, err
		}

		to_purchase_components_prices[component_id] = float64(count_needed) * price_per_component
	}

	// Посчитать общую стоимость
	for _, purchased_component_total_price := range purchased_components_prices {
		total_price += purchased_component_total_price
	}

	for _, to_purchase_component_total_price := range to_purchase_components_prices {
		total_price += to_purchase_component_total_price
	}

	return total_price, purchased_components_prices, to_purchase_components_prices, nil
}
