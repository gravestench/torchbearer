package models

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
