package repos_mysql_supplier_component

import (
	"sync"

	"gorm.io/gorm"
)

type repoSupplierComponent struct {
	conn *gorm.DB
	m    *sync.RWMutex
}

func NewSupplierComponentRepo(db *gorm.DB) *repoSupplierComponent {
	return &repoSupplierComponent{
		conn: db,
		m:    &sync.RWMutex{},
	}
}
