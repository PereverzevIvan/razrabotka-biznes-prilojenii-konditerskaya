package services

import "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/models"

type IOrderRepo interface {
	GetAll() ([]models.Order, error)
}
