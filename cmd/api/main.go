package main

import "card-match/internal/cmd/api"

func main() {
	cmd := api.Cmd()
	cmd.Execute()
}
