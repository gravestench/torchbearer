package models

import (
	"strings"
)

type FacilityTypeFlag int

const (
	FacilityAncestralVault FacilityTypeFlag = 1 << iota
	FacilityDocks
	FacilityDreamhouse
	FacilityFarmsOrPastures
	FacilityFlophouse
	FacilityGuildHall
	FacilityHedgeWitch
	FacilityHotel
	FacilityInn
	FacilityMarket
	FacilityHomes
	FacilityRiverOrCanal
	FacilityShrine
	FacilityStables
	FacilityStreets
	FacilityTavern
	FacilityTemple
	FacilityThievesGuild
	FacilityWall
	FacilityWell
	FacilityWizardsTower
)

func (ft FacilityTypeFlag) String() string {
	var s []string

	if ft&FacilityAncestralVault != 0 {
		s = append(s, "Ancestral Vault")
	}

	if ft&FacilityDocks != 0 {
		s = append(s, "Docks")
	}

	if ft&FacilityDreamhouse != 0 {
		s = append(s, "Dreamhouse")
	}

	if ft&FacilityFarmsOrPastures != 0 {
		s = append(s, "Farms / Pastures")
	}

	if ft&FacilityFlophouse != 0 {
		s = append(s, "Flophouse")
	}

	if ft&FacilityGuildHall != 0 {
		s = append(s, "Guild Hall")
	}

	if ft&FacilityHedgeWitch != 0 {
		s = append(s, "Hedge Witch")
	}

	if ft&FacilityHotel != 0 {
		s = append(s, "Hotel")
	}

	if ft&FacilityInn != 0 {
		s = append(s, "Inn")
	}

	if ft&FacilityMarket != 0 {
		s = append(s, "Market")
	}

	if ft&FacilityHomes != 0 {
		s = append(s, "Homes")
	}

	if ft&FacilityRiverOrCanal != 0 {
		s = append(s, "River / Canal")
	}

	if ft&FacilityShrine != 0 {
		s = append(s, "Shrine")
	}

	if ft&FacilityStables != 0 {
		s = append(s, "Stables")
	}

	if ft&FacilityStreets != 0 {
		s = append(s, "Streets")
	}

	if ft&FacilityTavern != 0 {
		s = append(s, "Tavern")
	}

	if ft&FacilityTemple != 0 {
		s = append(s, "Temple")
	}

	if ft&FacilityThievesGuild != 0 {
		s = append(s, "Thieves Guild")
	}

	if ft&FacilityWall != 0 {
		s = append(s, "Wall")
	}

	if ft&FacilityWell != 0 {
		s = append(s, "Well")
	}

	if ft&FacilityWizardsTower != 0 {
		s = append(s, "Wizard's Tower")
	}

	return strings.Join(s, ", ")
}

func (ft FacilityTypeFlag) Strings() []string {
	return strings.Split(ft.String(), ", ")
}

func (ft FacilityTypeFlag) List() []string {
	return strings.Split(ft.String(), ", ")
}
