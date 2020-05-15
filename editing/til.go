package editing

type Til struct {
	UUID       string     `json:"uuid"`
	Title      string     `json:"title"`
	Content    string     `json:"content"`
	Visibility Visibility `json:"visibility"`
	TagNames   []string   `json:"tag_names"`
}
