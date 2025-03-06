package order_params

import "time"

type UpdateStatusParams struct {
	UserID      int
	UserRole    string
	OrderID     int
	NewStatusID int                    `json:"new_status_id"`
	Data        map[string]interface{} `json:"data"`
}

func (params *UpdateStatusParams) Validate() []string {
	errs := make([]string, 0)

	if params.OrderID <= 0 {
		errs = append(errs, "order_id is required")
	}

	if params.NewStatusID <= 0 {
		errs = append(errs, "new_status_id is required")
	}

	if params.UserID <= 0 {
		errs = append(errs, "user_id is required")
	}

	if params.UserRole == "" {
		errs = append(errs, "user_role is required")
	}

	return errs
}

type UpdateStatusToCanceledData struct {
	DeclineReason string `json:"decline_reason"`
}

func (params *UpdateStatusToCanceledData) Validate() []string {
	if params == nil {
		return []string{"params is required"}
	}

	errs := make([]string, 0)
	if params.DeclineReason == "" {
		errs = append(errs, "decline_reason is required")
	}

	return errs
}

type UpdateStatusToConfirmationData struct {
	Cost                *float64   `json:"cost"`
	PlannedCompletionAt *time.Time `json:"planned_completion_at"`
}

func (params *UpdateStatusToConfirmationData) Validate() []string {
	errs := make([]string, 0)

	if params.Cost == nil {
		errs = append(errs, "cost is required")
	}

	if params.PlannedCompletionAt == nil {
		errs = append(errs, "planned_completion_at is required")
	}

	return errs
}
