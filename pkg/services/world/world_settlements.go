package world

import (
	"slices"

	"github.com/google/uuid"

	"torchbearer/pkg/models"
)

// when initially creating a new world, we need to create settlements
func (w *World) generateNewWorldSettlements() {
	for _, t := range []models.SettlementType{
		models.SettlementBorderlandFortress,
		models.SettlementBustlingMetropolis,
		models.SettlementBusyCrossroads,
		models.SettlementDilapidatedPort,
		models.SettlementDwarvenHalls,
		models.SettlementElfhome,
		models.SettlementForgottenTemple,
		models.SettlementProsperousWayhouse,
		models.SettlementReligiousBastion,
		models.SettlementRemoteVillage,
		models.SettlementShire,
		models.SettlementSteading,
		models.SettlementWalledTown,
		models.SettlementWizardsTower,
	} {
		settlement := &models.Settlement{
			WorldID:      w.WorldID,
			SettlementID: uuid.New(),
		}

		settlement.Name = w.generateNewSettlementName()
		settlement.Seed = w.Seed + models.GenerateSeedFromString(settlement.Name)
		settlement.Type = t
		settlement.Facilities = w.generateNewSettlementFacilities(t)
		settlement.Culture.Traits = w.generateNewSettlementTraits(settlement.Type)
		settlement.Culture.Skills = w.generateNewSettlementSkills(settlement.Type)

		settlement.Culture.Government = models.RandomMundaneGovernment()
		settlement.Culture.ShadowGovernment = models.GovernmentUnknown

		if !settlement.Culture.Government.IsPossibleShadowGovernment() {
			if w.rng.Intn(2) == 0 {
				for !settlement.Culture.ShadowGovernment.IsPossibleShadowGovernment() {
					settlement.Culture.ShadowGovernment = models.RandomGovernment()
				}
			}
		}

		settlement.Culture.Laws = w.generateNewSettlementLaws()

		w.Settlements = append(w.Settlements, settlement)
	}
}

func (w *World) generateNewSettlementTraits(t models.SettlementType) (result []models.TraitRecord) {
	traits := make([]models.TraitRecord, 0)
	switch t {
	case models.SettlementElfhome:
		for _, trait := range w.records.Traits() {
			switch trait.Name {
			case "Calm", "Quiet":
				traits = append(traits, trait)
			}
		}
	case models.SettlementDwarvenHalls:
		for _, trait := range w.records.Traits() {
			switch trait.Name {
			case "Cunning", "Fiery":
				traits = append(traits, trait)
			}
		}
	case models.SettlementReligiousBastion:
		for _, trait := range w.records.Traits() {
			switch trait.Name {
			case "Defender", "Scarred":
				traits = append(traits, trait)
			}
		}
	case models.SettlementBustlingMetropolis:
		for _, trait := range w.records.Traits() {
			switch trait.Name {
			case "Extravagant", "Jaded":
				traits = append(traits, trait)
			}
		}
	case models.SettlementWizardsTower:
		for _, trait := range w.records.Traits() {
			switch trait.Name {
			case "Skeptical", "Thoughtful":
				traits = append(traits, trait)
			}
		}
	case models.SettlementRemoteVillage:
		for _, trait := range w.records.Traits() {
			switch trait.Name {
			case "Early Riser", "Rough Hands":
				traits = append(traits, trait)
			}
		}
	case models.SettlementBusyCrossroads:
		for _, trait := range w.records.Traits() {
			switch trait.Name {
			case "Foolhardy", "Quick-Witted":
				traits = append(traits, trait)
			}
		}
	default:
		for _, trait := range w.records.Traits() {
			traits = append(traits, trait)
		}

		for _, idx := range w.generateUniqueIndexes(3, len(traits)) {
			result = append(result, traits[idx])
		}
	}

	return result
}

func (w *World) generateNewSettlementSkills(t models.SettlementType) (result []models.SkillRecord) {
	skills := make([]models.SkillRecord, 0)

	switch t {
	case models.SettlementElfhome:
		for _, skill := range w.records.Skills() {
			switch skill.Name {
			case "Healer", "Mentor", "Pathfinder":
				skills = append(skills, skill)
			}
		}
	case models.SettlementDwarvenHalls:
		for _, skill := range w.records.Skills() {
			switch skill.Name {
			case "Armorer", "Laborer", "Stonemason":
				skills = append(skills, skill)
			}
		}
	case models.SettlementReligiousBastion:
		for _, skill := range w.records.Skills() {
			switch skill.Name {
			case "Cartographer", "Scholar", "Theologian":
				skills = append(skills, skill)
			}
		}
	case models.SettlementBustlingMetropolis:
		for _, skill := range w.records.Skills() {
			switch skill.Name {
			case "Haggler", "Sailor", "Steward":
				skills = append(skills, skill)
			}
		}
	case models.SettlementWizardsTower:
		for _, skill := range w.records.Skills() {
			switch skill.Name {
			case "Alchemist", "Lore Master", "Scholar":
				skills = append(skills, skill)
			}
		}
	case models.SettlementRemoteVillage:
		for _, skill := range w.records.Skills() {
			switch skill.Name {
			case "Carpenter", "Peasant", "Weaver":
				skills = append(skills, skill)
			}
		}
	case models.SettlementBusyCrossroads:
		for _, skill := range w.records.Skills() {
			switch skill.Name {
			case "Cook", "Haggler", "Rider":
				skills = append(skills, skill)
			}
		}
	default:
		for _, skill := range w.records.Skills() {
			skills = append(skills, skill)
		}

		for _, idx := range w.generateUniqueIndexes(3, len(skills)) {
			result = append(result, skills[idx])
		}
	}

	return result
}

func (w *World) generateNewSettlementLaws() (selected []models.Law) {
	laws := w.settlementLaws()

	for _, idx := range w.generateUniqueIndexes(3, len(laws)) {
		selected = append(selected, models.Law(laws[idx]))
	}

	return selected
}

func (w *World) generateUniqueIndexes(n, max int, selections ...int) []int {
	if n > max {
		return nil
	}

	if len(selections) >= n {
		return selections
	}

	index := w.rng.Intn(max)

	for slices.Contains(selections, index) {
		index = w.rng.Intn(max)
	}

	return w.generateUniqueIndexes(n, max, append(selections, index)...)
}

func (w *World) generateNewSettlementFacilities(t models.SettlementType) (f models.FacilityTypeFlag) {
	return t.RequiredFacilities() | t.RandomOptionalFacilities(w.rng)
}

func (s *Service) settlementLaws() []string {
	return []string{
		"No magic allowed within city limits.",
		"All citizens must pay a monthly tax to support the local guards.",
		"Weapons must be peace-bonded and registered with the city.",
		"Curfew is imposed after sunset; citizens must stay indoors.",
		"No trading with outsiders without approval from the council.",
		"Trespassing in the forest is forbidden; it's inhabited by dangerous creatures.",
		"Public intoxication is punishable by a fine or community service.",
		"Wearing a hood or mask in public is prohibited to maintain security.",
		"All disputes between citizens must be settled by the town's appointed mediator.",
		"Citizens must participate in a yearly festival to honor the town's patron deity.",
		"Magic users must register with the local mage's guild and undergo regular inspections.",
		"Taverns must close by midnight and not serve alcohol to minors.",
		"Hunting is only permitted in designated areas outside the city.",
		"Thievery is punishable by amputation of a hand or a heavy fine.",
		"Citizens are required to maintain their homes and keep them in good repair.",
		"Defamation of a noble's honor is a criminal offense.",
		"Foreigners must obtain a temporary permit to reside in the city.",
		"Beggars are not allowed within city walls; they must seek help at the temple.",
		"All magical artifacts must be reported to the authorities.",
		"Dueling is only allowed in designated areas and with the consent of both parties.",
		"Citizens must take part in a weekly community cleanup day.",
		"The city council reserves the right to requisition private property in times of need.",
		"In times of war, all able-bodied citizens must serve in the town militia.",
		"Littering is subject to a fine, and repeat offenders may face community service.",
		"Blacksmiths must be licensed and adhere to strict quality standards.",
		"Casting harmful spells on fellow citizens is a capital offense.",
		"All trade caravans must pay a toll to enter the city.",
		"Unsanctioned gatherings of more than 10 people are forbidden.",
		"Religious proselytizing is only allowed within designated areas.",
		"Fishermen must obtain a permit and follow seasonal fishing regulations.",
		"The possession of dangerous magical creatures is strictly prohibited.",
		"Punishment for theft may include branding and exile from the city.",
		"Citizens must maintain a minimum level of hygiene.",
		"All disputes involving debts are handled by a specialized debt court.",
		"Selling spoiled food is subject to heavy fines and potential closure of the establishment.",
		"Children must attend mandatory education provided by the city.",
		"The city's archives are restricted to authorized personnel only.",
		"Public defacement or vandalism is punishable by public flogging.",
		"Visitors must provide proof of vaccination before entering the city.",
		"Trading with known pirate vessels is a criminal offense.",
		"Citizens are encouraged to report any suspicious activities to the authorities.",
		"Horse-drawn carriages must yield to pedestrians and obey traffic regulations.",
		"All gambling establishments must be licensed and regulated.",
		"The use of necromancy is strictly forbidden and punishable by death.",
		"Citizens are responsible for maintaining their section of the city wall.",
		"All taverns and inns must keep a register of guests for security purposes.",
		"Public drunkenness is subject to arrest and a night in the stocks.",
		"Citizens are encouraged to participate in the annual harvest festival.",
		"Public nudity is prohibited within city limits.",
		"Street performers must obtain a permit and follow designated performance zones.",
		"The possession of cursed objects is subject to confiscation and disposal.",
		"Citizens must pay a fee for the disposal of waste and sewage.",
		"Unauthorized entry into the sewers is a criminal offense.",
		"All businesses must prominently display their license.",
		"A mandatory census of all citizens is conducted every five years.",
		"The sale of illegal drugs is subject to severe penalties.",
		"Citizens must assist the city guard in times of civil unrest.",
		"Unauthorized use of teleportation magic is strictly prohibited.",
		"Archery practice is only allowed in designated areas.",
		"Private duels to the death must be sanctioned by the city council.",
		"The possession of forbidden books is subject to confiscation and investigation.",
		"Citizens must adhere to a strict dress code for formal events.",
		"All animals within city limits must be properly registered and vaccinated.",
		"The city's water supply is strictly regulated; tampering is a criminal offense.",
		"Citizens are required to report any sightings of rare magical creatures.",
		"The use of necromantic magic for resurrection is forbidden.",
		"Citizens must obtain a permit for any construction or renovations.",
		"Vagrancy is punishable by a week of community service.",
		"Unlicensed fortune-tellers and seers are subject to arrest.",
		"Citizens must observe a moment of silence during the annual Remembrance Day.",
		"All goods transported by river must be inspected at the city docks.",
		"Citizens must participate in monthly fire drills.",
		"Unauthorized use of fire magic is a capital offense.",
		"All citizens must serve on a jury when called upon.",
		"The city's alchemical experiments are restricted to authorized personnel.",
		"Citizens are encouraged to report any sightings of undead creatures.",
		"The sale of counterfeit currency is subject to a lengthy prison sentence.",
		"All public celebrations must be approved by the city council.",
		"Private libraries must be registered with the city's library association.",
		"Citizens must provide food and lodging to visiting city officials.",
		"Unauthorized possession of magical artifacts is a criminal offense.",
		"Citizens must carry identification at all times.",
		"Bards and musicians must audition for a license to perform in taverns.",
		"Unauthorized use of illusion magic is subject to imprisonment.",
		"Citizens must participate in annual disaster preparedness drills.",
		"The city's sewers are patrolled regularly to prevent illegal activities.",
		"Public gatherings during the annual lunar eclipse are forbidden.",
		"Unauthorized use of transformation magic is strictly prohibited.",
		"Citizens must maintain their own section of the city's gardens.",
		"All magical creatures must be registered with the city's magical creature bureau.",
		"Citizens are required to participate in the annual citywide clean-up event.",
		"Unauthorized possession of dragon scales or body parts is a capital offense.",
		"Citizens must report any incidents of illegal potion brewing.",
	}
}
