package models

type Location any
type Shrine any
type Temple any
type Law string

type Settlement struct {
	Name    string
	Type    SettlementType
	Culture struct {
		Government
		ShadowGovernment Government
		Traits           []TraitRecord
		Laws             []Law
		Shrines          []Shrine
		Temples          []Temple
	}
	Location
	Haggling struct {
		Obstruction
		Results string
	}
	TellingTales struct {
		Obstruction
		Attempts int
	}
	UnpayedDebts  []string
	NotableEvents []string
	NumDisasters  int
	EconomicLevel int
	Facilities    FacilityTypeFlag            // composite of all facilities present
	FacilityNotes map[FacilityTypeFlag]string // not a composite
	Personalities []Adventurer
}
