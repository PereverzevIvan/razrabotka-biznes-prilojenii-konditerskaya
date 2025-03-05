
Рефакторим сущность example

## Структура

Создаём следующую структуру:

```
internal
    domains
        example
            deps/deps.go
            controller/
                controller.go
                get_by_id.go
            repo
                mysql
                    repo.go
                    get_by_id.go
            usecase
                usecase.go
                get_by_id.go
            params
                get_by_id.go
```

## Названия пакетов

> названия пакетов составляются так: 
    example_controller
    example_params
    example_usecase
    example_deps
    example_repo_mysql

## Перенос содержимого в domain/example

переносим соответствующее содержимое:
1. controllers/example -> domain/example/controller 
    + необходимые **интерфейсы** из `controllers/interfaces.go`
2. services/example -> domain/example/usecase
    + необходимые **интерфейсы** из `services/interfaces.go`
3. repos/mysql/example -> domain/example/repo/mysql
4. models/params -> domain/

## Нейминг

Работаем в domain/example/

1. controller/
    - controller.go: `IExampleService` -> `IExampleUsecase`
    - get_by_id.go: `(controller *exampleController)` -> `(c *exampleController)`
2. usecase/
    - файл `service.go` -> `usecase.go`
    - usecase.go: `exampleService` -> `ExampleUsecase`
    - get_by_id.go: `(service *ExampleUsecase)` -> `(u *ExampleUsecase)` 
3. repo/mysql/
    - repo.go: `exampleRepo` -> ``ExampleRepo`
    - get_by_id.go: `(repo *ExampleRepo)` -> `(r *ExampleRepo)`

## Удалить мьютекс из repo

Работаем в domain/example/
        
2. repo/mysql/
    - repo.go: удаляем мьютекс (`m sync.RWMutex`) из структуры `ExampleRepo` если есть 

## deps

### Создаём структуру Deps

в domains/example/deps/deps.go создаём

```go
type Deps struct {
	db *gorm.DB

	exampleRepo    *example_repo_mysql.ExampleRepo
	exampleUsecase *example_usecase.ExampleUsecase
}

func NewDeps(db *gorm.DB) *Deps {
	return &Deps{
		db: db,
	}
}

func (d *Deps) ExampleRepo() *example_repo_mysql.ExampleRepo {
	if d.exampleRepo == nil {
		d.exampleRepo = example_repo_mysql.NewExampleRepo(d.db)
	}
	return d.exampleRepo
}

func (d *Deps) ExampleUsecase() *example_usecase.ExampleUsecase {
	if d.exampleUsecase == nil {
		d.exampleUsecase = example_usecase.NewExampleUsecase(d.ExampleRepo())
	}
	return d.exampleUsecase
}
```

### Изменяем AddRoutes и используем deps в controller.go

Переходим в `domains/example/controller/controller.go`

Изменяем функцию

БЫЛО:
```go
func AddExampleControllerRoutes(api fiber.Router, exampleService IExampleService) {
    controller := &toolFailureController{
		exampleService: exampleService,
	}

    api.Get("/examples/:id", controller.GetByID)
}
``` 

---

СТАЛО:
```go
func AddRoutes(api fiber.Router, storage *storage.Storage) {
    deps := example_deps.NewDeps(storage.DB)

	controller := &toolController{
		exampleUsecase: deps.ExampleUsecase(),
	}

	apiExamples := api.Group("/examples")
	apiExamples.Get("/:id", controller.GetAll)
}
```

### Добавляем в слайс в internal/app/init_routes.go


```go
func (app *App) initRoutes() error {

	addRoutes := []AddRouteFn{
		tool_controller.AddRoutes,
        example_controller.AddRoutes, // Добавили
	}

	for _, fn := range addRoutes {
		fn(app.fiberApp, app.Storage)
	}
}
```

# ГОТОВО