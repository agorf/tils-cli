package showing

import (
	"fmt"
	"strings"
	"time"

	"github.com/agorf/tilboard-cli/editing"
)

type Til struct {
	Title      string             `json:"title"`
	Content    string             `json:"content"`
	Visibility editing.Visibility `json:"visibility"`
	CreatedAt  time.Time          `json:"created_at"`
	Archived   bool               `json:"archived"`
	TagNames   []string           `json:"tag_names"`
}

func (t Til) String() string {
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
