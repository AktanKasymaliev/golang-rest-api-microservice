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

	if server.health() != "OK" {
		server.logger.Fatal("Health check failed, maybe config or router is not set")
		return errors.New("Health check failed, maybe config or router is not set")
	}

	server.logger.Info("Starting Server")
	// Create an HTTP server instance.
	httpServer := &http.Server{
		Addr:           server.config.Host + ":" +  server.config.Port,
		Handler:        server.router,
		// ...
	}
	api.RouteRegister(server.router)
	return httpServer.ListenAndServe()
}