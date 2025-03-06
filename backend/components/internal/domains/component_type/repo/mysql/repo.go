package component_type_repo_mysql

import (
	"gorm.io/gorm"
)

type ComponentTypeRepo struct {
	db *gorm.DB
}

func NewComponentTypeRepo(db *gorm.DB) *ComponentTypeRepo {
	return &ComponentTypeRepo{
		db: db,
	}
}
