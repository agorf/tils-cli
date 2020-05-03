package cmd

import (
	"github.com/agorf/tilboard-cli/api"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:  "delete",
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return delete(args[0])
	},
}

func delete(uuid string) error {
	err := api.DestroyTil(uuid)
	if err != nil {
		return err
	}
	return nil
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
