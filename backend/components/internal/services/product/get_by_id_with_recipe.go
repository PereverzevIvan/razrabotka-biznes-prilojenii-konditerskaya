package services_product

import (
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models/logic_errors"
)

func (service *productService) GetByIDWithRecipe(id int) (*models.Product, error) {
	main_product, err := service.productRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	// кэш для полученных ранее продуктов и компонентов
	obtained_products_map := map[int]*models.Product{
		main_product.ID: main_product,
	}
	obtained_components_map := map[int]*models.Component{}

	obtained_tool_types_map := map[int]*models.ToolType{}

	// Проверка на цикл в графе рецепте
	// Если в графе рецепта есть цикл, то возвращаем ошибку
	// Привмер:
	// Для продукта A в рецепте нужен продукт B,
	// который в свою очередь содержит в рецепте продукт A.
	// Тогда, чтобы сделать A нужен B, но для B нужен A следовательно есть цикл
	path_product_set := map[int]bool{
		main_product.ID: true,
	}

	var dfsProductRecipe func(cur_product *models.Product) error
	dfsProductRecipe = func(cur_product *models.Product) error {

		// Получаем компоненты этого полуфабриката,
		// если компонент ранее был получен, то подставляем его
		for _, recipe_component := range cur_product.RecipeComponents {

			component_id := recipe_component.ComponentID

			// Если ранее получали этот компонент, то используем его
			if component, ok := obtained_components_map[component_id]; ok {
				recipe_component.Component = component
			} else {
				// Иначе получаем компонент из БД
				component, err = service.componentRepo.GetByID(component_id)
				if err != nil {
					return err
				}

				obtained_components_map[component_id] = component
				recipe_component.Component = component
			}
		}

		// Получаем типы приборов для производства этого полуфабриката,
		// если прибор ранее был получен, то подставляем его
		for _, recipe_operation := range cur_product.RecipeOperations {

			if recipe_operation.ToolType == nil {
				continue
			}
			tool_type_id := *recipe_operation.ToolTypeID

			// Если ранее получали этот инструмент, то используем его
			if tool_type, ok := obtained_tool_types_map[tool_type_id]; ok {
				recipe_operation.ToolType = tool_type
			} else {
				// Иначе получаем из БД
				tool_type, err = service.toolTypeRepo.GetByID(tool_type_id)
				if err != nil {
					return err
				}

				obtained_tool_types_map[tool_type_id] = tool_type
				recipe_operation.ToolType = tool_type
			}
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
				// Иначе получаем полуфабрикат из БД
				semiproduct, err = service.productRepo.GetByID(semiproduct_id)
				if err != nil {
					return err
				}

				obtained_products_map[semiproduct_id] = semiproduct
				recipe_semiproduct.Semiproduct = semiproduct
			}

			path_product_set[semiproduct_id] = true

			err = dfsProductRecipe(recipe_semiproduct.Semiproduct)
			if err != nil {
				return err
			}
		}

		return nil
	}

	err = dfsProductRecipe(main_product)
	if err != nil {
		switch err {
		case logic_errors.ErrCycleDetectedInProductRecipe:
			return main_product, logic_errors.ErrCycleDetectedInProductRecipe
		}

		return nil, err
	}

	return main_product, nil
}
