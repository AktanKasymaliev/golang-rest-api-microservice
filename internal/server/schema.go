package server

import (
	"github.com/sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

// Config represents the configuration settings for the server.
type Config struct {
	Host string		`json:"host" envconfig:"HOST" default:"localhost"`
	Port string		`json:"port" envconfig:"PORT" default:"8000"`
	LogLevel string	`json:"log_level" envconfig:"LOG_LEVEL" default:"debug"`
}

// APIServer represents the HTTP server for the API.
type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *gin.Engine
	health func() string 
}