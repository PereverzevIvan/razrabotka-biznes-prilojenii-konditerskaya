package services_user

import "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/gateway/models"

func (s *UserService) GetByID(user_id int) (*models.User, error) {
	return s.userRepo.GetByID(user_id)
}
