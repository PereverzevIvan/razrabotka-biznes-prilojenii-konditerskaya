package controllers

import "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/gateway/models"

type IUserService interface {
	// GetUser(userId int) (string, error)
	GetByID(user_id int) (*models.User, error)
}
