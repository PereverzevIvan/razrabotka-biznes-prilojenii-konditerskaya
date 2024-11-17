package main

import (
	"database/sql"
	"fmt"

	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/gateway/configs"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/gateway/pkg/config_loader"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	main_cfg := configs.Config{}
	config_loader.MustLoad("./configs/default_config.yaml", main_cfg)

	cfg := main_cfg.DBConfig

	var dsn string

	switch cfg.DBType {
	case "mysql":
		dsn = fmt.Sprintf(
			"%v:%v@tcp(%v:%v)/%v?multiStatements=true",
			cfg.AdminName,
			cfg.DBPassword,
			cfg.Host,
			cfg.Port,
			cfg.DBName,
		)
	default:
		panic("Неверный тип базы данных")
	}

	db, err := sql.Open(cfg.DBType, dsn)

	if err != nil {
		panic(err.Error())
	}

	var driver database.Driver

	switch cfg.DBType {
	case "mysql":
		driver, err = mysql.WithInstance(db, &mysql.Config{})
	default:
		panic("Неверный тип базы данных")

	}

	if err != nil {
		panic(err.Error())
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://./storage/migrations",
		"confectionary",
		driver,
	)
	if err != nil {
		panic(err.Error())
	}

	// m.Down()
	err = m.Up()
	if err != nil {
		fmt.Println(err.Error())
	}
	// m.Steps()
}
