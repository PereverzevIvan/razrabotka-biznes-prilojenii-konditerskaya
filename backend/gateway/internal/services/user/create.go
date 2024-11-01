package services_user

import (
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/gateway/models"
	params_auth "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/gateway/models/params/auth"
	"golang.org/x/crypto/bcrypt"
)

func (s *UserService) Create(params *params_auth.RegisterParams) (*models.User, error) {
	user := params.ToUser()

	user.RoleID = 1

	hashed_password_bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user.Password = string(hashed_password_bytes)

	return user, s.userRepo.Create(user)
}
