package params_tool_failure

import "time"

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
