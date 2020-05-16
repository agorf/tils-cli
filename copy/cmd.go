package copy

import (
	"fmt"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/AlecAivazis/survey/v2/terminal"
	"github.com/atotto/clipboard"
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
	tilIndex := 0
	tilPrompt := &survey.Select{
		Message: "Select til to copy:",
		Options: titles,
	}
	err = survey.AskOne(tilPrompt, &tilIndex)
	if err == terminal.InterruptErr {
		return nil
	}
	til := tils[tilIndex]

	allLines := strings.Split(strings.Replace(til.Content, "\r\n", "\n", -1), "\n")
	var lines []string
	for _, line := range allLines {
		if len(line) > 0 && line[0:3] != "```" {
			lines = append(lines, line)
		}
	}

	lineIndex := 0
	if len(lines) > 1 {
		linePrompt := &survey.Select{
			Message: "Select line to copy:",
			Options: lines,
		}
		err = survey.AskOne(linePrompt, &lineIndex)
		if err == terminal.InterruptErr {
			return nil
		}
	}
	line := lines[lineIndex]

	err = clipboard.WriteAll(line)
	if err != nil {
		return err
	}
	fmt.Println("Copied to clipboard")

	return nil
}
