package phase

import (
	"fmt"
	"strings"
)

type Phase struct {
	Name        string
	Description string
	SubPhase    []Phase `json:",omitempty"`
}

func (p Phase) String() string {
	var s string

	if p.Name != "" && p.Description != "" {
		s = fmt.Sprintf("%s: %s", p.Name, p.Description)
	} else if p.Description == "" {
		s = p.Name
	} else {
		s = p.Description
	}

	for _, subPhase := range p.SubPhase {
		for _, line := range strings.Split(subPhase.String(), "\n") {
			if len(subPhase.SubPhase) == 0 {
				s = fmt.Sprintf("%s\r\n  - %s", s, line)
			} else {
				s = fmt.Sprintf("%s\r\n  %s", s, line)
			}
		}
	}

	return s
}
