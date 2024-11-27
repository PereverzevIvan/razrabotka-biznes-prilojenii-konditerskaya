package services_product

import (
	"container/heap"
	"fmt"

	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models/logic_errors"
	params_tool "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models/params/tool"
)

// Вспомогательные структуры
type taskMakeSemiproducts struct {
	Semiproduct             *models.Product
	MinutesToCreate         int
	MinStartAt              int
	Amount                  int // Количество полуфабрикатов, которые нужно сделать
	CountRecipeSemiproducts int // Количество полуфабрикатов в рецепте (длина массива RecipeSemiproducts)
	ParentTasks             []*taskMakeSemiproducts
}

// Вспомогательные структуры
type taskOperation struct {
	TaskMake     *taskMakeSemiproducts
	OperationIdx int
	Operation    *models.RecipeOperation
	MinStartAt   int
	Number       int
	// DurationMinutes int
}
type taskOperationHeap []taskOperation

func (h taskOperationHeap) Len() int { return len(h) }

// Вариант 2 - Приоритизация по времени создания
// 1. Сначала пойдут те операции, которые относятся к полуфабрикату с бОльшим временем создания
// 2. Однако, если до них есть операции,
// которые можно выполнить до их начала,
// то они будут выше по приоритету и достанутся первыми из heap'ы
// 3. Если время создания полуфабриката равно,
// то операции с меньшим временем начала создания будут первее
func (h taskOperationHeap) Less(i, j int) bool {

	// fmt.Println("Less")
	if h[i].TaskMake.MinutesToCreate > h[j].TaskMake.MinutesToCreate {
		// fmt.Println(h[i].MinStartAt < h[j].MinStartAt+h[j].DurationMinutes, "a")
		return h[i].MinStartAt < h[j].MinStartAt+h[j].Operation.DurationMinutes
	}
	if h[i].TaskMake.MinutesToCreate == h[j].TaskMake.MinutesToCreate {
		return h[i].MinStartAt < h[j].MinStartAt
	}

	// fmt.Println(h[i].MinStartAt+h[i].DurationMinutes <= h[j].MinStartAt, "b")
	return h[i].MinStartAt+h[i].Operation.DurationMinutes <= h[j].MinStartAt
}

func (h taskOperationHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *taskOperationHeap) Push(x interface{}) { *h = append(*h, x.(taskOperation)) }

func (h *taskOperationHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// Вспомогательные структуры
type toolAvailability struct {
	ToolID      int
	AvailableAt int
}

type toolAvailabilityMinHeap []toolAvailability

func (h toolAvailabilityMinHeap) Len() int { return len(h) }

func (h toolAvailabilityMinHeap) Less(i, j int) bool {
	return h[i].AvailableAt < h[j].AvailableAt
}

func (h toolAvailabilityMinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *toolAvailabilityMinHeap) Push(x interface{}) { *h = append(*h, x.(toolAvailability)) }

func (h *toolAvailabilityMinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

/*
Вариант реализации алгоритма:

# Получить всё необходимое оборудование

# Создать полуфабрикаты и продукт
1. Находим все конечные полуфабрикаты - листы в графе рецепта
2. Создаём из них задачи (taskMakeSemiproducts) и добавляем в maxHeap
  - MaxHeap сортирует в порядке убывания времени создания 1го полуфабриката
  - MaxHeap будет содержать доступные для создания полуфабрикаты

2. Достаём из maxHeap задачу на создание полуфабрикатов
3. Когда задача выполнена, проходимся по ParentTasks
  - для них уменьшаем CountRecipeSemiproducts и если она равна 0,
    то добавляем в maxHeap добавляем задачи на их создание

Нужно:
- Отслеживать выполнение очередного полуфабриката
  - Когда он сделан -> проверить можно ли создать родительский полуфабрикат

maxHeap по времени создания полуфабриката

для полуфабриката нужно создать дочерние полуфабрикаты + выполнить его операции

Для задачи на создание полуфабрикатов нужно:
-
*/
func (service *productService) CalcProductionMinTime(product *models.Product) (
	work_intervals_by_tool map[int][]models.ToolWorkInterval,
	err error,
) {

	// hashset необходимых типов инструментов
	required_tool_types := initRequiredToolTypes(product)

	// ключ - tool_type_id, значение - все исправные инструменты
	tools_by_tool_type, err := service.initToolsByToolType(required_tool_types)
	if err != nil {
		return nil, err
	}

	// ключ - tool_type_id, значение - minHeap по времени доступности инструментов
	tool_availability_heaps_by_tool_type := initToolAvailabilityHeapsByToolType(tools_by_tool_type)

	// Всё задачи, ключ - semiproduct_id, значение - taskMakeSemiproducts
	tasks_by_semiproduct := initTasksBySemiproduct(product)
	for _, task := range tasks_by_semiproduct {
		if len(task.Semiproduct.RecipeOperations) == 0 {
			return nil, fmt.Errorf("%w in semiproduct %s, id: %v",
				logic_errors.ErrNoRecipeOperationsInProduct,
				task.Semiproduct.Name,
				task.Semiproduct.ID,
			)
		}
	}

	task_operation_heap := &taskOperationHeap{}

	// Добавляем задачи на создание полуфабрикатов,
	// которые не содержат другие полуфабрикаты в рецепте
	for _, task := range tasks_by_semiproduct {
		if task.CountRecipeSemiproducts == 0 {
			for i := 0; i < task.Amount; i++ {
				heap.Push(task_operation_heap, taskOperation{
					TaskMake:     task,
					OperationIdx: 0,
					Operation:    task.Semiproduct.RecipeOperations[0],
					MinStartAt:   0,
					Number:       i + 1,
				})
			}
		}
	}

	// work_intervals_by_tool - это наш ответ,
	// который содержит интервалы работы каждого инструмента.
	// По нему можно будет построить диаграмму ганта
	work_intervals_by_tool = make(map[int][]models.ToolWorkInterval)

	for task_operation_heap.Len() > 0 {

		task_operation := heap.Pop(task_operation_heap).(taskOperation)

		operation := task_operation.Operation

		operation_end_at := task_operation.MinStartAt + operation.DurationMinutes

		tool_type_id := operation.ToolTypeID

		// Обработать случай, когда надо делать без оборудования
		if tool_type_id == nil {
			// Обработать случай, когда надо делать без оборудования
			work_intervals_by_tool[0] = append(work_intervals_by_tool[0], models.ToolWorkInterval{
				Start:       task_operation.MinStartAt,
				End:         operation_end_at,
				Description: describeToolWorkInterval(&task_operation),
			})
		} else {

			tool_availability := heap.Pop(tool_availability_heaps_by_tool_type[*tool_type_id]).(toolAvailability)

			// Если инструмент не доступен для выполнения операции
			// добавляем кё обратно в очередь, с MinStartAt равным AvailableAt
			if task_operation.MinStartAt < tool_availability.AvailableAt {
				task_operation.MinStartAt = tool_availability.AvailableAt
				heap.Push(task_operation_heap, task_operation)
				heap.Push(tool_availability_heaps_by_tool_type[*tool_type_id], tool_availability)
				continue
			}

			tool_id := tool_availability.ToolID

			work_interval := models.ToolWorkInterval{
				Start:       task_operation.MinStartAt,
				End:         operation_end_at,
				Description: describeToolWorkInterval(&task_operation),
			}

			// Добавляем интервал работы инструмента
			work_intervals_by_tool[tool_id] = append(work_intervals_by_tool[tool_id], work_interval)

			// Обновляем доступность инструмента и добавляем его в очередь
			tool_availability.AvailableAt = operation_end_at
			heap.Push(tool_availability_heaps_by_tool_type[*tool_type_id], tool_availability)
		}

		task_make := task_operation.TaskMake

		// Добавляем следущую операцию в очередь, если она не последняя
		if task_operation.OperationIdx+1 < len(task_make.Semiproduct.RecipeOperations) {

			next_operation := taskOperation{
				TaskMake:     task_make,
				OperationIdx: task_operation.OperationIdx + 1,
				Operation:    task_make.Semiproduct.RecipeOperations[task_operation.OperationIdx+1],
				MinStartAt:   operation_end_at,
				Number:       task_operation.Number,
			}

			heap.Push(task_operation_heap, next_operation)
			continue
		}

		task_make.Amount--
		if task_make.Amount > 0 {
			continue
		}

		// Все полуфабрикаты этого типа созданы, обрабатываем создание родительских
		for _, parent_task := range task_make.ParentTasks {
			parent_task.CountRecipeSemiproducts--

			parent_task.MinStartAt = max(parent_task.MinStartAt, operation_end_at)

			if parent_task.CountRecipeSemiproducts > 0 {
				continue
			}

			// Все необходимые полуфабрикаты для создания родительского готовы,
			// переходим к добавлению операций для создания родительских
			for i := 0; i < parent_task.Amount; i++ {
				heap.Push(task_operation_heap, taskOperation{
					TaskMake:     parent_task,
					OperationIdx: 0,
					Operation:    parent_task.Semiproduct.RecipeOperations[0],
					MinStartAt:   parent_task.MinStartAt,
					Number:       i + 1,
				})
			}
		}
	}

	return work_intervals_by_tool, nil
}

func describeToolWorkInterval(task_operation *taskOperation) string {
	return fmt.Sprintf("%s - %v - %s",
		task_operation.TaskMake.Semiproduct.Name,
		task_operation.Number,
		task_operation.Operation.Desctiption,
	)
}

func SumOfProductRecipeOperationsDuration(product *models.Product) int {
	if len(product.RecipeOperations) == 0 {
		return 0
	}

	var sum_duratation_minutes int

	for _, operation := range product.RecipeOperations {
		sum_duratation_minutes += operation.DurationMinutes
	}

	return sum_duratation_minutes
}

func initRequiredToolTypes(main_product *models.Product) map[int]bool {
	required_tool_types := make(map[int]bool)

	var dfsProductRecipe func(product *models.Product)
	dfsProductRecipe = func(product *models.Product) {

		for _, recipe_operation := range product.RecipeOperations {
			if recipe_operation.ToolType != nil {
				required_tool_types[recipe_operation.ToolType.ID] = true
			}
		}

		// Получаем полуфабрикаты текущего продукта,
		// Продолжаем рекурсию
		for _, recipe_semiproduct := range product.RecipeSemiProducts {
			dfsProductRecipe(recipe_semiproduct.Semiproduct)
		}
	}

	dfsProductRecipe(main_product)
	return required_tool_types
}

func (service *productService) initToolsByToolType(required_tool_types map[int]bool) (map[int][]models.Tool, error) {

	tool_type_tools_map := make(map[int][]models.Tool)

	for tool_type_id := range required_tool_types {

		tools, err := service.toolRepo.GetAll(
			&params_tool.GetAllParams{
				ToolTypeID:      tool_type_id,
				OnlyServiceable: true,
			})
		if err != nil {
			return nil, err
		}

		if len(tools) == 0 {
			return nil, fmt.Errorf("%w: tool_type_id: %d",
				logic_errors.ErrNoAvailableTools,
				tool_type_id)
		}

		tool_type_tools_map[tool_type_id] = tools
	}

	return tool_type_tools_map, nil
}

func initToolAvailabilityHeapsByToolType(
	tool_type_tools_map map[int][]models.Tool,
) (tool_availability_heaps_by_tool_type map[int]*toolAvailabilityMinHeap) {

	tool_availability_heaps_by_tool_type = make(map[int]*toolAvailabilityMinHeap)

	for tool_type_id, tools := range tool_type_tools_map {

		tools_availabilities_heap := toolAvailabilityMinHeap{}
		for _, tool := range tools {
			tools_availabilities_heap = append(
				tools_availabilities_heap,
				toolAvailability{
					ToolID:      tool.ID,
					AvailableAt: 0,
				})
		}

		heap.Init(&tools_availabilities_heap)

		tool_availability_heaps_by_tool_type[tool_type_id] = &tools_availabilities_heap
	}

	return tool_availability_heaps_by_tool_type
}

func initTasksBySemiproduct(main_product *models.Product) map[int]*taskMakeSemiproducts {
	tasks_by_semiproduct := make(map[int]*taskMakeSemiproducts)

	var dfsProductRecipe func(parent_task *taskMakeSemiproducts, product *models.Product, amount int)
	dfsProductRecipe = func(parent_task *taskMakeSemiproducts, product *models.Product, amount int) {

		var task *taskMakeSemiproducts

		// Создаём задачу, или получаем уже существующую
		if task_in_map, ok := tasks_by_semiproduct[product.ID]; ok {
			task = task_in_map
		} else {
			task = &taskMakeSemiproducts{
				Semiproduct:             product,
				MinutesToCreate:         SumOfProductRecipeOperationsDuration(product),
				Amount:                  0,
				CountRecipeSemiproducts: len(product.RecipeSemiProducts),
				ParentTasks:             []*taskMakeSemiproducts{},
			}
			tasks_by_semiproduct[product.ID] = task
		}

		// Обновляем данные задачи
		task.Amount += amount
		// Отслеживание родительских задач
		if parent_task != nil {
			task.ParentTasks = append(task.ParentTasks, parent_task)
		}

		// Получаем полуфабрикаты текущего продукта,
		// Продолжаем рекурсию
		for _, recipe_semiproduct := range product.RecipeSemiProducts {
			dfsProductRecipe(
				task,
				recipe_semiproduct.Semiproduct,
				amount*recipe_semiproduct.Count,
			)
		}
	}

	dfsProductRecipe(nil, main_product, 1)

	return tasks_by_semiproduct
}
