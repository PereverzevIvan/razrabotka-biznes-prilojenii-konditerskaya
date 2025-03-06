package quality_control_params

type UpdateParams struct {
	ID          int     `json:"id"`
	IsSatisfies *bool   `json:"is_satisfies"`
	Comment     string  `json:"comment"`
	Parameter   *string `json:"parameter"`
}

func (p *UpdateParams) Validate() []string {
	errs := make([]string, 0)

	if p.ID == 0 {
		errs = append(errs, "id is required")
	}

	if p.IsSatisfies != nil &&
		!*p.IsSatisfies &&
		p.Comment == "" {
		errs = append(errs, "comment is required if is_satisfies is false")
	}

	if p.Parameter != nil && *p.Parameter == "" {
		errs = append(errs, "parameter is required")
	}

	return errs
}
