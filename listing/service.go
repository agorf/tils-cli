package listing

import (
	"fmt"
)

type client interface {
	Get(string, interface{}) error
}

func Run(c client) error {
	var tils []Til

	err := c.Get("/tils", &tils)
	if err != nil {
		return err
	}

	for _, til := range tils {
		fmt.Println(til)
	}

	return nil
}
