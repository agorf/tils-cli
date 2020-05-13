package editing

import (
	"errors"
	"fmt"
)

type client interface {
	Get(string, interface{}) error
	Put(string, interface{}, interface{}) error
}

func Run(c client, uuid string) error {
	var til Til

	err := c.Get(fmt.Sprintf("/tils/%s", uuid), &til)
	if err != nil {
		return err
	}

	text, err := MarshalTil(&til)
	if err != nil {
		return err
	}

	newText, changed, err := CaptureInputFromEditor(text)
	if err != nil {
		return err
	}
	if !changed {
		return errors.New("File did not change")
	}

	newTil, err := UnmarshalTil(newText)
	if err != nil {
		return err
	}

	err = c.Put(
		fmt.Sprintf("/tils/%s", uuid),
		map[string]Til{"til": *newTil},
		&til,
	)
	if err != nil {
		return err
	}

	return nil
}
