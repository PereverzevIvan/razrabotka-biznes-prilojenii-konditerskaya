package purchased_component_repo_mysql

import (
	"gorm.io/gorm"
)

type PurchasedComponentRepo struct {
	db *gorm.DB
}

func NewPurchasedComponentRepo(db *gorm.DB) *PurchasedComponentRepo {
	return &PurchasedComponentRepo{
		db: db,
	}
}
