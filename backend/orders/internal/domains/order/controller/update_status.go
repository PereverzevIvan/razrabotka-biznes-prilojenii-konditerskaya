package order_controller

import (
	"errors"

	order_params "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/domains/order/params"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/models/logic_errors"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

func (c *orderController) UpdateStatus(ctx fiber.Ctx) error {
	user_role, err := c.jwtService.GetRoleFromAccessTokenCookie(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	user_id, err := c.jwtService.GetUserIDFromAccessTokenCookie(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	update_params := &order_params.UpdateStatusParams{}
	if err = ctx.Bind().Body(update_params); err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	update_params.OrderID = fiber.Params[int](ctx, "order_id")
	update_params.UserID = user_id
	update_params.UserRole = user_role

	if errs := update_params.Validate(); len(errs) != 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"errors": errs})
	}

	order, err := c.orderUsecase.UpdateStatus(update_params)
	if err != nil {
		log.Error(err)
		switch {
		case errors.Is(err, logic_errors.ErrInvalidUpdate):
			return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.Status(fiber.StatusOK).JSON(order)
}
