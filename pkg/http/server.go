// Package http implements a standard HTTP server.
package http

import (
	"fmt"
	"net/http"
)

type server struct {
	publicServer *http.Server
}

func ping() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("pong"))
	}
}

func NewHttpServer(port int) *server {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", ping())

	return &server{
		publicServer: &http.Server{
			Addr:    fmt.Sprintf(":%d", port),
			Handler: mux,
		},
	}
}
