package models

import (
	"fmt"
)

type AdventurerSkill struct {
	RecordKey   string
	Level       int
	Advancement struct {
		Passes   int
		Failures int
	}
}

func (s AdventurerSkill) String() string {
	return fmt.Sprintf("%s %d", s.RecordKey, s.Level)
}
