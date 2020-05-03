package api

import (
	"encoding/json"
	"fmt"
	"time"
)

type Visibility byte

const (
	Public Visibility = iota
	Unlisted
	Private
)

type Til struct {
	UUID       string     `json:"uuid"`
	Title      string     `json:"title"`
	Content    string     `json:"content"`
	Visibility Visibility `json:"visibility"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	Archived   bool       `json:"archived"`
	TagNames   []string   `json:"tag_names"`
}

func (v Visibility) String() string {
	return [...]string{"public", "unlisted", "private"}[v]
}

func VisibilityString(s string) (Visibility, error) {
	switch s {
	case "public":
		return Public, nil
	case "unlisted":
		return Unlisted, nil
	case "private":
		return Private, nil
	default:
		return Public, fmt.Errorf("invalid visibility string %v", s)
	}
}

func (v *Visibility) UnmarshalJSON(b []byte) error {
	var s string

	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	visibility, err := VisibilityString(s)
	if err != nil {
		return err
	}

	*v = visibility

	return nil
}

func FetchTils() ([]Til, error) {
	var tils []Til

	err := Get("/tils", &tils)
	if err != nil {
		return nil, err
	}

	return tils, nil
}

func FetchTil(uuid string) (*Til, error) {
	var til Til

	err := Get(fmt.Sprintf("/tils/%s", uuid), &til)
	if err != nil {
		return nil, err
	}

	return &til, nil
}
