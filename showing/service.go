package showing

import (
	"fmt"
)

type store interface {
	GetTil(string, interface{}) error
}

func Run(s store, uuid string) error {
	var til Til

	err := s.GetTil(uuid, &til)
	if err != nil {
		return err
	}

	fmt.Println(til)

	return nil
}
