package services

import (
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/models"
)

type IToolTypeRepo interface {
	GetAll() ([]models.ToolType, error)
}
