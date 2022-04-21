package main

import "card-match/internal/cmd/test"

func main() {
	cmd := test.Cmd()
	cmd.Execute()
}
