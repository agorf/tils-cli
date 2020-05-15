package adding

import (
	"github.com/agorf/tilboard-cli/editing"
)

type Til struct {
	Title      string             `json:"title"`
	Content    string             `json:"content"`
	Visibility editing.Visibility `json:"visibility"`
	TagNames   []string           `json:"tag_names"`
	URL        string             `json:"url"`
}
