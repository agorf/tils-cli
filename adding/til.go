package adding

import (
	"github.com/agorf/tilboard-cli/editing"
)

type Til editing.Til

func (t Til) String() string {
	return t.URL
}
