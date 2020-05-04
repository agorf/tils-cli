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
	URL        string     `json:"url"`
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

func (v Visibility) MarshalJSON() ([]byte, error) {
	b, err := json.Marshal(v.String())
	if err != nil {
		return nil, err
	}
	return b, nil
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

func UpdateTil(uuid string, til *Til) error {
	err := Put(fmt.Sprintf("/tils/%s", uuid), til)
	if err != nil {
		return err
	}

	return nil
}

func DestroyTil(uuid string) error {
	err := Delete(fmt.Sprintf("/tils/%s", uuid))
	if err != nil {
		return err
	}

	return nil
}

func CreateTil(til *Til) error {
	err := Post("/tils", til)
	if err != nil {
		return err
	}

	return nil
}
