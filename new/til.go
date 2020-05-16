package new

import (
	"github.com/agorf/tils-cli/edit"
)

type Til struct {
	Title      string          `json:"title"`
	Content    string          `json:"content"`
	Visibility edit.Visibility `json:"visibility"`
	TagNames   []string        `json:"tag_names"`
	URL        string          `json:"url"`
}
