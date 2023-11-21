package models

import (
	"fmt"
)

type Stock struct {
	Race
	Class
	Nature
	Description   string
	LevelBenefits [][]Benefit
}

func (s Stock) String() string {
	return fmt.Sprintf("%s %s", s.Race, s.Class)
}

type Benefit struct {
	Name        string
	Description string
	Selectable  bool
}
