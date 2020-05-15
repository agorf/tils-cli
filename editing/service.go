package editing

import (
	"fmt"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/AlecAivazis/survey/v2/terminal"
)

type store interface {
	GetTils(interface{}) error
	UpdateTil(string, interface{}, interface{}) error
}

func Run(s store) error {
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

	title := ""
	titlePrompt := &survey.Input{
		Message: "Title:",
		Default: til.Title,
	}
	err = survey.AskOne(titlePrompt, &title)
	if err == terminal.InterruptErr {
		return nil
	}

	tagNames := ""
	tagNamesPrompt := &survey.Input{
		Message: "Tags:",
		Default: strings.Join(til.TagNames, " "),
	}
	err = survey.AskOne(tagNamesPrompt, &tagNames)
	if err == terminal.InterruptErr {
		return nil
	}

	visibilityStr := ""
	visibilityPrompt := &survey.Select{
		Message: "Visibility:",
		Options: []string{"public", "unlisted", "private"},
		Default: til.Visibility.String(),
	}
	err = survey.AskOne(visibilityPrompt, &visibilityStr)
	if err == terminal.InterruptErr {
		return nil
	}
	visibility, _ := VisibilityString(visibilityStr)

	content := ""
	contentPrompt := &survey.Editor{
		Message:       "Content:",
		Default:       til.Content,
		AppendDefault: true,
	}
	err = survey.AskOne(contentPrompt, &content)
	if err == terminal.InterruptErr {
		return nil
	}

	newTil := Til{
		Title:      title,
		Content:    string(content),
		TagNames:   strings.Split(tagNames, " "),
		Visibility: visibility,
	}

	err = s.UpdateTil(til.UUID, newTil, &til)
	if err != nil {
		return err
	}

	fmt.Println("Updated")

	return nil
}
