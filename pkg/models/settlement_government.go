package models

import (
	"math/rand"
)

type Government string

const (
	GovernmentUnknown             Government = "Unknown"
	GovernmentMonarchy            Government = "Monarchy"
	GovernmentOligarchy           Government = "Oligarchy"
	GovernmentDemocracy           Government = "Democracy"
	GovernmentTheocracy           Government = "Theocracy"
	GovernmentTribalCouncil       Government = "Tribal Council"
	GovernmentFeudalism           Government = "Feudalism"
	GovernmentRepublic            Government = "Republic"
	GovernmentAutocracy           Government = "Autocracy"
	GovernmentAnarchy             Government = "Anarchy"
	GovernmentConfederation       Government = "Confederation"
	GovernmentMagocracy           Government = "Magocracy"
	GovernmentMerchantRepublic    Government = "Merchant Republic"
	GovernmentPlutocracy          Government = "Plutocracy"
	GovernmentMilitaryJunta       Government = "Military Junta"
	GovernmentGuildocracy         Government = "Guildocracy"
	GovernmentAIGovernance        Government = "AI Governance"
	GovernmentVampireCourt        Government = "Vampire Court"
	GovernmentCouncilOfElders     Government = "Council of Elders"
	GovernmentHiveMind            Government = "Hive Mind"
	GovernmentNecrocracy          Government = "Necrocracy"
	GovernmentChiefdom            Government = "Viking Clan Chiefdom"
	GovernmentRunestoneCouncil    Government = "Runestone Council"
	GovernmentFjordlordship       Government = "Fjordlordship"
	GovernmentValkyrieSisterhood  Government = "Valkyrie Sisterhood"
	GovernmentCouncilOfMages      Government = "Council of Mages"
	GovernmentUnderworldSyndicate Government = "Underworld Syndicate"
	GovernmentNomadicClans        Government = "Nomadic Clans"
	GovernmentCult                Government = "Cult"
	GovernmentBureaucracy         Government = "Bureaucracy"
)

// Description returns the description of the government.
func (g Government) Description() string {
	description, found := map[Government]string{
		GovernmentUnknown:             "The type of government is unknown.",
		GovernmentMonarchy:            "A single ruler, often a king or queen, holds absolute power and authority over the settlement.",
		GovernmentOligarchy:           "A small, privileged group of individuals, such as noble families or council members, governs the settlement.",
		GovernmentDemocracy:           "The settlement's citizens have the power to elect representatives who make decisions on their behalf.",
		GovernmentTheocracy:           "Religious leaders or a divine entity guide and rule the settlement, and religious laws are paramount.",
		GovernmentTribalCouncil:       "The settlement is governed by a council of tribal elders or chieftains who make decisions for the community.",
		GovernmentFeudalism:           "Landowners (lords) grant land and protection to vassals in exchange for loyalty and military service.",
		GovernmentRepublic:            "Elected officials, often from different regions or districts, make decisions for the settlement's citizens.",
		GovernmentAutocracy:           "A single ruler with absolute power controls all aspects of government and society.",
		GovernmentAnarchy:             "There is no formal government, and the settlement operates without central authority, relying on local customs or traditions.",
		GovernmentConfederation:       "Several smaller settlements form an alliance, with each retaining a degree of autonomy while cooperating on common issues.",
		GovernmentMagocracy:           "Wizards, sorcerers, or magic-users hold political power and make decisions based on their magical abilities.",
		GovernmentMerchantRepublic:    "Wealthy merchants and trade guilds control the government and economy of the settlement.",
		GovernmentPlutocracy:          "Political power is concentrated in the hands of the wealthiest citizens, often through financial influence.",
		GovernmentMilitaryJunta:       "A group of military officers seizes control of the government, making decisions through force.",
		GovernmentGuildocracy:         "Various guilds, such as craftsmen or adventurers, govern the settlement, each representing their specific interests.",
		GovernmentAIGovernance:        "An advanced artificial intelligence system or magical construct governs the settlement, making decisions based on algorithms or programming.",
		GovernmentVampireCourt:        "A group of powerful vampires or other supernatural beings rule over the settlement's mortal inhabitants.",
		GovernmentCouncilOfElders:     "A council composed of the oldest and wisest members of the settlement makes important decisions.",
		GovernmentHiveMind:            "The settlement's inhabitants are interconnected in a psychic or telepathic network, making collective decisions.",
		GovernmentNecrocracy:          "Ruled by undead or necromancers, the government may involve dark and sinister forces.",
		GovernmentChiefdom:            "A powerful Viking clan, led by a chieftain, governs the settlement, often through martial strength and seafaring prowess.",
		GovernmentRunestoneCouncil:    "A council of wise elders makes decisions for the community, with guidance sought from ancient runestones and Norse gods.",
		GovernmentFjordlordship:       "The settlement is ruled by a fjordlord, a seafaring ruler who controls access to vital coastal resources.",
		GovernmentValkyrieSisterhood:  "An order of valkyries, warrior maidens of Norse mythology, governs the settlement and ensures bravery in battle.",
		GovernmentCouncilOfMages:      "A council of powerful mages and sorcerers governs the settlement, wielding immense magical influence.",
		GovernmentUnderworldSyndicate: "Criminal organizations control the government, with a web of illegal activities and secret operations.",
		GovernmentNomadicClans:        "Groups of nomadic tribes or clans govern through cooperation and survival in harsh wilderness.",
		GovernmentCult:                "A secretive and fanatical cult rules the settlement, often worshiping dark or mysterious entities.",
		GovernmentBureaucracy:         "A complex and bureaucratic system with various offices and officials governs the settlement's affairs.",
	}[g]

	if !found {
		return "Unknown"
	}

	return description
}

func (g Government) IsPossibleShadowGovernment() bool {
	switch g {
	case
		GovernmentMagocracy,
		GovernmentPlutocracy,
		GovernmentMilitaryJunta,
		GovernmentGuildocracy,
		GovernmentAIGovernance,
		GovernmentVampireCourt,
		GovernmentCouncilOfElders,
		GovernmentHiveMind,
		GovernmentNecrocracy,
		GovernmentRunestoneCouncil,
		GovernmentCouncilOfMages,
		GovernmentUnderworldSyndicate,
		GovernmentCult:
		return true
	}

	return false
}

func RandomMundaneGovernment() Government {
	g := []Government{
		GovernmentMonarchy,
		GovernmentOligarchy,
		GovernmentDemocracy,
		GovernmentTheocracy,
		GovernmentTribalCouncil,
		GovernmentFeudalism,
		GovernmentRepublic,
		GovernmentAutocracy,
		GovernmentAnarchy,
		GovernmentConfederation,
		GovernmentMerchantRepublic,
		GovernmentPlutocracy,
		GovernmentMilitaryJunta,
		GovernmentGuildocracy,
		GovernmentCouncilOfElders,
		GovernmentChiefdom,
		GovernmentRunestoneCouncil,
		GovernmentFjordlordship,
		GovernmentNomadicClans,
		GovernmentBureaucracy,
	}

	return g[rand.Intn(len(g))]
}

func RandomGovernment() Government {
	g := []Government{
		GovernmentUnknown,
		GovernmentMonarchy,
		GovernmentOligarchy,
		GovernmentDemocracy,
		GovernmentTheocracy,
		GovernmentTribalCouncil,
		GovernmentFeudalism,
		GovernmentRepublic,
		GovernmentAutocracy,
		GovernmentAnarchy,
		GovernmentConfederation,
		GovernmentMagocracy,
		GovernmentMerchantRepublic,
		GovernmentPlutocracy,
		GovernmentMilitaryJunta,
		GovernmentGuildocracy,
		GovernmentAIGovernance,
		GovernmentVampireCourt,
		GovernmentCouncilOfElders,
		GovernmentHiveMind,
		GovernmentNecrocracy,
		GovernmentChiefdom,
		GovernmentRunestoneCouncil,
		GovernmentFjordlordship,
		GovernmentValkyrieSisterhood,
		GovernmentCouncilOfMages,
		GovernmentUnderworldSyndicate,
		GovernmentNomadicClans,
		GovernmentCult,
		GovernmentBureaucracy,
	}

	return g[rand.Intn(len(g))]
}
