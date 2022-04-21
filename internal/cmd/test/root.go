package test

import (
	"card-match/internal/cmd/api"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

func Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "test",
		Short: "Run integration tests",
		Run: func(cmd *cobra.Command, _ []string) {
			server := api.NewDefaultServer()
			go server.StartApplicationServer()

			resp, err := http.Get("http://localhost:8080/ping")

			if err != nil {
				panic(err)
			}

			if resp.StatusCode != 200 {
				panic("Did not receive a successful response")
			}

			fmt.Print("Received a successful response")

		}}

	return cmd
}
