package models

type AdventurerRewards struct {
	Fate struct {
		Earned int
		Spent  int
	}
	Persona struct {
		Earned int
		Spent  int
	}
	Checks struct {
		Earned int
		Spent  int
	}
}
