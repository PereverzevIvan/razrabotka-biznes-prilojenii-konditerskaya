package order_controller

import (
	order_params "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/domains/order/params"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

func (c *orderController) GetAll(ctx fiber.Ctx) error {
	user_role, err := c.jwtService.GetRoleFromAccessTokenCookie(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	user_id, err := c.jwtService.GetUserIDFromAccessTokenCookie(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	params := &order_params.GetAllParams{
		UserId:   user_id,
		UserRole: user_role,
	}
	if errs := params.Validate(); len(errs) != 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"errors": errs})
	}

	orders, err := c.orderUsecase.GetAll(params)
	if err != nil {
		log.Error(err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.Status(fiber.StatusOK).JSON(orders)
}
