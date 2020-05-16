package show

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/AlecAivazis/survey/v2/terminal"
)

type store interface {
	GetTils(interface{}) error
}

func Run(s store) error {
	fmt.Println("Fetching tils...")
	var tils []Til
	err := s.GetTils(&tils)
	if err != nil {
		return err
	}

	var titles []string
	for _, til := range tils {
		titles = append(titles, til.Title)
	}
	index := 0
	tilPrompt := &survey.Select{
		Message: "Select til to edit:",
		Options: titles,
	}
	err = survey.AskOne(tilPrompt, &index)
	if err == terminal.InterruptErr {
		return nil
	}
	til := tils[index]

	fmt.Println(til.Content)

	return nil
}
