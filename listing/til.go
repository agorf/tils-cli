package listing

import (
	"fmt"
	"strings"
	"time"

	"github.com/agorf/tilboard-cli/editing"
)

type Til struct {
	UUID       string             `json:"uuid"`
	Title      string             `json:"title"`
	Visibility editing.Visibility `json:"visibility"`
	CreatedAt  time.Time          `json:"created_at"`
	TagNames   []string           `json:"tag_names"`
}

func (t Til) String() string {
	prefixedTags := make([]string, len(t.TagNames))
	for i, tagName := range t.TagNames {
		prefixedTags[i] = "#" + tagName
	}

	visibility := ""
	if t.Visibility != editing.Public {
		visibility = colorize(fmt.Sprintf("(%s) ", t.Visibility), yellow)
	}

	return fmt.Sprintf(
		"%s  %s  %s%s  %s",
		colorize(t.UUID, purple),
		colorize(t.CreatedAt.Format("02 Jan 2006"), blue),
		visibility,
		t.Title,
		colorize(strings.Join(prefixedTags, " "), blue),
	)
}
