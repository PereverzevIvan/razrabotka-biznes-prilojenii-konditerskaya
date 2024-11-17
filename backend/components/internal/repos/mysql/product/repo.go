package repos_mysql_product

import (
	"sync"

	"gorm.io/gorm"
)

type productRepo struct {
	conn *gorm.DB
	m    *sync.RWMutex
}

func NewProductRepo(conn *gorm.DB) *productRepo {
	return &productRepo{
		conn: conn,
		m:    &sync.RWMutex{},
	}
}
