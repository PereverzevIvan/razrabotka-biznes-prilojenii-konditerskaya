package services_user

import "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/gateway/internal/services"

// type RoleRepo interface {
// 	GetByID(role_id int) (models.Role, error)
// }

type UserService struct {
	userRepo services.IUserRepo
}

func NewUserService(userRepo services.IUserRepo) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

// func (s UserService) GetByEmail(email string) (*models.User, error) {
// 	return s.userRepo.GetByEmail(email)
// }

// func (s UserService) GetAll(params map[string]string) (*[]models.User, error) {
// 	users, err := s.userRepo.GetAll(params)
// 	return users, err
// }

// func (s UserService) IsActive(user_id int) (bool, error) {
// 	is_active, err := s.userRepo.IsActive(user_id)
// 	if err != nil {
// 		return false, err
// 	}
// 	return is_active, err
// }

// func (s UserService) IsAdmin(user_id int) (bool, error) {
// 	is_admin, err := s.userRepo.IsAdmin(user_id)
// 	if err != nil {
// 		return false, err
// 	}
// 	return is_admin, err
// }

// func HashPassword(password string) ([]byte, error) {
// 	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
// 	if err != nil {
// 		return []byte{}, err
// 	}

// 	return hashedPassword, nil
// }

// func (s UserService) Create(user *models.User) error {
// 	hashedPassword, err := HashPassword(user.Password)
// 	if err != nil {
// 		return err
// 	}

// 	user.Password = string(hashedPassword)
// 	return s.userRepo.Create(user)
// }

// func (s UserService) Update(user *models.User) error {
// 	return s.userRepo.Update(user)
// }

// func (s UserService) UpdateActive(user_id int, is_active bool) error {
// 	return s.userRepo.UpdateActive(user_id, is_active)
// }
