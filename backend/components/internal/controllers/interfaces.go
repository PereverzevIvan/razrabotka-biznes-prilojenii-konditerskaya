package controllers

import (
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models"
	params_purchased_component "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models/params/purchased_component"
	results_purchased_component "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models/results/purchased_component"
	"github.com/gofiber/fiber/v3"
)

type IJWTService interface {
	GetRoleFromAccessTokenCookie(ctx fiber.Ctx) (string, error)
	GetUserIDFromAccessTokenCookie(ctx fiber.Ctx) (int, error)
}

type IComponentCategoryService interface {
	GetByName(name string) (*models.ComponentCategory, error)
}

type IComponentTypeService interface {
	GetAll(component_category_id int) ([]models.ComponentType, error)
}

type IComponentService interface {
	GetAll() ([]models.Component, error)
	GetByID(component_id int) (*models.Component, error)
}

type IPurchasedComponentService interface {
	GetAll(component_category_id int, params *params_purchased_component.GetAllParams) (*results_purchased_component.GetAllResults, error)
}