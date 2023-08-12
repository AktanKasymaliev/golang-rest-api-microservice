package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConfig struct {
	user 		string
	name 		string
	password 	string
	host		string
	port 		string

	// Optional
	sslmode		string
	timezone 	string
}

func (db *DBConfig) makeDSN() string {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		db.host, db.user, db.password, db.name, db.port, db.sslmode, db.timezone,
	)
	return dsn
}


type Database struct {
	dbConfig 	*DBConfig
	database	*gorm.DB
}

func (database *Database) Open() error {
	db, err := gorm.Open(
		postgres.Open(database.dbConfig.makeDSN()),
	)
	if err != nil {
		return err
	}
	database.database = db
	return nil
}