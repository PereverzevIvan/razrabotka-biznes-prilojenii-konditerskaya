package services_tool_type

import "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/services"

type toolTypeService struct {
	toolTypeRepo services.IToolTypeRepo
}

func NewToolTypeService(toolTypeRepo services.IToolTypeRepo) *toolTypeService {
	return &toolTypeService{
		toolTypeRepo: toolTypeRepo,
	}
}
