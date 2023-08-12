package server

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

// NewInstanceConfig Initializing
func newInstanceServerConfig() *Config {
	return &Config {
		Host: os.Getenv("HOST"),
		Port: os.Getenv("PORT"),
		LogLevel: os.Getenv("LOG_LEVEL"),
	}
}

// NewInstanceServer Initializing
func NewInstanceServer() *APIServer{
	return &APIServer{
			config: newInstanceServerConfig(),
			logger: logrus.New(),
			router: gin.Default(),
			health: HealthChecker,
	}
}

// HealthChecker checks health of the service
func HealthChecker() string {
	return "OK"
}