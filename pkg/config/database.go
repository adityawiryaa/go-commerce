package config

import "os"

func DatabaseMysql() DatabaseMysqlConfig {
	return DatabaseMysqlConfig{
		User:     os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_NAME"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
	}
}
