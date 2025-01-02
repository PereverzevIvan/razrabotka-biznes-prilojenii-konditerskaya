package controllers_order

import (
	"errors"

	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/models/logic_errors"
	params_order "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/models/params/order"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

func (controller *orderController) UpdateStatus(ctx fiber.Ctx) error {
	user_role, err := controller.jwtService.GetRoleFromAccessTokenCookie(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	user_id, err := controller.jwtService.GetUserIDFromAccessTokenCookie(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	update_params := &params_order.UpdateStatusParams{}
	if err = ctx.Bind().Body(update_params); err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	update_params.OrderID = fiber.Params[int](ctx, "order_id")
	update_params.UserID = user_id
	update_params.UserRole = user_role

	if errs := update_params.Validate(); len(errs) != 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"errors": errs})
	}

	order, err := controller.orderService.UpdateStatus(update_params)
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
