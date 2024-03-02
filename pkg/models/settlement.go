package models

import (
	"fmt"

	"github.com/google/uuid"
)

type Location any
type Shrine any
type Temple any
type Law string

type Settlement struct {
	WorldID      uuid.UUID
	SettlementID uuid.UUID
	Seed         int64
	Name         string
	Type         SettlementType
	description  string
	Culture      struct {
		Government
		ShadowGovernment Government
		Traits           []TraitRecord
		Skills           []SkillRecord
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
	Townsfolk     map[uuid.UUID]*Townsfolk
}

func (s *Settlement) InitTownsfolk() {
	if s.Townsfolk == nil {
		s.Townsfolk = make(map[uuid.UUID]*Townsfolk)
	}
}

func (s *Settlement) NewTownsfolk() *Townsfolk {
	if s.Townsfolk == nil {
		s.Townsfolk = make(map[uuid.UUID]*Townsfolk)
	}

	t := NewTownsfolk()
	t.WorldID = s.WorldID
	t.HometownID = s.SettlementID
	t.Relationships = make(map[uuid.UUID]string)

	s.Townsfolk[t.TownsfolkID] = t

	return t
}

func (s *Settlement) AddTownsfolk(t Townsfolk) {
	if s.Townsfolk == nil {
		s.Townsfolk = make(map[uuid.UUID]*Townsfolk)
	}

	s.Townsfolk[t.TownsfolkID] = &t
}

func (s *Settlement) GetTownsfolkByName(name string) (*Townsfolk, error) {
	for _, t := range s.Townsfolk {
		if t.Name == name {
			return t, nil
		}
	}

	return nil, fmt.Errorf("townsfolk with name %q not found", name)
}

func (s *Settlement) GetTownsfolkByID(id uuid.UUID) (*Townsfolk, error) {
	for _, t := range s.Townsfolk {
		if t.TownsfolkID == id {
			return t, nil
		}
	}

	return nil, fmt.Errorf("townsfolk with ID %q not found", id.String())
}
