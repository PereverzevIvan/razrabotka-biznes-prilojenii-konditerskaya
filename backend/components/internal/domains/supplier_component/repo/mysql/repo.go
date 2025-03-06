package supplier_component_repo_mysql

import (
	"gorm.io/gorm"
)

type SupplierComponentRepo struct {
	db *gorm.DB
}

func NewSupplierComponentRepo(db *gorm.DB) *SupplierComponentRepo {
	return &SupplierComponentRepo{
		db: db,
	}
}
