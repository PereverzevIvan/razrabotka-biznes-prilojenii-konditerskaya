package quality_control_controller

import (
	quality_control_deps "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/domains/quality_control/deps"
	quality_control_params "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/domains/quality_control/params"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/models"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/storage"
	"github.com/gofiber/fiber/v3"
)

type IJWTUsecase interface {
	GetRoleFromAccessTokenCookie(ctx fiber.Ctx) (string, error)
	GetUserIDFromAccessTokenCookie(ctx fiber.Ctx) (int, error)
}

type IQualityControlUsecase interface {
	Create(params *quality_control_params.CreateParams) (*models.QualityControl, error)
	Update(params *quality_control_params.UpdateParams) (*models.QualityControl, error)
	GetByOrderID(orderID int) ([]models.QualityControl, error)
}

type qualityControllController struct {
	jwtUsecase            IJWTUsecase
	qualityControlUsecase IQualityControlUsecase
}

func AddRoutes(
	api fiber.Router,
	storage *storage.Storage,
) {

	deps := quality_control_deps.NewDeps(
		storage.DB,
		storage.JwtConfig,
	)

	controller := qualityControllController{
		jwtUsecase:            deps.JwtService(),
		qualityControlUsecase: deps.QualityControlUsecase(),
	}

	apiQualityControls := api.Group("/quality-controls")

	apiQualityControls.Route("/").
		Post(controller.Create).
		Patch(controller.Update)

	apiQualityControls.Get("/order/:order_id", controller.GetByOrderId)
}
