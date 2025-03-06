package component_type_controller

import (
	component_type_deps "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/domains/component_type/deps"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/storage"
	"github.com/gofiber/fiber/v3"
)

type IComponentCategoryUsecase interface {
	GetByName(name string) (*models.ComponentCategory, error)
}

type IComponentTypeUsecase interface {
	GetAll(component_category_id int) ([]models.ComponentType, error)
}

type componentTypeController struct {
	// componentCategoriesMap   map[string]int
	componentCategoryUsecase IComponentCategoryUsecase
	componentTypeUsecase     IComponentTypeUsecase
}

func AddRoutes(
	api fiber.Router,
	storage *storage.Storage,
) {
	deps := component_type_deps.NewDeps(storage.DB)

	controller := componentTypeController{
		componentCategoryUsecase: deps.ComponentCategoryUsecase(),
		componentTypeUsecase:     deps.ComponentTypeUsecase(),
	}

	api.Get(":category_name/types", controller.GetAll)
}
