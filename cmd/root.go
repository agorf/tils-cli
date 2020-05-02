package cmd

import (
	"fmt"
	"os"

	"github.com/agorf/tilboard-cli/api"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:           "tilboard",
	SilenceErrors: true,
	SilenceUsage:  true,
}

func Execute() {
	api.Token = os.Getenv("TILBOARD_API_TOKEN")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
