package repos_mysql_component_type

import (
	"sync"

	"gorm.io/gorm"
)

type componentTypeRepo struct {
	conn *gorm.DB
	m    *sync.RWMutex
}

func NewComponentTypeRepo(conn *gorm.DB) *componentTypeRepo {
	return &componentTypeRepo{
		conn: conn,
		m:    &sync.RWMutex{},
	}
}
