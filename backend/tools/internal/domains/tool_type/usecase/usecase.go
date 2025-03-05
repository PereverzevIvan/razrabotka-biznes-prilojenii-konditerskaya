package tool_type_usecase

import "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/services"

type ToolTypeUsecase struct {
	toolTypeRepo services.IToolTypeRepo
}

func NewToolTypeUsecase(toolTypeRepo services.IToolTypeRepo) *ToolTypeUsecase {
	return &ToolTypeUsecase{
		toolTypeRepo: toolTypeRepo,
	}
}
