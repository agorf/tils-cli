package adding

import (
	"errors"
	"fmt"

	"github.com/agorf/tilboard-cli/editing"
)

type client interface {
	Post(string, interface{}, interface{}) error
}

func Run(c client) error {
	til := editing.Til{
		Title:    "Title",
		Content:  "Content",
		TagNames: []string{"tag1", "tag2", "tag3"},
	}

	text, err := editing.MarshalTil(&til)
	if err != nil {
		return err
	}

	newText, changed, err := editing.CaptureInputFromEditor(text)
	if err != nil {
		return err
	}
	if !changed {
		return errors.New("File did not change")
	}

	newTil, err := editing.UnmarshalTil(newText)
	if err != nil {
		return err
	}

	err = c.Post("/tils", map[string]editing.Til{"til": *newTil}, &til)
	if err != nil {
		return err
	}

	fmt.Println(Til(til))

	return nil
}
