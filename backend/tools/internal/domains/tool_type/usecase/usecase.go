package tool_type_usecase

import (
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/models"
)

type IToolTypeRepo interface {
	GetAll() ([]models.ToolType, error)
}

type ToolTypeUsecase struct {
	toolTypeRepo IToolTypeRepo
}

func NewToolTypeUsecase(toolTypeRepo IToolTypeRepo) *ToolTypeUsecase {
	return &ToolTypeUsecase{
		toolTypeRepo: toolTypeRepo,
	}
}
