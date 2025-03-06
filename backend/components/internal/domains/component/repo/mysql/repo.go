package component_repo_mysql

import (
	"gorm.io/gorm"
)

type ComponentRepo struct {
	db *gorm.DB
}

func NewComponentRepo(db *gorm.DB) *ComponentRepo {
	return &ComponentRepo{
		db: db,
	}
}
