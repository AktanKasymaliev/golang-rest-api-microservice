package server

import (
	"github.com/sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

type Config struct {
	Host string
	Port string
	LogLevel string
}

type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *gin.Engine
	health func() string 
}