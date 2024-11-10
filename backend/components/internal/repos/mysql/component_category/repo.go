package repos_mysql_component_category

import (
	"sync"

	"gorm.io/gorm"
)

type componentCategoryRepo struct {
	conn *gorm.DB
	m    *sync.RWMutex
}

func NewComponentCategoryRepo(conn *gorm.DB) *componentCategoryRepo {
	return &componentCategoryRepo{
		conn: conn,
		m:    &sync.RWMutex{},
	}
}
