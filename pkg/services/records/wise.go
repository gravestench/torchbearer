package records

import (
	"torchbearer/pkg/models"
)

type WisesRecords map[string]models.WiseRecord

func (s *Service) initWisesTable() WisesRecords {
	table := make(WisesRecords)

	for _, record := range []models.WiseRecord{
		{
			Name:        "Elven Lore-wise",
			Description: "The Elven Lore-wise is a knowledge skill that represents a character's deep understanding of the culture, history, and traditions of the Elves. Characters with this wise are well-versed in Elven customs, magic, and their unique way of life. They can use this knowledge to navigate Elven communities, negotiate with Elves, or decipher Elven artifacts. It can also help them identify and understand Elven-related mysteries and secrets.",
		},
		{
			Name:        "Folly of Humanity-wise",
			Description: "The Folly of Humanity-wise is a perceptive skill that highlights a character's insight into the flaws, weaknesses, and idiosyncrasies of human society. Characters with this wise have a keen awareness of human politics, biases, and moral complexities. They can use this wisdom to manipulate or outwit humans, detect hidden agendas, and navigate the treacherous waters of human interactions. It may also serve as a cautionary tool to avoid falling into common human pitfalls.",
		},
		{
			Name:        "Folly of Dwarves-wise",
			Description: "The Folly of Dwarves-wise represents a character's acute understanding of the peculiarities, quirks, and foibles of Dwarven culture and behavior. Characters with this wise possess insights into Dwarven craftsmanship, grudges, and social hierarchies. They can employ this knowledge to gain favor with Dwarves, mediate in Dwarven disputes, and uncover Dwarven secrets. It may also be used to avoid offending Dwarves or making costly mistakes when dealing with them.",
		},
		{
			Name:        "Dwarven Chronicles-wise",
			Description: "The Dwarven Chronicles-wise is a specialized knowledge skill that reflects a character's familiarity with the extensive history, legends, and lore of the Dwarves. Characters with this wise are well-versed in the epic tales, lineages of Dwarven clans, and significant events in Dwarven history. This knowledge can be valuable when seeking Dwarven allies, deciphering ancient Dwarven texts, or understanding the motivations of Dwarven individuals or groups. It provides a key to unlocking the deep and storied past of the Dwarves.",
		},
		{
			Name:        "Shrewd Appraisal-wise",
			Description: "The Shrewd Appraisal-wise represents a character's skill in assessing the true value, quality, and authenticity of items, goods, and treasures. Characters with this wise have a keen eye for detail, knowing when something is a valuable artifact or a mere trinket. They can use this wisdom to haggle for better prices, identify counterfeit items, and recognize the potential worth of seemingly insignificant objects. This wise can be invaluable in trade, bartering, and making informed decisions about what to keep or discard during adventures.",
		},
		{
			Name:        "Home-wise",
			Description: "The Home-wise represents a Halfling's intimate knowledge of their own cozy abode. Halflings take great pride in their homes and can use this wise to navigate their way around, find hidden nooks and crannies, and even uncover forgotten or secret compartments within their dwellings. This wise can be valuable for planning defenses, hiding precious items, or making visitors feel welcome.",
		},
		{
			Name:        "Needs a Little Salt-wise",
			Description: "The Needs a Little Salt-wise reflects a Halfling's culinary expertise and understanding of food preparation. Halflings are known for their culinary skills, and this wise allows them to enhance the flavor and quality of meals. It can be used to identify the missing ingredients in a dish, improve rations on the road, or even impress others with a delectable feast. Halflings understand that a little salt can make a world of difference in a meal.",
		},
		{
			Name:        "Dungeon Traps-wise",
			Description: "The Dungeon Traps-wise represents a character's expertise in recognizing and disarming traps commonly found in dungeons and ancient ruins. Characters with this wise can identify pressure plates, tripwires, and hidden mechanisms that trigger deadly traps. It's a valuable skill for delvers and treasure hunters seeking to avoid perilous surprises.",
		},
		{
			Name:        "Alchemy Ingredients-wise",
			Description: "Alchemy Ingredients-wise denotes a character's knowledge of various rare and mystical ingredients used in alchemical recipes. Characters with this wise can identify and gather exotic herbs, minerals, and other materials required for alchemical concoctions. This knowledge can be essential for alchemists and potion-makers.",
		},
		{
			Name:        "Beastmen Tribes-wise",
			Description: "The Beastmen Tribes-wise represents a character's deep understanding of the various tribes and clans of beastmen in the world. Characters with this wise can recognize the unique customs, behaviors, and hierarchies of different beastmen tribes. It can be helpful when negotiating with or avoiding conflicts with beastmen.",
		},
		{
			Name:        "Mythical Relics-wise",
			Description: "The Mythical Relics-wise represents a character's knowledge of legendary and mythical artifacts scattered throughout the world. Characters with this wise can identify, research, and understand the powers and histories of these relics. This knowledge is invaluable for adventurers seeking to uncover and harness the potential of such items.",
		},
		{
			Name:        "Insectology-wise",
			Description: "Insectology-wise is specialized knowledge about the behavior, habits, and ecological roles of insects. Characters with this wise can identify different insect species, understand their significance in ecosystems, and even predict insect-related phenomena. This knowledge can be beneficial for herbalists, trackers, and those studying the natural world.",
		},
		{
			Name:        "Goblin Markets-wise",
			Description: "Goblin Markets-wise represents a character's familiarity with the secretive and enigmatic Goblin Markets. Characters with this wise know the rules, customs, and dangers of these mysterious gatherings where rare and magical goods are traded. This wise can be helpful when navigating the intricacies of Goblin Markets and striking deals with goblin merchants.",
		},
		{
			Name:        "Gemstone Identification-wise",
			Description: "Gemstone Identification-wise represents a character's expertise in recognizing and appraising precious and semi-precious gemstones. Characters with this wise can identify various gemstone types, assess their quality, and determine their market value. This knowledge is valuable for gem traders, jewelers, and treasure hunters.",
		},
		{
			Name:        "Whispering Winds-wise",
			Description: "Whispering Winds-wise is knowledge of the secrets carried by the winds and the ability to interpret their whispers. Characters with this wise can predict weather patterns, discern hidden messages in the wind's sound, and even communicate with entities that dwell in the air. It's a skill useful for sailors, weather forecasters, and mystics.",
		},
		{
			Name:        "Ancient Martial Arts-wise",
			Description: "Ancient Martial Arts-wise represents a character's deep understanding of ancient and esoteric fighting styles. Characters with this wise can recognize martial arts techniques, understand their philosophies, and adapt their combat strategies accordingly. This knowledge can be valuable when facing skilled martial artists or deciphering ancient combat manuscripts.",
		},
		{
			Name:        "Song of the Elements-wise",
			Description: "Song of the Elements-wise is knowledge of the harmonious interplay between the four classical elements: earth, air, fire, and water. Characters with this wise can identify elemental imbalances, restore elemental harmony, and harness elemental energies without magical abilities. This skill is particularly valuable for elemental scholars and healers.",
		},
		{
			Name:        "Abyssal Symbols-wise",
			Description: "Abyssal Symbols-wise represents a character's familiarity with the dark and sinister symbols used in abyssal rituals and cults. Characters with this wise can recognize abyssal sigils, decipher cursed markings, and understand the significance of infernal symbols. It's knowledge that can be used to counteract dark magic or investigate cult activities.",
		},
		{
			Name:        "Antique Firearms-wise",
			Description: "Antique Firearms-wise is specialized knowledge of the history, mechanics, and use of antique firearms. Characters with this wise can identify and maintain historical firearms, assess their rarity and value, and even fire them with accuracy. This skill is useful for historians, collectors, and adventurers in settings with old-fashioned firearms.",
		},
		{
			Name:        "Shadow Realms-wise",
			Description: "Shadow Realms-wise reflects a character's understanding of the mysterious and shadowy dimensions that exist alongside the mortal world. Characters with this wise can navigate the shadow realms, detect breaches between worlds, and decipher the rules and dangers of these parallel planes. This knowledge can be crucial when dealing with otherworldly threats.",
		},
		{
			Name:        "Ethereal Alloys-wise",
			Description: "Ethereal Alloys-wise is knowledge of rare and ethereal metals and alloys with unique properties. Characters with this wise can identify, work with, and understand the mystical properties of such materials. This knowledge can be valuable for blacksmiths, enchanters, and those seeking to craft magical items.",
		},
		{
			Name:        "Fungi-wise",
			Description: "Fungi-wise represents a character's knowledge of various mushrooms and fungi found in the wilderness. Characters with this wise can identify edible and toxic fungi, understand their growth patterns, and even use them for medicinal or alchemical purposes. It's a skill valuable for foragers, herbalists, and survivalists.",
		},
		{
			Name:        "Poisonous Plant-wise",
			Description: "Poisonous Plant-wise is knowledge of dangerous and toxic plants that can harm or even kill when ingested or touched. Characters with this wise can recognize poisonous plants, understand their effects, and avoid accidental exposure. It's a skill for herbalists, botanists, and those who need to navigate hazardous natural environments.",
		},
		{
			Name:        "Animal Tracking-wise",
			Description: "Animal Tracking-wise represents a character's ability to track and follow the trails and signs left by animals. Characters with this wise can identify animal tracks, deduce their behaviors, and predict their movements. It's a skill useful for hunters, rangers, and those who wish to find or avoid wildlife.",
		},
		{
			Name:        "Forest Lore-wise",
			Description: "Forest Lore-wise is knowledge of the secrets and ecology of forests and woodlands. Characters with this wise can navigate forests, identify different tree species, and find sources of food and water within wooded areas. It's a skill for woodsmen, druids, and those who dwell in or explore forests.",
		},
		{
			Name:        "Geology-wise",
			Description: "Geology-wise represents a character's understanding of the geological features and formations of the land. Characters with this wise can identify rock types, predict terrain hazards, and even uncover valuable minerals and resources. It's a skill for miners, prospectors, and those who study the earth's composition.",
		},
		{
			Name:        "Weather Prediction-wise",
			Description: "Weather Prediction-wise is knowledge of weather patterns and the ability to predict upcoming weather conditions. Characters with this wise can anticipate storms, changes in temperature, and other atmospheric phenomena. It's a skill for sailors, farmers, and anyone who relies on accurate weather forecasts.",
		},
		{
			Name:        "Herbal Remedies-wise",
			Description: "Herbal Remedies-wise represents a character's proficiency in using herbs and plants for medicinal purposes. Characters with this wise can identify healing herbs, prepare poultices, and treat injuries and ailments with natural remedies. It's a skill for healers, herbalists, and those who practice traditional medicine.",
		},
		{
			Name:        "River Navigation-wise",
			Description: "River Navigation-wise is knowledge of safely navigating rivers, understanding currents, and recognizing hazards. Characters with this wise can steer boats, rafts, or canoes down rivers without mishap. It's a skill for riverboat captains, fishermen, and travelers who rely on river transport.",
		},
		{
			Name:        "Wilderness Survival-wise",
			Description: "Wilderness Survival-wise represents a character's knowledge of survival skills in the wild. Characters with this wise can build shelters, find food and water, and navigate through wilderness areas. It's a skill for outdoor enthusiasts, scouts, and those who venture into untamed landscapes.",
		},
		{
			Name:        "Edible Berries-wise",
			Description: "Edible Berries-wise is knowledge of identifying and safely consuming various wild berries. Characters with this wise can differentiate between edible and poisonous berries, assess their nutritional value, and recognize seasonal availability. It's a skill for foragers and those living off the land.",
		},
		{
			Name:        "Birdsong Interpretation-wise",
			Description: "Birdsong Interpretation-wise is the ability to understand and interpret the songs and calls of birds. Characters with this wise can discern bird behaviors, predict weather changes, and even communicate simple messages using bird calls. It's a skill for birdwatchers, ornithologists, and naturalists.",
		},
		{
			Name:        "River Fishing-wise",
			Description: "River Fishing-wise represents a character's expertise in catching fish from rivers and streams. Characters with this wise can use various fishing techniques, identify fish species, and even understand fish behavior. It's a skill for fishermen, anglers, and those who rely on riverine food sources.",
		},
		{
			Name:        "Wild Animal Behavior-wise",
			Description: "Wild Animal Behavior-wise is knowledge of the behaviors and instincts of wild animals. Characters with this wise can predict animal reactions, avoid dangerous encounters, and even communicate nonverbally with some animals. It's a skill for wildlife enthusiasts, trackers, and animal handlers.",
		},
		{
			Name:        "Floral Arrangement-wise",
			Description: "Floral Arrangement-wise represents a character's ability to create aesthetically pleasing flower arrangements. Characters with this wise can select and arrange flowers for decorative purposes, understand color harmonies, and create visually appealing displays. It's a skill for florists, gardeners, and decorators.",
		},
		{
			Name:        "Insect Identification-wise",
			Description: "Insect Identification-wise is knowledge of recognizing and understanding various insect species. Characters with this wise can identify beneficial and harmful insects, assess their roles in ecosystems, and even predict insect-related phenomena. It's a skill for entomologists, bug enthusiasts, and gardeners.",
		},
		{
			Name:        "Camouflage Techniques-wise",
			Description: "Camouflage Techniques-wise is expertise in blending into natural surroundings and avoiding detection. Characters with this wise can hide effectively in different environments, create natural disguises, and move stealthily without being noticed. It's a skill for hunters, spies, and scouts.",
		},
		{
			Name:        "Navigating the Urban Maze-wise",
			Description: "Navigating the Urban Maze-wise is knowledge of navigating complex city streets, alleyways, and urban environments. Characters with this wise can find shortcuts, identify hidden paths, and understand the layout of urban areas. It's a skill for city dwellers, thieves, and urban adventurers.",
		},
		{
			Name:        "Sailing Knots-wise",
			Description: "Sailing Knots-wise represents a character's expertise in tying and using various knots for sailing and maritime activities. Characters with this wise can secure sails, anchor ships, and perform other essential tasks on the water. It's a skill for sailors, pirates, and maritime enthusiasts.",
		},
		{
			Name:        "Folklore Legends-wise",
			Description: "Folklore Legends-wise is knowledge of local myths, legends, and folktales. Characters with this wise can recall and interpret folklore stories, understand their cultural significance, and use them for storytelling or problem-solving. It's a skill for bards, storytellers, and cultural historians.",
		},
		{
			Name:        "Craftsmanship Techniques-wise",
			Description: "Craftsmanship Techniques-wise represents a character's understanding of various crafting methods and techniques. Characters with this wise can identify craftsmanship styles, assess the quality of handmade items, and even replicate traditional crafting methods. It's a skill for artisans, craftsmen, and blacksmiths.",
		},
		{
			Name:        "Wine Tasting-wise",
			Description: "Wine Tasting-wise is knowledge of wine varieties, flavors, and the art of wine tasting. Characters with this wise can identify wine characteristics, appreciate wine pairings, and assess the quality of wines. It's a skill for sommeliers, wine connoisseurs, and gourmets.",
		},
		{
			Name:        "Cultural Etiquette-wise",
			Description: "Cultural Etiquette-wise is knowledge of social customs and etiquette in various cultures. Characters with this wise can navigate different cultural norms, show respect in foreign lands, and avoid cultural misunderstandings. It's a skill for diplomats, ambassadors, and world travelers.",
		},
		{
			Name:        "Ocean Currents-wise",
			Description: "Ocean Currents-wise represents a character's understanding of oceanic currents and their effects on navigation. Characters with this wise can use ocean currents for faster travel, predict sea conditions, and navigate safely on the open sea. It's a skill for seafarers, sailors, and oceanographers.",
		},
		{
			Name:        "Ancient Architecture-wise",
			Description: "Ancient Architecture-wise is knowledge of historical architectural styles and structures. Characters with this wise can recognize architectural features, understand construction techniques, and assess the historical significance of buildings. It's a skill for historians, archaeologists, and architecture enthusiasts.",
		},
		{
			Name:        "Dragon Lore-wise",
			Description: "Dragon Lore-wise is a knowledge skill that reflects a character's deep understanding of the culture, history, and behavior of dragons. Characters with this wise are well-versed in dragon customs, abilities, and their unique way of life. They can use this knowledge to negotiate with dragons, decipher ancient draconic texts, or unravel dragon-related mysteries and secrets.",
		},
		{
			Name:        "Serpent Whispering-wise",
			Description: "Serpent Whispering-wise is knowledge of understanding and communicating with serpents and reptilian creatures. Characters with this wise can interpret serpent behaviors, predict their movements, and even use serpent symbolism in rituals or negotiations. It's a skill for snake charmers, herpetologists, and those who interact with reptiles.",
		},
		{
			Name:        "Pirate Code-wise",
			Description: "Pirate Code-wise represents a character's expertise in the unwritten rules and customs of piracy. Characters with this wise can navigate the pirate underworld, understand pirate hierarchies, and negotiate with pirate crews. It can also help them avoid treacherous situations when dealing with the scallywags of the high seas.",
		},
		{
			Name:        "Ancient Ruins-wise",
			Description: "Ancient Ruins-wise is knowledge of exploring and understanding long-forgotten ruins and archaeological sites. Characters with this wise can identify architectural styles, decipher ancient scripts, and navigate safely through ancient structures. It's a skill for archaeologists, tomb raiders, and adventurers uncovering the secrets of the past.",
		},
		{
			Name:        "Mystic Runes-wise",
			Description: "Mystic Runes-wise represents a character's understanding of magical runes and symbols. Characters with this wise can decipher runic inscriptions, activate runic enchantments, and harness the power of runic magic. It's a skill for runemasters, wizards, and those who delve into the mystical arts.",
		},
		{
			Name:        "Celestial Navigation-wise",
			Description: "Celestial Navigation-wise is knowledge of using celestial objects for navigation. Characters with this wise can determine direction and time using the stars, sun, and moon. It's a skill for sailors, astronomers, and explorers who rely on the heavens to find their way in the world.",
		},
		{
			Name:        "Underworld Secrets-wise",
			Description: "Underworld Secrets-wise is insight into the hidden aspects of criminal organizations and illicit activities. Characters with this wise have knowledge of the criminal underworld, its hierarchies, and the codes of silence. It can help them navigate criminal networks, gather information discreetly, and avoid drawing the wrong kind of attention.",
		},
		{
			Name:        "Ancient Forest Guardians-wise",
			Description: "Ancient Forest Guardians-wise represents knowledge of the mystical creatures and guardians that protect ancient forests. Characters with this wise can identify these beings, understand their motivations, and seek their assistance or favor. It's a skill for forest shamans, protectors of nature, and those who wish to preserve sacred woodlands.",
		},
		{
			Name:        "Shipwreck Survival-wise",
			Description: "Shipwreck Survival-wise is knowledge of surviving at sea after a shipwreck or maritime disaster. Characters with this wise can find flotsam and makeshift rafts, predict currents, and increase their chances of rescue. It's a skill for sailors, adventurers, and those who traverse dangerous waters.",
		},
		{
			Name:        "Criminal Psychology-wise",
			Description: "Criminal Psychology-wise is insight into the minds and motivations of criminals. Characters with this wise can profile criminals, understand their tactics, and anticipate their actions. It can help in solving crimes, tracking down suspects, or staying one step ahead of criminal elements.",
		},
		{
			Name:        "Orchard Care-wise",
			Description: "Orchard Care-wise is knowledge of maintaining and nurturing fruit orchards. Characters with this wise can prune fruit trees, prevent diseases, and optimize fruit production. It's a skill for orchard keepers, horticulturists, and those who harvest bountiful orchards.",
		},
		{
			Name:        "Master of Disguise-wise",
			Description: "Master of Disguise-wise represents a character's expertise in assuming different identities and disguises. Characters with this wise can change their appearance, mannerisms, and behavior to blend into various social settings or deceive others. It's a skill for spies, infiltrators, and covert operatives.",
		},
		{
			Name:        "Volcanic Phenomena-wise",
			Description: "Volcanic Phenomena-wise is knowledge of volcanic activity and its effects on the environment. Characters with this wise can predict volcanic eruptions, understand lava flows, and navigate volcanic terrain. It's a skill for volcanologists, adventurers in volcanic regions, and those who seek to harness volcanic resources.",
		},
		{
			Name:        "Astrological Navigation-wise",
			Description: "Astrological Navigation-wise is knowledge of using celestial objects for precise navigation, including interstellar travel. Characters with this wise can chart courses through space, calculate launch windows, and navigate spacecraft. It's a skill for spacefarers, astronauts, and those who explore the cosmos.",
		},
		{
			Name:        "Insect Swarm Control-wise",
			Description: "Insect Swarm Control-wise is knowledge of managing and controlling swarms of insects. Characters with this wise can redirect insect swarms away from themselves or others, utilize insects for specific purposes, and understand swarm behavior. It's a skill for entomancers, pest controllers, and those who interact with insect hordes.",
		},
		{
			Name:        "Musical Instrument Crafting-wise",
			Description: "Musical Instrument Crafting-wise is expertise in crafting and repairing musical instruments. Characters with this wise can create and repair a wide range of instruments, understand acoustics, and enhance sound quality. It's a skill for luthiers, musicians, and instrument enthusiasts.",
		},
	} {
		table[record.Name] = record
	}

	return table
}
