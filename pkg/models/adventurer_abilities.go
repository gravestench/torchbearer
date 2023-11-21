package models

// all Modifiers are offsets to

type AdventurerAbilities struct {
	Raw  RawAbilities
	Town TownAbilities
}

type RawAbilities struct {
	Will              int
	Health            int
	Nature            int
	NatureTax         int
	NatureDescriptors []string
	Might             int
}

type TownAbilities struct {
	Circles    int
	Precedence int
	Resources  int
}
