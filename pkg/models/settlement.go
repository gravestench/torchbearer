package models

import (
	"github.com/google/uuid"
)

type Location any
type Shrine any
type Temple any
type Law string

type Settlement struct {
	UUID        uuid.UUID
	Seed        int64
	Name        string
	Type        SettlementType
	description string
	Culture     struct {
		Government
		ShadowGovernment Government
		Traits           []TraitRecord
		Skills           []Record
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
