package cmd

import (
	"errors"
	"fmt"

	"github.com/agorf/tilboard-cli/api"
	"github.com/agorf/tilboard-cli/editing"
	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:  "new",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		return new()
	},
}

func new() error {
	var til api.Til = api.Til{
		Title:    "Title",
		Content:  "Content",
		TagNames: []string{"tag1", "tag2", "tag3"},
	}

	text, err := editing.MarshalTil(&til)
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

	err = api.CreateTil(newTil)
	if err != nil {
		return err
	}

	fmt.Println(newTil.URL)

	return nil
}

func init() {
	rootCmd.AddCommand(newCmd)
}
