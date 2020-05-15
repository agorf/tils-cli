package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/AlecAivazis/survey/v2/terminal"
	"github.com/agorf/tilboard-cli/copy"
	"github.com/agorf/tilboard-cli/delete"
	"github.com/agorf/tilboard-cli/edit"
	"github.com/agorf/tilboard-cli/new"
	"github.com/agorf/tilboard-cli/show"
	"github.com/agorf/tilboard-cli/store/http"
)

const (
	defaultBaseURL = "https://tils.dev/api/"
	version        = "0.1.0"
)

func run() error {
	if len(os.Args) > 2 {
		help()
	}

	baseURL := os.Getenv("TILBOARD_API_BASE_URL")
	if baseURL == "" {
		baseURL = defaultBaseURL
	}

	apiToken := os.Getenv("TILBOARD_API_TOKEN")
	if apiToken == "" {
		handleError(errors.New("TILBOARD_API_TOKEN environment variable is blank"))
	}

	command := ""
	if len(os.Args) == 1 {
		prompt := &survey.Select{
			Message: "Command:",
			Options: []string{"new", "show", "copy", "edit", "delete", "version", "quit"},
		}
		err := survey.AskOne(prompt, &command)
		if err == terminal.InterruptErr {
			os.Exit(0)
		}
	} else {
		command = os.Args[1]
	}

	store := http.NewStore(baseURL, apiToken)

	switch command {
	case "new":
		if err := new.Run(store); err != nil {
			handleError(err)
		}
	case "show":
		if err := show.Run(store); err != nil {
			handleError(err)
		}
	case "copy":
		if err := copy.Run(store); err != nil {
			handleError(err)
		}
	case "edit":
		if err := edit.Run(store); err != nil {
			handleError(err)
		}
	case "delete":
		if err := delete.Run(store); err != nil {
			handleError(err)
		}
	case "version":
		fmt.Println(version)
	case "quit":
		os.Exit(0)
	default:
		help()
	}

	return nil
}

func handleError(err error) {
	fmt.Fprintf(os.Stderr, "%v\n", err)
	os.Exit(1)
}

func help() {
	fmt.Println("help")
	os.Exit(1)
}

func main() {
	if err := run(); err != nil {
		handleError(err)
	}
}
