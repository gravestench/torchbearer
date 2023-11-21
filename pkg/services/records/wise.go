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
	} {
		table[record.Name] = record
	}

	return table
}
