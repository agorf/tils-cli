package cmd

import (
	"fmt"
	"os"

	"github.com/agorf/tilboard-cli/api"
	"github.com/spf13/cobra"
)

const defaultBaseURL = "https://tils.dev/api/"

var rootCmd = &cobra.Command{
	Use:           "tilboard",
	SilenceErrors: true,
	SilenceUsage:  true,
}

func Execute() {
	api.BaseURL = os.Getenv("TILBOARD_API_BASE_URL")

	if api.BaseURL == "" {
		api.BaseURL = defaultBaseURL
	}

	api.Token = os.Getenv("TILBOARD_API_TOKEN")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
