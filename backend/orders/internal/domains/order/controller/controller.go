package order_controller

import (
	order_deps "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/domains/order/deps"
	order_params "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/domains/order/params"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/models"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/storage"
	"github.com/gofiber/fiber/v3"
)

type IJWTUsecase interface {
	GetRoleFromAccessTokenCookie(ctx fiber.Ctx) (string, error)
	GetUserIDFromAccessTokenCookie(ctx fiber.Ctx) (int, error)
}
type IOrderUsecase interface {
	GetAll(params *order_params.GetAllParams) ([]models.Order, error)
	Create(params *order_params.CreateParams) (*models.Order, error)
	// GetByName(name string) (*models.ComponentCategory, error)
	UpdateStatus(params *order_params.UpdateStatusParams) (*models.Order, error)
}

type orderController struct {
	jwtService   IJWTUsecase
	orderUsecase IOrderUsecase
}

func AddRoutes(
	api fiber.Router,
	storage *storage.Storage,
) {
	deps := order_deps.NewDeps(
		storage.DB,
		storage.JwtConfig,
		storage.ComponentConfig.BaseUrl,
	)

	controller := orderController{
		jwtService:   deps.JwtService(),
		orderUsecase: deps.OrderUsecase(),
	}

	apiOrders := api.Group("/orders")

	apiOrders.Get("/", controller.GetAll)
	apiOrders.Post("/", controller.Create)
	apiOrders.Patch("/:order_id/status", controller.UpdateStatus)
}
