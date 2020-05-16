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

func run() error {
	if len(os.Args) > 2 {
		help()
	}

	baseURL := os.Getenv("TILS_CLI_API_BASE_URL")
	if baseURL == "" {
		baseURL = defaultBaseURL
	}

	apiToken := os.Getenv("TILS_CLI_API_TOKEN")
	if apiToken == "" {
		handleError(errors.New("TILS_CLI_API_TOKEN environment variable is blank"))
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
	case "open":
		if err := open.Run(store); err != nil {
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
	case "archive":
		if err := archive.Run(store); err != nil {
			handleError(err)
		}
	case "delete":
		if err := delete.Run(store); err != nil {
			handleError(err)
		}
	case "version":
		fmt.Println(version.Version)
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
	fmt.Printf("%s [command]\n", os.Args[0])
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println()
	fmt.Println("    new      Create til")
	fmt.Println("    show     Show til content in the terminal")
	fmt.Println("    open     Open til in the browser")
	fmt.Println("    copy     Copy til to the clipboard")
	fmt.Println("    edit     Edit til")
	fmt.Println("    archive  Archive til")
	fmt.Println("    delete   Delete til")
	fmt.Println("    version  Print the current version")
	fmt.Println("    help     Print this help text")
	fmt.Println()
	fmt.Println("If a command is not provided, a picker will ask for one")
	fmt.Println()
	fmt.Println("Environment variables:")
	fmt.Println()
	fmt.Println("    TILS_CLI_API_TOKEN  Token to access the API with")

	os.Exit(1)
}

func main() {
	if err := run(); err != nil {
		handleError(err)
	}
}
