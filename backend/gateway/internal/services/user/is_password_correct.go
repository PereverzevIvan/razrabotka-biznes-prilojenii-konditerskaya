package services_user

import (
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/gateway/models"
	"golang.org/x/crypto/bcrypt"
)

func (s *UserService) IsPasswordCorrect(user *models.User, password string) bool {
	hash_byte := []byte(user.Password)

	err := bcrypt.CompareHashAndPassword(hash_byte, []byte(password))
	return err == nil
}
