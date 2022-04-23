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

			// Add retry policy in case the server is slow to start up.
			for i := 0; i <= 50; i++ {
				resp, err := http.Get("http://localhost:8080/ping")

				if err != nil || resp.StatusCode != 200 {
					continue
				}

				if resp.StatusCode == 200 {
					fmt.Print("Received a succesful response")
					return
				}

			}

			panic("Did not receive a successful response")
		}}

	return cmd
}
