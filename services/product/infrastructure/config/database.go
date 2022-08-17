package config

import (
	"database/sql"
	"fmt"
	"go-commerce/pkg/config"

	_ "github.com/go-sql-driver/mysql"
)

func InitMysqlConfig() *sql.DB {
	cfg := config.DatabaseMysql()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}
