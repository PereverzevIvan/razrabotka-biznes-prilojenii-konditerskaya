package repos_mysql_quality_control

import (
	"sync"

	"gorm.io/gorm"
)

type qualityControlRepo struct {
	conn *gorm.DB
	m    *sync.RWMutex
}

func NewQualityControlRepo(conn *gorm.DB) *qualityControlRepo {
	return &qualityControlRepo{
		conn: conn,
		m:    &sync.RWMutex{},
	}
}
