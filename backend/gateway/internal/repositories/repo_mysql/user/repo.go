package repo_mysql_user

import (
	"sync"

	"gorm.io/gorm"
)

type userRepo struct {
	Conn *gorm.DB
	m    *sync.RWMutex
}

func NewUserRepo(conn *gorm.DB) *userRepo {
	return &userRepo{
		Conn: conn,
		m:    &sync.RWMutex{},
	}
}
