package cmd

import (
	"errors"

	"github.com/agorf/tilboard-cli/api"
	"github.com/agorf/tilboard-cli/editing"
	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:  "edit",
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return edit(args[0])
	},
}

func edit(uuid string) error {
	til, err := api.FetchTil(uuid)
	if err != nil {
		return err
	}

	text, err := editing.MarshalTil(til)
	if err != nil {
		return err
	}

	newText, changed, err := editing.CaptureInputFromEditor(text)
	if err != nil {
		return err
	}
	if !changed {
		return errors.New("File did not change")
	}

	newTil, err := editing.UnmarshalTil(newText)
	if err != nil {
		return err
	}

	err = api.UpdateTil(uuid, newTil)
	if err != nil {
		return err
	}

	return nil
}

func init() {
	rootCmd.AddCommand(editCmd)
}
