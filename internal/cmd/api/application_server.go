package api

import (
	"fmt"
	"net/http"
)

func ping() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("pong"))
	}
}

// An HTTP server with a single /ping endpoint.
func NewApplicationServer(port int) *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", ping())

	applicationServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}

	applicationServer.RegisterOnShutdown(func() {
		fmt.Print("Shut down application server")
	})

	return applicationServer
}
