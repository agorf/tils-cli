package http

import (
	"fmt"
)

type store struct {
	client *Client
}

func NewStore(baseURL, apiToken string) *store {
	client := NewClient(baseURL, apiToken)
	return &store{&client}
}

func (s *store) GetTil(uuid string, target interface{}) error {
	err := s.client.Get(fmt.Sprintf("/tils/%s", uuid), target)
	if err != nil {
		return err
	}
	return nil
}

func (s *store) GetTils(target interface{}) error {
	err := s.client.Get("/tils", target)
	if err != nil {
		return err
	}
	return nil
}

func (s *store) AddTil(newTil interface{}, target interface{}) error {
	err := s.client.Post("/tils", map[string]interface{}{"til": newTil}, target)
	if err != nil {
		return err
	}
	return nil
}

func (s *store) RemoveTil(uuid string) error {
	err := s.client.Delete(fmt.Sprintf("/tils/%s", uuid))
	if err != nil {
		return err
	}
	return nil
}

func (s *store) UpdateTil(uuid string, newTil interface{}, target interface{}) error {
	err := s.client.Put(
		fmt.Sprintf("/tils/%s", uuid),
		map[string]interface{}{"til": newTil},
		target,
	)
	if err != nil {
		return err
	}
	return nil
}
