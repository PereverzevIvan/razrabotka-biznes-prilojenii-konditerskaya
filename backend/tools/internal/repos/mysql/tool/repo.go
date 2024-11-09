package repos_mysql_tool

import (
	"sync"

	"gorm.io/gorm"
)

type toolRepo struct {
	conn *gorm.DB
	m    *sync.RWMutex
}

func NewToolRepo(conn *gorm.DB) *toolRepo {
	return &toolRepo{
		conn: conn,
		m:    &sync.RWMutex{},
	}
}
