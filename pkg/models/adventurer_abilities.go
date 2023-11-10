package models

// all Modifiers are offsets to

type AdventurerAbilities struct {
	WillModifier   int
	HealthModifier int
	NatureModifier int
	TaxedNature    int
	MightModifier  int
	Town           struct {
		CirclesModifier    int
		PrecedenceModifier int
		ResourcesModifier  int
	}
}
