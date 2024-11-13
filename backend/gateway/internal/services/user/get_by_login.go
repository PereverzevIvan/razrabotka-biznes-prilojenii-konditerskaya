package services_user

import "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/gateway/models"

func (s *UserService) GetByLogin(email string) (*models.User, error) {
	return s.userRepo.GetByLogin(email)
}
