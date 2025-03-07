package tool_failure_params

import (
	"time"

	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/proto/pkg/api/tool_failure"
)

type AddFixedAtParams struct {
	ToolFailureID int        `json:"tool_failure_id"`
	FixedAt       *time.Time `json:"fixed_at"`
}

func (p *AddFixedAtParams) Validate() []string {
	errs := make([]string, 0)

	if p.ToolFailureID == 0 {
		errs = append(errs, "tool_failure_id is required")
	}

	if p.FixedAt == nil {
		errs = append(errs, "fixed_at is required")
	}

	return errs
}

func AddFixedAddParamsFromGRPC(req *tool_failure.ToolFailureAddFixedAtRequest) *AddFixedAtParams {
	params := &AddFixedAtParams{
		ToolFailureID: int(req.ToolFailureId),
	}
	if req.FixedAt != nil {
		fixedAt := req.FixedAt.AsTime()
		params.FixedAt = &fixedAt
	}
	return params
}
