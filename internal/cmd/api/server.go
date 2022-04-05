package api

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type server struct {
	applicationServer *http.Server
}

func NewServer(config *Config) *server {
	applicationServer := NewApplicationServer(config.applicationServerPort)

	return &server{
		applicationServer: applicationServer,
	}

}

func (s *server) StartApplicationServer() error {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-quit
		fmt.Print("Gracefully shutting down application server\n")
		s.applicationServer.Shutdown(context.Background())
	}()

	return s.applicationServer.ListenAndServe()
}
