package tool_type_repo_mysql

import (
	"gorm.io/gorm"
)

type ToolTypeRepo struct {
	db *gorm.DB
}

func NewToolTypeRepo(db *gorm.DB) *ToolTypeRepo {
	return &ToolTypeRepo{
		db: db,
	}
}
