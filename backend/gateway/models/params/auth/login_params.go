package params_auth

type LoginParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
