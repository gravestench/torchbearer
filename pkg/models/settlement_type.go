package models

import (
	"math/rand"
)

type SettlementType string

const (
	SettlementBorderlandFortress SettlementType = "Borderland Fortress"
	SettlementBustlingMetropolis SettlementType = "Bustling Metropolis"
	SettlementBusyCrossroads     SettlementType = "Busy Crossroads"
	SettlementDilapidatedPort    SettlementType = "Dilapidated Port"
	SettlementDwarvenHalls       SettlementType = "Dwarven Halls"
	SettlementElfhome            SettlementType = "Elfhome"
	SettlementForgottenTemple    SettlementType = "Forgotten Temple"
	SettlementProsperousWayhouse SettlementType = "Prosperous Wayhouse"
	SettlementReligiousBastion   SettlementType = "Religious Bastion"
	SettlementRemoteVillage      SettlementType = "Remote Village"
	SettlementShire              SettlementType = "Shire"
	SettlementSteading           SettlementType = "Steading"
	SettlementWalledTown         SettlementType = "Walled Town"
	SettlementWizardsTower       SettlementType = "Wizard’s Tower"
)

func (st SettlementType) RequiredFacilities() FacilityTypeFlag {
	return map[SettlementType]FacilityTypeFlag{
		SettlementBorderlandFortress: 0 |
			FacilityStables |
			FacilityWall |
			FacilityWell,

		SettlementBustlingMetropolis: 0 |
			FacilityDocks |
			FacilityFlophouse |
			FacilityGuildHall |
			FacilityInn |
			FacilityMarket |
			FacilityHomes |
			FacilityRiverOrCanal |
			FacilityStables |
			FacilityStreets |
			FacilityTavern |
			FacilityTemple |
			FacilityThievesGuild |
			FacilityWell,

		SettlementBusyCrossroads: 0 |
			FacilityFlophouse |
			FacilityGuildHall |
			FacilityInn |
			FacilityMarket |
			FacilityHomes |
			FacilityStables |
			FacilityStreets |
			FacilityTavern |
			FacilityThievesGuild |
			FacilityWell,

		SettlementDilapidatedPort: 0 |
			FacilityDocks |
			FacilityFlophouse |
			FacilityGuildHall |
			FacilityMarket |
			FacilityHomes |
			FacilityRiverOrCanal |
			FacilityShrine |
			FacilityStables |
			FacilityStreets |
			FacilityTavern |
			FacilityThievesGuild,

		SettlementDwarvenHalls: 0 |
			FacilityAncestralVault |
			FacilityGuildHall |
			FacilityInn |
			FacilityMarket |
			FacilityHomes |
			FacilityRiverOrCanal |
			FacilityStables |
			FacilityStreets |
			FacilityTavern |
			FacilityWall |
			FacilityWell,

		SettlementElfhome: 0 |
			FacilityDreamhouse |
			FacilityHomes |
			FacilityShrine |
			FacilityTemple |
			FacilityWell,

		SettlementForgottenTemple: 0 |
			FacilityShrine |
			FacilityTemple |
			FacilityWell,

		SettlementProsperousWayhouse: 0 |
			FacilityInn |
			FacilityShrine |
			FacilityStables |
			FacilityWall |
			FacilityWell,

		SettlementReligiousBastion: 0 |
			FacilityHomes |
			FacilityShrine |
			FacilityStreets |
			FacilityTemple |
			FacilityWell,

		SettlementRemoteVillage: 0 |
			FacilityFarmsOrPastures |
			FacilityHomes |
			FacilityStables |
			FacilityStreets |
			FacilityWell,

		SettlementShire: 0 |
			FacilityFarmsOrPastures |
			FacilityHomes |
			FacilityStables |
			FacilityTavern |
			FacilityWell,

		SettlementSteading: 0 |
			FacilityFarmsOrPastures |
			FacilityHomes |
			FacilityStables |
			FacilityWell,

		SettlementWalledTown: 0 |
			FacilityDocks |
			FacilityFlophouse |
			FacilityGuildHall |
			FacilityMarket |
			FacilityHomes |
			FacilityRiverOrCanal |
			FacilityStables |
			FacilityStreets |
			FacilityTavern |
			FacilityThievesGuild |
			FacilityWall |
			FacilityWell,

		SettlementWizardsTower: 0 |
			FacilityWell |
			FacilityWizardsTower,
	}[st]
}

func (st SettlementType) OptionalFacilities() FacilityTypeFlag {
	return map[SettlementType]FacilityTypeFlag{
		SettlementBorderlandFortress: 0 |
			FacilityDocks |
			FacilityFlophouse |
			FacilityHedgeWitch |
			FacilityInn |
			FacilityHomes |
			FacilityRiverOrCanal |
			FacilityShrine |
			FacilityStreets |
			FacilityTavern |
			FacilityTemple |
			FacilityWizardsTower,

		SettlementBustlingMetropolis: 0 |
			FacilityFarmsOrPastures |
			FacilityHedgeWitch |
			FacilityHotel |
			FacilityShrine |
			FacilityWall |
			FacilityWizardsTower,

		SettlementBusyCrossroads: 0 |
			FacilityDocks |
			FacilityFarmsOrPastures |
			FacilityHedgeWitch |
			FacilityHotel |
			FacilityRiverOrCanal |
			FacilityShrine |
			FacilityWall,

		SettlementDilapidatedPort: 0 |
			FacilityFarmsOrPastures |
			FacilityHedgeWitch |
			FacilityHotel |
			FacilityInn |
			FacilityTemple |
			FacilityWall |
			FacilityWell,

		SettlementDwarvenHalls: 0 |
			FacilityDocks |
			FacilityFarmsOrPastures |
			FacilityFlophouse |
			FacilityHotel |
			FacilityShrine,

		SettlementElfhome: 0 |
			FacilityHotel |
			FacilityInn |
			FacilityMarket |
			FacilityTavern |
			FacilityRiverOrCanal,

		SettlementForgottenTemple: 0 |
			FacilityFlophouse |
			FacilityRiverOrCanal |
			FacilityStables |
			FacilityStreets |
			FacilityTavern |
			FacilityThievesGuild |
			FacilityWall,

		SettlementProsperousWayhouse: 0 |
			FacilityDocks |
			FacilityFlophouse |
			FacilityHedgeWitch |
			FacilityHotel |
			FacilityHomes |
			FacilityRiverOrCanal |
			FacilityTavern,

		SettlementReligiousBastion: 0 |
			FacilityDocks |
			FacilityFlophouse |
			FacilityGuildHall |
			FacilityHotel |
			FacilityInn |
			FacilityMarket |
			FacilityRiverOrCanal |
			FacilityStables |
			FacilityTavern |
			FacilityThievesGuild |
			FacilityWall,

		SettlementRemoteVillage: 0 |
			FacilityDocks |
			FacilityFlophouse |
			FacilityMarket |
			FacilityRiverOrCanal |
			FacilityShrine |
			FacilityTavern |
			FacilityThievesGuild,

		SettlementShire: 0 |
			FacilityInn |
			FacilityMarket |
			FacilityRiverOrCanal,

		SettlementSteading: 0 |
			FacilityRiverOrCanal |
			FacilityShrine |
			FacilityWall,

		SettlementWalledTown: 0 |
			FacilityFarmsOrPastures |
			FacilityHedgeWitch |
			FacilityHotel |
			FacilityInn |
			FacilityShrine |
			FacilityTemple,

		SettlementWizardsTower: 0 |
			FacilityHedgeWitch |
			FacilityInn |
			FacilityHomes |
			FacilityRiverOrCanal |
			FacilityShrine |
			FacilityStables |
			FacilityStreets |
			FacilityTavern,
	}[st]
}

func (st SettlementType) RandomOptionalFacilities() FacilityTypeFlag {
	flags := st.OptionalFacilities()
	// Initialize the result with all bits set to false (0)
	result := 0

	// Iterate through each bit of the flags
	for i := 0; i < 32; i++ { // Assuming a 32-bit integer
		bitMask := 1 << i

		// Check if the corresponding bit in the flags is set to true
		if int(flags)&bitMask != 0 {
			// If the flag bit is true, flip a coin to decide whether to set the bit in the result
			if rand.Intn(2) == 1 {
				result |= bitMask
			}
		}
	}

	return FacilityTypeFlag(result)
}
