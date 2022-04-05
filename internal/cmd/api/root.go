package api

import (
	"log"

	"github.com/spf13/cobra"
)

func Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "api",
		Short: "Run the application server",
		Run: func(cmd *cobra.Command, _ []string) {
			applicationServerPort, err := cmd.PersistentFlags().GetInt(FlagApplicationServerPort)

			if err != nil {
				log.Fatal(err)
			}

			config := &Config{
				applicationServerPort: applicationServerPort,
			}
			server := NewServer(config)
			server.StartApplicationServer()
		}}

	persistentFlags := cmd.PersistentFlags()
	persistentFlags.Int(FlagApplicationServerPort, 8080, "Port to listen on")

	return cmd
}
