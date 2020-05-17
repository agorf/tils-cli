package config

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/AlecAivazis/survey/v2/terminal"
)

func Run() error {
	token := ""
	tokenPrompt := &survey.Input{
		Message: "API token:",
	}
	err := survey.AskOne(
		tokenPrompt,
		&token,
		survey.WithValidator(survey.Required),
	)
	if err == terminal.InterruptErr {
		return nil
	}

	err = Write(token)
	if err != nil {
		return err
	}
	fmt.Println("Wrote config")

	return nil
}
