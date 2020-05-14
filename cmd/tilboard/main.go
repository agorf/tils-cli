package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/agorf/tilboard-cli/adding"
	"github.com/agorf/tilboard-cli/editing"
	"github.com/agorf/tilboard-cli/listing"
	"github.com/agorf/tilboard-cli/removing"
	"github.com/agorf/tilboard-cli/showing"
	"github.com/agorf/tilboard-cli/store/http"
)

const defaultBaseURL = "https://tils.dev/api/"

func run() error {
	if len(os.Args) == 1 {
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

	store := http.NewStore(baseURL, apiToken)

	cmd, args := parseArgs()

	switch cmd {
	case "list":
		if len(args) != 0 {
			help()
		}
		if err := listing.Run(store); err != nil {
			handleError(err)
		}
	case "show":
		if len(args) != 1 {
			help()
		}
		if err := showing.Run(store, args[0]); err != nil {
			handleError(err)
		}
	case "new":
		if len(args) != 0 {
			help()
		}
		if err := adding.Run(store); err != nil {
			handleError(err)
		}
	case "edit":
		if len(args) != 1 {
			help()
		}
		if err := editing.Run(store, args[0]); err != nil {
			handleError(err)
		}
	case "delete":
		if len(args) != 1 {
			help()
		}
		if err := removing.Run(store, args[0]); err != nil {
			handleError(err)
		}
	default:
		help()
	}

	return nil
}

func parseArgs() (string, []string) {
	if len(os.Args) == 1 {
		return "", []string{}
	}
	return os.Args[1], os.Args[2:]
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
