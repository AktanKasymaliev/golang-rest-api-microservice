package server

import (
	"base/internal/core"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// NewInstanceConfig Initializing
func newInstanceServerConfig() *Config {
	return &Config {
		host: core.Getenv("HOST", "localhost"),
		port: core.Getenv("PORT", "8080"),
		logLevel: core.Getenv("LOG_LEVEL", "debug"),
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