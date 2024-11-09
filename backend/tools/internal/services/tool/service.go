package services_tool

import (
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/services"
)

type toolService struct {
	toolRepo services.IToolRepo
}

func NewToolService(toolRepo services.IToolRepo) *toolService {
	return &toolService{
		toolRepo: toolRepo,
	}
}
