package cmd

import (
	"fmt"

	"github.com/agorf/tilboard-cli/api"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:  "list",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		return list()
	},
}

func list() error {
	tils, err := api.FetchTils()
	if err != nil {
		return err
	}

	for _, til := range tils {
		fmt.Println(til)
	}

	return nil
}

func init() {
	rootCmd.AddCommand(listCmd)
}
