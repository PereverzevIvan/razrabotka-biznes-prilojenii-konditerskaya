package storage

import (
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/configs"
	"gorm.io/gorm"
)

type Storage struct {
	DB              *gorm.DB
	JwtConfig       *configs.JWTConfig
	ComponentConfig *configs.ComponentsAPIConfig
}

func NewStorage(
	dbConfig *configs.DBConfig,
	jwtConfig *configs.JWTConfig,
	componentConfig *configs.ComponentsAPIConfig,
) (*Storage, error) {
	storage := &Storage{
		JwtConfig:       jwtConfig,
		ComponentConfig: componentConfig,
	}

	var err error
	storage.DB, err = NewDB(dbConfig)
	if err != nil {
		return nil, err
	}

	return storage, nil
}
