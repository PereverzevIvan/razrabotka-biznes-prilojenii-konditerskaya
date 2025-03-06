package component_category_repo_mysql

import (
	"gorm.io/gorm"
)

type ComponentCategoryRepo struct {
	db *gorm.DB
}

func NewComponentCategoryRepo(db *gorm.DB) *ComponentCategoryRepo {
	return &ComponentCategoryRepo{
		db: db,
	}
}
