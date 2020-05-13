package editing

type Til struct {
	UUID       string     `json:"uuid"`
	Title      string     `json:"title"`
	Content    string     `json:"content"`
	Visibility Visibility `json:"visibility"`
	Archived   bool       `json:"archived"`
	TagNames   []string   `json:"tag_names"`
	URL        string     `json:"url"` // Only used by adding
}
