package api

import (
	"fmt"
	"net/http"
	"strconv"
)

func ping() http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("pong"))
	}
}

func getRandomBoard() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		// TODO: Actually make this random.
		randomizedBoard := 12231434
		w.Write([]byte(strconv.Itoa(randomizedBoard)))

	}

}

// An HTTP server with a single /ping endpoint.
func NewApplicationServer(port int) *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", ping())
	mux.HandleFunc("/getRandomBoard", getRandomBoard())

	applicationServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}

	applicationServer.RegisterOnShutdown(func() {
		fmt.Print("Shut down application server")
	})

	return applicationServer
}
