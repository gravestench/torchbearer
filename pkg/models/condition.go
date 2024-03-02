package models

import (
	"strings"
)

type Condition int

const (
	Fresh Condition = 1 << iota
	HungryAndThirsty
	Exhausted
	Angry
	Afraid
	Sick
	Injured
	Dead
)

func (c Condition) String() string {
	return strings.Join(c.Strings(), ", ")
}

func (c Condition) Strings() (conditions []string) {
	if (c & Fresh) > 0 {
		conditions = append(conditions, "Fresh")
	}

	if (c & HungryAndThirsty) > 0 {
		conditions = append(conditions, "HungryAndThirsty")
	}

	if (c & Exhausted) > 0 {
		conditions = append(conditions, "Exhausted")
	}

	if (c & Angry) > 0 {
		conditions = append(conditions, "Angry")
	}

	if (c & Afraid) > 0 {
		conditions = append(conditions, "Afraid")
	}

	if (c & Sick) > 0 {
		conditions = append(conditions, "Sick")
	}

	if (c & Injured) > 0 {
		conditions = append(conditions, "Injured")
	}

	if (c & Dead) > 0 {
		conditions = append(conditions, "Dead")
	}

	if len(conditions) < 1 {
		conditions = append(conditions, "None")
	}

	return
}
