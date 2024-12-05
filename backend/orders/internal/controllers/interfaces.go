package controllers

import (
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/models"
	params_order "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/models/params/order"
	"github.com/gofiber/fiber/v3"
)

type IJWTService interface {
	GetRoleFromAccessTokenCookie(ctx fiber.Ctx) (string, error)
	GetUserIDFromAccessTokenCookie(ctx fiber.Ctx) (int, error)
}

type IOrderService interface {
	GetAll(params *params_order.GetAllParams) ([]models.Order, error)
	Create(params *params_order.CreateParams) (*models.Order, error)
	// GetByName(name string) (*models.ComponentCategory, error)
}
