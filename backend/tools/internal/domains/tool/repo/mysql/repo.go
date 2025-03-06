package tool_repo_mysql

import (
	"gorm.io/gorm"
)

type ToolRepo struct {
	db *gorm.DB
}

func NewToolRepo(db *gorm.DB) *ToolRepo {
	return &ToolRepo{
		db: db,
	}
}
