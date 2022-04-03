package main

import (
	"card-match/internal/cmd/api"
)

func main() {
	server := api.NewServer()
	server.StartApplicationServer()
}
