package api

import (
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
	req, err := http.NewRequest("GET", baseURL+path, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Authorization", "Bearer "+Token)

	return req, nil
}
