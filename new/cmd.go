package new

import (
	"fmt"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/AlecAivazis/survey/v2/terminal"
	"github.com/agorf/tilboard-cli/edit"
)

type store interface {
	AddTil(interface{}, interface{}) error
}

func Run(s store) error {
	title := ""
	titlePrompt := &survey.Input{
		Message: "Title:",
	}
	err := survey.AskOne(
		titlePrompt,
		&title,
		survey.WithValidator(survey.Required),
	)
	if err == terminal.InterruptErr {
		return nil
	}

	tagNames := ""
	tagNamesPrompt := &survey.Input{
		Message: "Tags:",
	}
	err = survey.AskOne(tagNamesPrompt, &tagNames)
	if err == terminal.InterruptErr {
		return nil
	}

	visibilityStr := ""
	visibilityPrompt := &survey.Select{
		Message: "Visibility:",
		Options: []string{"public", "unlisted", "private"},
	}
	err = survey.AskOne(visibilityPrompt, &visibilityStr)
	if err == terminal.InterruptErr {
		return nil
	}
	visibility, _ := edit.VisibilityString(visibilityStr)

	content := ""
	contentPrompt := &survey.Editor{
		Message:       "Content:",
		AppendDefault: true,
	}
	err = survey.AskOne(
		contentPrompt,
		&content,
		survey.WithValidator(survey.Required),
	)
	if err == terminal.InterruptErr {
		return nil
	}

	create := false
	createPrompt := &survey.Confirm{
		Message: "Create?",
	}
	err = survey.AskOne(createPrompt, &create)
	if err == terminal.InterruptErr {
		return nil
	}
	if !create {
		fmt.Println("Aborted")
		return nil
	}

	newTil := Til{
		Title:      title,
		Content:    string(content),
		TagNames:   strings.Split(tagNames, " "),
		Visibility: visibility,
	}

	var til Til
	err = s.AddTil(newTil, &til)
	if err != nil {
		return err
	}

	fmt.Println(til.URL)

	return nil
}
