package repos_mysql_tool_failure

import (
	"sync"

	"gorm.io/gorm"
)

type toolFailureRepo struct {
	conn *gorm.DB
	m    *sync.RWMutex
}

func NewToolFailureRepo(conn *gorm.DB) *toolFailureRepo {
	return &toolFailureRepo{
		conn: conn,
		m:    &sync.RWMutex{},
	}
}
