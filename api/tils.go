package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
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
	return [...]string{"Public", "Unlisted", "Private"}[v]
}

func (v *Visibility) UnmarshalJSON(b []byte) error {
	var s string

	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	switch s {
	case "public":
		*v = Public
	case "unlisted":
		*v = Unlisted
	case "private":
		*v = Private
	}

	return nil
}

func FetchTils() ([]Til, error) {
	req, err := NewRequest("GET", "/tils", nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}

	var tils []Til

	err = json.NewDecoder(resp.Body).Decode(&tils)
	if err != nil {
		return nil, err
	}

	return tils, nil
}

func FetchTil(uuid string) (*Til, error) {
	req, err := NewRequest("GET", fmt.Sprintf("/tils/%s", uuid), nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}

	var til Til

	err = json.NewDecoder(resp.Body).Decode(&til)
	if err != nil {
		return nil, err
	}

	return &til, nil
}
