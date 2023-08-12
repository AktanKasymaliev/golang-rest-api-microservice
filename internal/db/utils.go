package db

import "base/internal/core"


func newDBConfigInstance() *DBConfig {
	return &DBConfig{
		user: core.Getenv("DB_USER", "user"),
		name: core.Getenv("DB_NAME", "name"),
		password: core.Getenv("DB_PASSWORD", "password"),
		host: core.Getenv("DB_HOST", "localhost"),
		port: core.Getenv("DB_PORT", "5432"),
		sslmode: core.Getenv("DB_SSLMODE", "disable"),
		timezone: core.Getenv("DB_TIMEZONE", "UTC"),
	}
}

func NewDatabaseInstance() *Database {
	return &Database{
		dbConfig: 	newDBConfigInstance(),
	}
}