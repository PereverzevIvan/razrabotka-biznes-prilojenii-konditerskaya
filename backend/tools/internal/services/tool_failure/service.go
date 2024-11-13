package services_tool_failure

import "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/services"

type toolFailureService struct {
	toolFailureRepo services.IToolFailureRepo
}

func NewToolFailureService(toolFailureRepo services.IToolFailureRepo) *toolFailureService {
	return &toolFailureService{
		toolFailureRepo: toolFailureRepo,
	}
}
