package storage

import (
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/configs"
	"gorm.io/gorm"
)

type Storage struct {
	DB        *gorm.DB
	JwtConfig *configs.JWTConfig
}

func NewStorage(
	dbConfig *configs.DBConfig,
	jwtConfig *configs.JWTConfig,
) (*Storage, error) {
	storage := &Storage{
		JwtConfig: jwtConfig,
	}

	var err error
	storage.DB, err = NewDB(dbConfig)
	if err != nil {
		return nil, err
	}

	return storage, nil
}
