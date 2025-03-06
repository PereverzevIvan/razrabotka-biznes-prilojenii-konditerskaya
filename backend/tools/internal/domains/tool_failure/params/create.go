package tool_failure_params

import "time"

type CreateParams struct {
	ToolID   int `json:"tool_id"`
	MasterID int `json:"master_id"`

	FailureReason   string `json:"reason"`
	FailureReasonID *int   `json:"reason_id"`

	FailureAt *time.Time `json:"failure_at"`
}

func (params *CreateParams) Validate() []string {
	errs := make([]string, 0)

	if params.ToolID == 0 {
		errs = append(errs, "tool_id is required")
	}

	if params.FailureReason == "" && params.FailureReasonID == nil {
		errs = append(errs, "reason or reason_id is required")
	}

	if params.FailureReason != "" && params.FailureReasonID != nil {
		errs = append(errs, "reason and reason_id can't be set at the same time")
	}

	if params.FailureReasonID != nil && *params.FailureReasonID == 0 {
		errs = append(errs, "reason_id can't be 0, but it can be null")
	}

	return errs
}
