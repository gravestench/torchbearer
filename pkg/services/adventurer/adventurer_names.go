package adventurer

import (
	"fmt"
	"math/rand"
)

func (s *Service) GenerateFirstName() string {
	return firstName[rand.Intn(len(firstName))]
}

func (s *Service) GenerateLastName() string {
	return lastName[rand.Intn(len(lastName))]
}

func (s *Service) GenerateAdventurerName(withTitle, withEpithet bool) string {
	first := firstName[rand.Intn(len(firstName))]
	last := lastName[rand.Intn(len(lastName))]

	name := fmt.Sprintf("%s %s", first, last)

	if withTitle {
		title := titles[rand.Intn(len(titles))]
		name = fmt.Sprintf("%s %s", title, name)
	}

	if withEpithet {
		epithet := epithets[rand.Intn(len(epithets))]
		name = fmt.Sprintf("%s %s", name, epithet)
	}

	return name
}

func (s *Service) GenerateRandomName() string {
	withTitle := rand.Intn(3)   // 33% chance of having a title
	withEpithet := rand.Intn(3) // 33% chance of having an epithet

	if withTitle == 0 {
		title := titles[rand.Intn(len(titles))]
		first := firstName[rand.Intn(len(firstName))]
		last := lastName[rand.Intn(len(lastName))]
		return fmt.Sprintf("%s %s %s", title, first, last)
	} else if withEpithet == 0 {
		epithet := epithets[rand.Intn(len(epithets))]
		first := firstName[rand.Intn(len(firstName))]
		last := lastName[rand.Intn(len(lastName))]
		return fmt.Sprintf("%s %s %s", first, last, epithet)
	} else {
		first := firstName[rand.Intn(len(firstName))]
		last := lastName[rand.Intn(len(lastName))]
		return fmt.Sprintf("%s %s", first, last)
	}
}

var titles = []string{
	"Sir", "Lady", "Lord", "Dame", "Baron", "Baroness", "Count", "Countess", "Duke", "Duchess",
	"Archmage", "High Priest", "Grandmaster", "Knight Commander", "Captain", "Master Thief", "Guildmaster",
	"Ranger Captain", "Bard", "Sorcerer", "Warrior", "Assassin", "Mercenary", "Enchanter", "Sellsword",
	"Dragon Slayer", "Alchemist", "Explorer", "Admiral", "High Inquisitor", "Shadowblade", "Spellweaver",
	"Druid", "Warden", "Champion", "Nobleman", "Noblewoman",
}

var epithets = []string{
	"the Brave", "the Wise", "the Fearless", "the Cunning", "the Valiant", "the Kind", "the Noble", "the Mysterious",
	"the Enigmatic", "the Heroic", "the Reckless", "the Fierce", "the Just", "the Mighty", "the Swift", "the Honorable",
	"the Silent", "the Shadow", "the Resolute", "the Cursed", "the Lucky", "the Merciful", "the Unyielding",
	"the Charmed", "the Relentless", "the Radiant", "the Dark", "the Eminent", "the Ironclad", "the Vengeful",
	"the Vigilant", "the Serene", "the Merciless", "the Indomitable", "the Stalwart", "the Fearbringer", "the Serpent's Bane",
	"the Stormcaller", "the Whisperer", "the Moonwalker", "the Starbringer", "the Voidbringer", "the Soulseeker",
	"the Lorekeeper", "the Windrider", "the Firestarter", "the Earthshaker", "the Tidecaller", "the Frostborn",
}

var firstName = []string{
	"Alden", "Elara", "Garrick", "Isolde", "Lysandra", "Orik", "Seraphina", "Thaddeus", "Xander", "Zephyra",
	"Aria", "Cyrus", "Elysia", "Faelan", "Gwendolyn", "Kael", "Lyra", "Malik", "Nadia", "Ryland",
	"Varian", "Elowen", "Soren", "Cassia", "Darian", "Kaelia", "Eowyn", "Thorne", "Arianna", "Finnian",
	"Elara", "Rowan", "Iliad", "Elowen", "Saelan", "Vaela", "Daeryn", "Caelia", "Korrin", "Baelor",
	"Kyra", "Thalon", "Aeliana", "Caelum", "Naelin", "Thyra", "Baelor", "Auriel", "Calian", "Maelis",
}

var lastName = []string{
	"Ironheart", "Stormrider", "Fireforge", "Moonshadow", "Ravenclaw", "Silverthorn", "Dawnstrider", "Shadowbane",
	"Dragonflame", "Starfall", "Blackthorn", "Wyrmslayer", "Thunderstrike", "Hawkeye", "Frosthammer", "Stoneshield",
	"Bloodraven", "Darkbane", "Loreseeker", "Shadowstrike", "Skydancer", "Dragonheart", "Soulreaper", "Starwhisper",
	"Swiftblade", "Wintersong", "Firebrand", "Nightshade", "Doomhammer", "Wildheart", "Stormbringer", "Blackthistle",
	"Shadowstalker", "Sundancer", "Silvershadow", "Eagleeye", "Darkthorn", "Lightrider", "Flameheart", "Thundershout",
	"Frostwalker", "Stonehearth", "Redcloak", "Whitethorn", "Soulbinder", "Bloodthorn", "Silentstrike", "Oathbreaker",
	"Voidwalker", "Ghostweaver", "Moonrider", "Starlighter", "Steelheart", "Runekeeper", "Flamesong", "Winterfury",
}
