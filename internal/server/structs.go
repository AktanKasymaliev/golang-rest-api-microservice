package server

import (
	"base/internal/db"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Config represents the configuration settings for the server.
type Config struct {
	host 		string		`envconfig:"HOST" default:"localhost"`
	port 		string		`envconfig:"PORT" default:"8000"`
	logLevel 	string		`envconfig:"LOG_LEVEL" default:"debug"`
}

// APIServer represents the HTTP server for the API.
type APIServer struct {
	config 		*Config
	logger 		*logrus.Logger
	router 		*gin.Engine
	database	*db.Database
	health func() string 
}

func (s *APIServer) startDB() error {
	db := db.NewDatabaseInstance()

	if err := db.Open(); err != nil {
		return err
	}

	s.database = db
	return nil
}