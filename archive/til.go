package archive

type Til struct {
	UUID     string `json:"uuid"`
	Title    string `json:"title"`
	Archived bool   `json:"archived"`
}
