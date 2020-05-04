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

	visibility := ""

	if t.Visibility != api.Public {
		visibility = Colorize(fmt.Sprintf("(%s) ", t.Visibility.String()), Yellow)
	}

	return fmt.Sprintf(
		"%s  %s  %s%s  %s",
		Colorize(t.UUID, Purple),
		Colorize(t.CreatedAt.Format("02 Jan 2006"), Blue),
		visibility,
		t.Title,
		Colorize(strings.Join(prefixedTags, " "), Blue),
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
