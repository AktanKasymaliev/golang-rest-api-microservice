package server

import (
	"errors"
	"net/http"
)

func StartServer() error {
	server := NewInstanceServer()

	if server.health() != "OK" {
		server.logger.Fatal("Health check failed, maybe config or router is not set")
		return errors.New("Health check failed, maybe config or router is not set")
	}

	server.logger.Info("Starting Server")
	http_s := &http.Server{
		Addr:           server.config.Host + ":" +  server.config.Port,
		Handler:        server.router,
		// ...
	}
	RouteRegister(server.router)
	return http_s.ListenAndServe()
}