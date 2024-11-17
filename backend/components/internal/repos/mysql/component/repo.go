package repos_mysql_component

import (
	"sync"

	"gorm.io/gorm"
)

type componentRepo struct {
	conn *gorm.DB
	m    *sync.RWMutex
}

func NewComponentRepo(conn *gorm.DB) *componentRepo {
	return &componentRepo{
		conn: conn,
		m:    &sync.RWMutex{},
	}
}
