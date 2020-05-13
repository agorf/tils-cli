package showing

import (
	"fmt"
)

type client interface {
	Get(string, interface{}) error
}

func Run(c client, uuid string) error {
	var til Til

	err := c.Get(fmt.Sprintf("/tils/%s", uuid), &til)
	if err != nil {
		return err
	}

	fmt.Println(til)

	return nil
}
