package cmd

import (
	"fmt"
	"strings"

	"github.com/agorf/tilboard-cli/api"
	"github.com/spf13/cobra"
)

type ListTil api.Til

func (t ListTil) String() string {
	prefixedTags := make([]string, len(t.TagNames))

	for i, tagName := range t.TagNames {
		prefixedTags[i] = "#" + tagName
	}

	return fmt.Sprintf(
		"%s  %s  %s  %s  %s",
		t.UUID,
		t.CreatedAt.Format("02 Jan 2006"),
		t.Visibility.String()[0:3],
		t.Title,
		strings.Join(prefixedTags, " "),
	)
}

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
		fmt.Println(ListTil(til))
	}

	return nil
}

func init() {
	rootCmd.AddCommand(listCmd)
}