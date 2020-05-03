package api

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

const (
	baseURL   = "https://tils.dev/api"
	userAgent = "https://github.com/agorf/tilboard-cli"
)

var (
	Token string // Injected
)

func NewRequest(method, path string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, baseURL+path, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+Token)
	req.Header.Set("User-Agent", userAgent)

	return req, nil
}

func Get(path string, target interface{}) error {
	req, err := NewRequest("GET", path, nil)
	if err != nil {
		return err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return errors.New(resp.Status)
	}

	err = json.NewDecoder(resp.Body).Decode(target)
	if err != nil {
		return err
	}

	return nil
}
