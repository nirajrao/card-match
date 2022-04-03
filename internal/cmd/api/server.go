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

func NewServer() *server {
	applicationServer := NewApplicationServer(8080)

	return &server{
		applicationServer: applicationServer,
	}

}

func (s *server) StartApplicationServer() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-quit
		fmt.Print("Gracefully shutting down application server\n")
		s.applicationServer.Shutdown(context.Background())
	}()

	s.applicationServer.ListenAndServe()
}
