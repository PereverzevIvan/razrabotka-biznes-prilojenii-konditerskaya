package product_usecase

import (
	"fmt"

	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models/logic_errors"
)

func (u *ProductUsecase) MakeProduct(id int) error {
	needed_recipe_components_map, err := u.CountNeededRecipeComponents(id)
	if err != nil {
		return err
	}

	for component_id, count_needed := range needed_recipe_components_map {
		remaining_count, err := u.purchasedComponentRepo.CountRemainingComponents(component_id)
		if err != nil {
			return err
		}
		if int64(count_needed) > remaining_count {
			return fmt.Errorf(
				"%w: component_id: %d, needed: %d, remaining: %d",
				logic_errors.ErrNotEnoughComponentsToMakeProduct, component_id, count_needed, remaining_count)
		}
	}

	// Списываем компоненты
	for component_id, count_needed := range needed_recipe_components_map {
		err := u.purchasedComponentRepo.DeductComponents(component_id, count_needed)
		if err != nil {
			return err
		}
	}

	return nil
}
