package models

import (
	"fmt"
)

type AdventurerTrait struct {
	RecordKey string
	Level     int
}

func (s AdventurerTrait) String() string {
	return fmt.Sprintf("%s %d", s.RecordKey, s.Level)
}
