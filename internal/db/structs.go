package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Config config struct
type Config struct {
	user 		string
	name 		string
	password 	string
	host		string
	port 		string

	// Optional
	sslmode		string
	timezone 	string
}

func (db *Config) makeDSN() string {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		db.host, db.user, db.password, db.name, db.port, db.sslmode, db.timezone,
	)
	return dsn
}

// Database struct
type Database struct {
	dbConfig 	*Config
	database	*gorm.DB
}

// Open let you open the connnection to postgresql
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