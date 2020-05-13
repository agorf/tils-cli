package removing

import (
	"fmt"
)

type client interface {
	Delete(string) error
}

func Run(c client, uuid string) error {
	err := c.Delete(fmt.Sprintf("/tils/%s", uuid))
	if err != nil {
		return err
	}

	return nil
}
