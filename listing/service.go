package listing

import (
	"fmt"
)

type store interface {
	GetTils(interface{}) error
}

func Run(s store) error {
	var tils []Til

	err := s.GetTils(&tils)
	if err != nil {
		return err
	}

	for _, til := range tils {
		fmt.Println(til)
	}

	return nil
}
