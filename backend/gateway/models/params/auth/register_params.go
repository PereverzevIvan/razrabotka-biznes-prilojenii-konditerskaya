package params_auth

import (
	"fmt"

	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/gateway/models"
)

type RegisterParams struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	// Birthday  time.Time `json:"birthday"`
}

func (params *RegisterParams) Validate() error {
	if params.Email == "" {
		return fmt.Errorf("email can't be empty")
	}
	if params.Password == "" {
		return fmt.Errorf("password can't be empty")
	}
	return nil
}

func (params *RegisterParams) ToUser() *models.User {
	return &models.User{
		ID:       0,
		RoleID:   1,
		Email:    params.Email,
		Password: params.Password,
		// FirstName: params.FirstName,
		// LastName:  params.LastName,
		// Birthday:  params.Birthday,
		// Active: true,
	}
}
