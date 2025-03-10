package order_controller

import (
	order_params "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/domains/order/params"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/models"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

func (c *orderController) Create(ctx fiber.Ctx) error {
	user_role, err := c.jwtService.GetRoleFromAccessTokenCookie(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	user_id, err := c.jwtService.GetUserIDFromAccessTokenCookie(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	params := &order_params.CreateParams{}
	if err = ctx.Bind().Body(params); err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	params.AuthorId = user_id
	params.AuthorRole = user_role

	switch user_role {
	case models.KRoleCustomer.Name():
		params.CustomerID = user_id
		params.PlannedCompletionAt = nil
	case models.KRoleClientManager.Name():
		params.ManagerID = &user_id
	}

	if errs := params.Validate(); len(errs) != 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"errors": errs})
	}

	order, err := c.orderUsecase.Create(params)
	if err != nil {
		log.Error(err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.Status(fiber.StatusCreated).JSON(order)
}
