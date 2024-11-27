package repos_mysql_tool

import (
	"sync"

	"gorm.io/gorm"
)

type repoTools struct {
	conn *gorm.DB
	m    *sync.RWMutex
}

func NewToolRepo(conn *gorm.DB) *repoTools {
	return &repoTools{
		conn: conn,
		m:    &sync.RWMutex{},
	}
}
