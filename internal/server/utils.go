package server

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

func NewInstanceConfig() *Config {
	return &Config {
		Host: os.Getenv("HOST"),
		Port: os.Getenv("PORT"),
		LogLevel: os.Getenv("LOG_LEVEL"),
	}
}

func NewInstanceServer() *APIServer{
	return &APIServer{
			config: NewInstanceConfig(),
			logger: logrus.New(),
			router: gin.Default(),
			health: HealthChecker,
	}
}

func HealthChecker() string {
	return "OK"
}