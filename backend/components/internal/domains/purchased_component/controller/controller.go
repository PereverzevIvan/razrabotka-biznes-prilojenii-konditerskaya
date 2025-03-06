package purchased_component_controller

import (
	purchased_component_deps "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/domains/purchased_component/deps"
	purchased_component_params "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/domains/purchased_component/params"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models"
	results_purchased_component "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models/results/purchased_component"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/storage"
	"github.com/gofiber/fiber/v3"
)

type IComponentCategoryUsecase interface {
	GetByName(name string) (*models.ComponentCategory, error)
}

type IPurchasedComponentService interface {
	GetAll(component_category_id int, params *purchased_component_params.GetAllParams) (*results_purchased_component.GetAllResults, error)
	DeleteEmpty(purchased_component_id int) error
	Create(purchased_component *models.PurchasedComponent) error
	Edit(purchased_component *models.PurchasedComponent) error
}

type purchasedComponentController struct {
	componentCategoryService  IComponentCategoryUsecase
	purchasedComponentService IPurchasedComponentService
}

func AddRoutes(
	api fiber.Router,
	storage *storage.Storage,
) {
	deps := purchased_component_deps.NewDeps(storage.DB)

	controller := purchasedComponentController{
		componentCategoryService:  deps.ComponentCategoryUsecase(),
		purchasedComponentService: deps.PurchasedComponentUsecase(),
	}

	apiPurchased := api.Group(":category_name/purchased")

	apiPurchased.Route("").
		Get(controller.GetAll).
		Post(controller.Create).
		Patch(controller.Edit)

	apiPurchased.Delete("/:purchased_component_id", controller.DeleteEmpty)
	// api.Delete(":category_name/purchased/:component_id", controller.Delete)
}
