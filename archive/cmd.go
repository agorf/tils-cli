package archive

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/AlecAivazis/survey/v2/terminal"
)

type store interface {
	GetTils(interface{}) error
	UpdateTil(string, interface{}, interface{}) error
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
		Message: "Select til to archive:",
		Options: titles,
	}
	err = survey.AskOne(tilPrompt, &index)
	if err == terminal.InterruptErr {
		return nil
	}
	til := tils[index]

	archive := false
	archivePrompt := &survey.Confirm{
		Message: "Archive?",
	}
	err = survey.AskOne(archivePrompt, &archive)
	if err == terminal.InterruptErr {
		return nil
	}
	if !archive {
		fmt.Println("Aborted")
		return nil
	}
	til.Archived = true

	fmt.Print("Archiving... ")
	err = s.UpdateTil(til.UUID, til, &til)
	if err != nil {
		return err
	}
	fmt.Println("Done")

	return nil
}
