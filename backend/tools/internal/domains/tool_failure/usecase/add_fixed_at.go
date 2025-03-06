package tool_failure_usecase

import tool_failure_params "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/domains/tool_failure/params"

func (service *ToolFailureUsecase) AddFixedAt(params *tool_failure_params.AddFixedAtParams) error {
	return service.toolFailureRepo.AddFixedAt(params)
}
