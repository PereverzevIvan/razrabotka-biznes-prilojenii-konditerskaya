package services

import "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/gateway/models"

type IUserRepo interface {
	GetByID(user_id int) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	Create(user *models.User) error
	// GetByEmail(email string) (*models.User, error)
	// GetAll(params map[string]string) (*[]models.User, error)
	// IsAdmin(user_id int) (bool, error)
	// Create(*models.User) error
	// IsActive(user_id int) (bool, error)
	// Update(user *models.User) error
	// UpdateActive(user_id int, is_active bool) error
	// Update(*models.User) error
	// Delete(user_id int) error
}
