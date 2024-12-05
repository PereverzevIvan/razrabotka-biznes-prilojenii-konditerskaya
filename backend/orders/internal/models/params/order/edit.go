package params_order

type EditParams struct {
	ID                  int      `json:"id"`
	Name                string   `json:"name"`
	Description         *string  `json:"description"`
	Cost                *float64 `json:"cost"`
	PlannedCompletionAt *string  `json:"planned_completion_at"`
}
