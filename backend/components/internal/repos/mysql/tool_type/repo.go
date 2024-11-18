package repos_mysql_tool_type

import (
	"sync"

	"gorm.io/gorm"
)

type toolTypeRepo struct {
	conn *gorm.DB
	m    *sync.RWMutex
}

func NewToolTypeRepo(conn *gorm.DB) *toolTypeRepo {
	return &toolTypeRepo{
		conn: conn,
		m:    &sync.RWMutex{},
	}
}
