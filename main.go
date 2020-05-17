package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/AlecAivazis/survey/v2/terminal"
	"github.com/agorf/tils-cli/archive"
	"github.com/agorf/tils-cli/copy"
	"github.com/agorf/tils-cli/delete"
	"github.com/agorf/tils-cli/edit"
	"github.com/agorf/tils-cli/new"
	"github.com/agorf/tils-cli/open"
	"github.com/agorf/tils-cli/show"
	"github.com/agorf/tils-cli/store/http"
	"github.com/agorf/tils-cli/version"
)

const (
	defaultBaseURL = "https://tils.dev/api/"
)

func run() (bool, error) {
	if len(os.Args) > 2 {
		return true, errors.New("Invalid parameters")
	}

	baseURL := os.Getenv("TILS_CLI_API_BASE_URL")
	if baseURL == "" {
		baseURL = defaultBaseURL
	}

	apiToken := os.Getenv("TILS_CLI_API_TOKEN")
	if apiToken == "" {
		return false, errors.New("TILS_CLI_API_TOKEN environment variable is blank\n\nVisit https://tils.dev/account to get your API token")
	}

	command := ""
	if len(os.Args) == 1 {
		prompt := &survey.Select{
			Message: "Select til command:",
			Options: []string{
				"new",
				"show",
				"open",
				"copy",
				"edit",
				"archive",
				"delete",
				"version",
				"quit",
			},
		}
		err := survey.AskOne(prompt, &command)
		if err == terminal.InterruptErr {
			return false, nil
		}
	} else {
		command = os.Args[1]
	}

	store := http.NewStore(baseURL, apiToken)

	switch command {
	case "new":
		return false, new.Run(store)
	case "show":
		return false, show.Run(store)
	case "open":
		return false, open.Run(store)
	case "copy":
		return false, copy.Run(store)
	case "edit":
		return false, edit.Run(store)
	case "archive":
		return false, archive.Run(store)
	case "delete":
		return false, delete.Run(store)
	case "version":
		fmt.Println(version.Version)
	case "quit":
		// Do nothing
	default:
		return true, errors.New("Unrecognized command")
	}

	return false, nil
}

func help() {
	fmt.Printf("%s [command]\n", os.Args[0])
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println()
	fmt.Println("    new        Create til")
	fmt.Println("    show       Show til content in the terminal")
	fmt.Println("    open       Open til in the browser")
	fmt.Println("    copy       Copy til to the clipboard")
	fmt.Println("    edit       Edit til")
	fmt.Println("    archive    Archive til")
	fmt.Println("    delete     Delete til")
	fmt.Println("    version    Print the current version")
	fmt.Println("    help       Print this help text")
	fmt.Println()
	fmt.Println("If a command is not provided, a picker will ask for one")
}

func main() {
	if showHelp, err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)

		if showHelp {
			help()
		}

		os.Exit(1)
	}
}
