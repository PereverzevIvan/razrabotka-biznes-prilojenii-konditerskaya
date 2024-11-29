package services

import (
	"time"

	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/models"
)

type IUserRepo interface {
	GetByID(id int) (*models.User, error)
}

type IOrderRepo interface {
	GetAll() ([]models.Order, error)
	Create(order *models.Order) (*models.Order, error)
	CountForDate(date time.Time) (int64, error)
}
