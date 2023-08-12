package server

import (
	"base/internal/server/api"
	"errors"
	"net/http"
)

// StartServer starts the HTTP server.
func StartServer() error {
	// Create a new instance of the APIServer.
	server := NewInstanceServer()

	if err := server.startDB(); err != nil {
		server.logger.Fatal("DB connection failure.")
	}

	if server.health() != "OK" {
		server.logger.Fatal("Health check failed, maybe config or router is not set")
		return errors.New("Health check failed, maybe config or router is not set")
	}

	server.logger.Info("Starting Server")
	// Create an HTTP server instance.
	httpServer := &http.Server{
		Addr:           server.config.host + ":" +  server.config.port,
		Handler:        server.router,
		// ...
	}
	api.RouteRegister(server.router)
	return httpServer.ListenAndServe()
}
