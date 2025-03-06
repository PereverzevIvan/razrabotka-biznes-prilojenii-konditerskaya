package storage

import (
	"fmt"

	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/configs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB(dbConfig *configs.DBConfig) (*gorm.DB, error) {
	var db *gorm.DB
	var err error
	var dsn string

	switch dbConfig.DBType {
	case "mysql":
		dsn = fmt.Sprintf(
			"%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
			dbConfig.AdminName,
			dbConfig.DBPassword,
			dbConfig.Host,
			dbConfig.Port,
			dbConfig.DBName,
		)
		fmt.Println(dsn)

		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			return nil, fmt.Errorf("не удалось подключиться к базе данных, причина: %v", err)
		}
	// case "postgres":
	// 	dsn = fmt.Sprintf(
	// 		"host=%v port=%v user=%v password=%v dbname=%v sslmode=%v",
	// 		cfg.Host,
	// 		cfg.Port,
	// 		cfg.AdminName,
	// 		cfg.DBPassword,
	// 		cfg.DBName,
	// 		cfg.SSL,
	// 	)

	// 	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// 	if err != nil {
	// 		panic("Не удалось подключиться к базе данных")
	// 	}
	default:
		return nil, fmt.Errorf("неизвестный тип базы данных")
	}

	return db, nil
}
