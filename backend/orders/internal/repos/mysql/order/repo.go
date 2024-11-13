package repos_mysql_order

import (
	"sync"

	"gorm.io/gorm"
)

type orderRepo struct {
	conn *gorm.DB
	m    *sync.RWMutex
}

func NewOrderRepo(conn *gorm.DB) *orderRepo {
	return &orderRepo{
		conn: conn,
		m:    &sync.RWMutex{},
	}
}
