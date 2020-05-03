package cmd

import (
	"fmt"
	"strings"

	"github.com/agorf/tilboard-cli/api"
	"github.com/spf13/cobra"
)

type ShowTil api.Til

func (t ShowTil) String() string {
	prefixedTags := make([]string, len(t.TagNames))

	for i, tagName := range t.TagNames {
		prefixedTags[i] = "#" + tagName
	}

	return fmt.Sprintf(
		"%s\n%s\n%s\n\n%s\n\n%s  %s  %s",
		strings.Repeat("=", len(t.Title)),
		t.Title,
		strings.Repeat("=", len(t.Title)),
		t.Content,
		t.CreatedAt.Format("Mon, 02 Jan 2006"),
		t.Visibility,
		strings.Join(prefixedTags, " "),
	)
}

var showCmd = &cobra.Command{
	Use:  "show",
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return show(args[0])
	},
}

func show(uuid string) error {
	til, err := api.FetchTil(uuid)
	if err != nil {
		return err
	}

	fmt.Println(ShowTil(*til))

	return nil
}

func init() {
	rootCmd.AddCommand(showCmd)
}
