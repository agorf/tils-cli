package editing

import (
	"bytes"
	"errors"
	"regexp"
	"strings"
	"text/template"

	"github.com/agorf/tilboard-cli/api"
)

const tilTemplate = `---
tags:{{with .TagNames}}{{range .}} {{.}}{{end}}{{end}}
visibility: {{.Visibility}}
archived: {{.Archived}}
---

# {{.Title}}

{{.Content}}
`

const tilTemplateRegex = `---\n` +
	`tags\s*:\s*(?P<tags>[ a-z0-9_]+)*\s*\n` +
	`visibility\s*:\s*(?P<visibility>public|unlisted|private)\s*\n` +
	`archived\s*:\s*(?P<archived>true|false)\s*\n` +
	`---\n` +
	`\n` +
	`#\s*(?P<title>.+)\n` +
	`\n` +
	`(?P<content>.*)`

var (
	re *regexp.Regexp
)

func MarshalTil(til *api.Til) ([]byte, error) {
	templ, err := template.New("til").Parse(tilTemplate)
	if err != nil {
		return nil, err
	}

	var data bytes.Buffer
	templ.Execute(&data, &til)

	return data.Bytes(), nil
}

func UnmarshalTil(data []byte) (*api.Til, error) {
	var til api.Til

	match := re.FindStringSubmatch(string(data))

	for i, name := range re.SubexpNames() {
		if i == 0 {
			continue
		}

		value := match[i]

		switch name {
		case "tags":
			til.TagNames = strings.Split(value, " ")
		case "visibility":
			v, err := api.VisibilityString(value)
			if err != nil {
				return nil, err
			}
			til.Visibility = v
		case "archived":
			til.Archived = value == "true"
		case "title":
			til.Title = value
		case "content":
			til.Content = value
		default:
			// TODO: Ask whether the user wants to re-edit the file
			return nil, errors.New("failed to parse file")
		}
	}

	return &til, nil
}

func init() {
	re = regexp.MustCompile(tilTemplateRegex)
}
