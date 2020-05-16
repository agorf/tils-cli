package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/agorf/tils-cli/version"
)

const (
	apiVersion = 1
	userAgent  = "tils-cli/" + version.Version
)

type ResponseErrors struct {
	Errors []string `json:"errors"`
}

type Client struct {
	client  *http.Client
	baseURL string
	token   string
}

func NewClient(baseUrl, token string) Client {
	return Client{
		client:  &http.Client{},
		baseURL: baseUrl,
		token:   token,
	}
}

func (c Client) newRequest(method, path string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, c.baseURL+path, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", fmt.Sprintf("application/vnd.tilboard.v%d", apiVersion))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.token)
	req.Header.Set("User-Agent", userAgent)

	return req, nil
}

func (c Client) Get(path string, target interface{}) error {
	req, err := c.newRequest("GET", path, nil)
	if err != nil {
		return err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return errorFromResponse(resp)
	}

	err = json.NewDecoder(resp.Body).Decode(target)
	if err != nil {
		return err
	}

	return nil
}

func (c Client) Put(path string, data interface{}, target interface{}) error {
	var body bytes.Buffer

	err := json.NewEncoder(&body).Encode(data)
	if err != nil {
		return err
	}

	req, err := c.newRequest("PUT", path, &body)
	if err != nil {
		return err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return errorFromResponse(resp)
	}

	err = json.NewDecoder(resp.Body).Decode(target)
	if err != nil {
		return err
	}

	return nil
}

func (c Client) Delete(path string) error {
	req, err := c.newRequest("DELETE", path, nil)
	if err != nil {
		return err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return errorFromResponse(resp)
	}

	return nil
}

func (c Client) Post(path string, data interface{}, target interface{}) error {
	var body bytes.Buffer

	err := json.NewEncoder(&body).Encode(data)
	if err != nil {
		return err
	}

	req, err := c.newRequest("POST", path, &body)
	if err != nil {
		return err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return errorFromResponse(resp)
	}

	err = json.NewDecoder(resp.Body).Decode(target)
	if err != nil {
		return err
	}

	return nil
}

func errorFromResponse(resp *http.Response) error {
	if resp.StatusCode != 422 {
		return errors.New(resp.Status)
	}

	var respErrors ResponseErrors
	err := json.NewDecoder(resp.Body).Decode(&respErrors)
	if err != nil {
		return err
	}
	return errors.New(strings.Join(respErrors.Errors, ", "))
}
