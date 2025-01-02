package params_quality_control

type CreateParams struct {
	MasterID    int     `json:"-"`
	OrderID     int     `json:"order_id"`
	IsSatisfies bool    `json:"is_satisfies"`
	Comment     *string `json:"comment"`
	Parameter   string  `json:"parameter"`
}

func (p *CreateParams) Validate() []string {
	errs := make([]string, 0)

	if p.MasterID == 0 {
		errs = append(errs, "master_id is required")
	}

	if p.OrderID == 0 {
		errs = append(errs, "order_id is required")
	}

	if p.Parameter == "" {
		errs = append(errs, "parameter is required")
	}

	if !p.IsSatisfies && p.Comment != nil && *p.Comment == "" {
		errs = append(errs, "comment is required if is_satisfies is false")
	}

	return errs
}
