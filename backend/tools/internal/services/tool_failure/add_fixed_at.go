package services_tool_failure

import (
	params_tool_failure "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/models/params/tool_failure"
)

func (service *toolFailureService) AddFixedAt(params *params_tool_failure.AddFixedAtParams) error {
	return service.toolFailureRepo.AddFixedAt(params)
}
