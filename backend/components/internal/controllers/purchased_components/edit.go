package controllers_purchased_component

import (
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

func (controller *purchasedComponentController) Edit(ctx fiber.Ctx) error {
	purchased_component := &models.PurchasedComponent{}

	err := ctx.Bind().Body(purchased_component)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	err = controller.purchasedComponentService.Edit(purchased_component)
	if err != nil {
		log.Error(err)
		switch err {
		case models.ErrFK:
			return ctx.Status(fiber.StatusBadRequest).SendString(models.ErrFK.Error())
		}
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.Status(fiber.StatusOK).JSON(purchased_component)
}
