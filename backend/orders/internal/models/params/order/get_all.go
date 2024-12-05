package params_order

type GetAllParams struct {
	UserId   int    `json:"author_id"`
	UserRole string `json:"author_role"`
}

func (params *GetAllParams) Validate() []string {
	errs := make([]string, 0)

	if params.UserId == 0 {
		errs = append(errs, "author_id is required")
	}
	if params.UserRole == "" {
		errs = append(errs, "author_role is required")
	}

	return errs
}
