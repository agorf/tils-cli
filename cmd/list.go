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

	titleColor := Reset

	if t.Visibility != api.Public {
		titleColor = Red
	}

	return fmt.Sprintf(
		"%s  %s  %s  %s  %s",
		Colorize(t.UUID, Purple),
		Colorize(t.CreatedAt.Format("02 Jan 2006"), Blue),
		Colorize(t.Visibility.String()[0:3], Green),
		Colorize(t.Title, titleColor),
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
