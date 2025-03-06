package product_usecase

import (
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models/logic_errors"
)

func (u *ProductUsecase) CountNeededRecipeComponents(
	id int,
) (map[int]int, error) {
	main_product, err := u.productRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	// кэш для полученных ранее продуктов и компонентов
	obtained_products_map := map[int]*models.Product{
		main_product.ID: main_product,
	}
	// obtained_components_map := map[int]*models.Component{}

	// key - id продукта, value - количество необходимых компонентов
	needed_recipe_components_map := map[int]int{}

	// Проверка на цикл в графе рецепте
	// Если в графе рецепта есть цикл, то возвращаем ошибку
	// Привмер:
	// Для продукта A в рецепте нужен продукт B,
	// который в свою очередь содержит в рецепте продукт A.
	// Тогда, чтобы сделать A нужен B, но для B нужен A следовательно есть цикл
	path_product_set := map[int]bool{
		main_product.ID: true,
	}

	var dfsProductRecipe func(cur_product *models.Product, cur_count int) error
	dfsProductRecipe = func(cur_product *models.Product, cur_count int) error {

		// Получаем компоненты этого полуфабриката,
		for _, recipe_component := range cur_product.RecipeComponents {

			component_id := recipe_component.ComponentID

			needed_recipe_components_map[component_id] += cur_count * recipe_component.Count
		}

		// Получаем полуфабрикаты текущего продукта,
		// если полуфабрикат ранее был получен, и не образует цикл, то подставляем его
		for _, recipe_semiproduct := range cur_product.RecipeSemiProducts {

			semiproduct_id := recipe_semiproduct.SemiproductID
			// Проверка на цикл
			if _, ok := path_product_set[semiproduct_id]; ok {
				return logic_errors.ErrCycleDetectedInProductRecipe
			}

			// Если ранее получали этот полуфабрикат, то используем его
			if semiproduct, ok := obtained_products_map[semiproduct_id]; ok {
				recipe_semiproduct.Semiproduct = semiproduct
			} else {
				// Иначе получаем из БД
				semiproduct, err = u.productRepo.GetByID(semiproduct_id)
				if err != nil {
					return err
				}

				obtained_products_map[semiproduct_id] = semiproduct
				recipe_semiproduct.Semiproduct = semiproduct
			}

			path_product_set[semiproduct_id] = true

			err = dfsProductRecipe(
				recipe_semiproduct.Semiproduct,
				cur_count*recipe_semiproduct.Count,
			)
			if err != nil {
				return err
			}
		}

		return nil
	}

	err = dfsProductRecipe(main_product, 1)
	if err != nil {
		switch err {
		case logic_errors.ErrCycleDetectedInProductRecipe:
			return needed_recipe_components_map, logic_errors.ErrCycleDetectedInProductRecipe
		}

		return nil, err
	}

	return needed_recipe_components_map, nil
}
