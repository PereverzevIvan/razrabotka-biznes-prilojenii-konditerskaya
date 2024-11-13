package params

type PaginateParams struct {
	Page  int `json:"page" query:"page" default:"1"`
	Limit int `json:"limit" query:"limit" default:"10"`
}
