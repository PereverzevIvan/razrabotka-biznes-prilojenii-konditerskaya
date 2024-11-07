package params_auth

import (
	"strings"

	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/gateway/models"
)

type RegisterParams struct {
	Login     string `json:"login"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	// Birthday  time.Time `json:"birthday"`
}

func (params *RegisterParams) Validate() []string {
	errs := make([]string, 0)

	if params.Login == "" {
		errs = append(errs, "логин не может быть пустым")
	}
	errs = append(errs, params.validatePassword()...)

	return errs
}

func (params *RegisterParams) validatePassword() []string {
	errs := make([]string, 0)
	if len(params.Password) < 5 {
		errs = append(errs, "длина пароля должна быть не менее 5 букв")
	}
	if len(params.Password) > 20 {
		errs = append(errs, "длина пароля должна быть не более 20 букв")
	}
	if strings.Contains(params.Password, params.Login) {
		errs = append(errs, "пароль не должен содержать логин")
	}

	hasLowerCaseCharacters := false
	hasUpperCaseCharacters := false
	for _, char := range params.Password {
		if char >= 'a' && char <= 'z' {
			hasLowerCaseCharacters = true
		} else if char >= 'A' && char <= 'Z' {
			hasUpperCaseCharacters = true
		}
	}
	if !hasLowerCaseCharacters {
		errs = append(errs, "пароль должен содержать строчные буквы")
	}
	if !hasUpperCaseCharacters {
		errs = append(errs, "пароль должен содержать заглавные буквы")
	}

	return errs
}

func (params *RegisterParams) ToUser() *models.User {
	return &models.User{
		ID:       0,
		RoleID:   1,
		Login:    params.Login,
		Password: params.Password,
		// FirstName: params.FirstName,
		// LastName:  params.LastName,
		// Birthday:  params.Birthday,
		// Active: true,
	}
}
