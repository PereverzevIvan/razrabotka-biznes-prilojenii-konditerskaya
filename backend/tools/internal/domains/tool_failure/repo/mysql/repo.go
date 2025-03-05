package tool_failure_repo_mysql

import (
	"gorm.io/gorm"
)

type ToolFailureRepo struct {
	conn *gorm.DB
}

func NewToolFailureRepo(conn *gorm.DB) *ToolFailureRepo {
	return &ToolFailureRepo{
		conn: conn,
	}
}
