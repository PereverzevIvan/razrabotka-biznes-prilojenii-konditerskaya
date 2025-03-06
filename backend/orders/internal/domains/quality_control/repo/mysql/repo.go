package quality_control_repo_mysql

import (
	"gorm.io/gorm"
)

type QualityControlRepo struct {
	db *gorm.DB
}

func NewQualityControlRepo(db *gorm.DB) *QualityControlRepo {
	return &QualityControlRepo{
		db: db,
	}
}
