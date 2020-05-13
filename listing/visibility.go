package listing

import (
	"encoding/json"
	"fmt"
)

type Visibility byte

const (
	Public Visibility = iota
	Unlisted
	Private
)

func (v Visibility) String() string {
	return [...]string{"public", "unlisted", "private"}[v]
}

func VisibilityString(s string) (Visibility, error) {
	switch s {
	case "public":
		return Public, nil
	case "unlisted":
		return Unlisted, nil
	case "private":
		return Private, nil
	default:
		return Public, fmt.Errorf("invalid visibility string %v", s)
	}
}

func (v *Visibility) UnmarshalJSON(b []byte) error {
	var s string

	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	visibility, err := VisibilityString(s)
	if err != nil {
		return err
	}

	*v = visibility

	return nil
}
