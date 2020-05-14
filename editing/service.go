package editing

import (
	"errors"
)

type store interface {
	GetTil(string, interface{}) error
	UpdateTil(string, interface{}, interface{}) error
}

func Run(s store, uuid string) error {
	var til Til

	err := s.GetTil(uuid, &til)
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

	err = s.UpdateTil(uuid, newTil, &til)
	if err != nil {
		return err
	}

	return nil
}
