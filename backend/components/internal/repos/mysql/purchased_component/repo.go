package repos_mysql_purchased_component

import (
	"sync"

	"gorm.io/gorm"
)

type purchasedComponentRepo struct {
	conn *gorm.DB
	m    *sync.RWMutex
}

func NewPurchasedComponentService(conn *gorm.DB) *purchasedComponentRepo {
	return &purchasedComponentRepo{
		conn: conn,
		m:    &sync.RWMutex{},
	}
}
