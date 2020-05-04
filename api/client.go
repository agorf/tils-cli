package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

const (
	userAgent = "https://github.com/agorf/tilboard-cli"
)

var (
	BaseURL string // Injected
	Token   string // Injected
)

func newRequest(method, path string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, BaseURL+path, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+Token)
	req.Header.Set("User-Agent", userAgent)

	return req, nil
}

func Get(path string, target interface{}) error {
	req, err := newRequest("GET", path, nil)
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

func Put(path string, til *Til) error {
	var body bytes.Buffer

	err := json.NewEncoder(&body).Encode(map[string]Til{"til": *til})
	if err != nil {
		return err
	}

	req, err := newRequest("PUT", path, &body)
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

	err = json.NewDecoder(resp.Body).Decode(til)
	if err != nil {
		return err
	}

	return nil
}

func Delete(path string) error {
	req, err := newRequest("DELETE", path, nil)
	if err != nil {
		return err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return errors.New(resp.Status)
	}

	return nil
}

func Post(path string, til *Til) error {
	var body bytes.Buffer

	err := json.NewEncoder(&body).Encode(map[string]Til{"til": *til})
	if err != nil {
		return err
	}

	req, err := newRequest("POST", path, &body)
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

	err = json.NewDecoder(resp.Body).Decode(til)
	if err != nil {
		return err
	}

	return nil
}
