package main

import (
	"fmt"
	"os"

	"github.com/agorf/tilboard-cli/adding"
	"github.com/agorf/tilboard-cli/editing"
	"github.com/agorf/tilboard-cli/http"
	"github.com/agorf/tilboard-cli/listing"
	"github.com/agorf/tilboard-cli/removing"
	"github.com/agorf/tilboard-cli/showing"
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

	client := http.NewClient(baseURL, os.Getenv("TILBOARD_API_TOKEN"))

	cmd, args := parseArgs()

	switch cmd {
	case "list":
		if len(args) != 0 {
			help()
		}
		if err := listing.Run(client); err != nil {
			handleError(err)
		}
	case "show":
		if len(args) != 1 {
			help()
		}
		if err := showing.Run(client, args[0]); err != nil {
			handleError(err)
		}
	case "new":
		if len(args) != 0 {
			help()
		}
		if err := adding.Run(client); err != nil {
			handleError(err)
		}
	case "edit":
		if len(args) != 1 {
			help()
		}
		if err := editing.Run(client, args[0]); err != nil {
			handleError(err)
		}
	case "delete":
		if len(args) != 1 {
			help()
		}
		if err := removing.Run(client, args[0]); err != nil {
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
