package records

import (
	"torchbearer/pkg/models"
)

type TraitRecords map[string]models.TraitRecord

func (s *Service) initTraitTable() TraitRecords {
	table := make(TraitRecords)

	for _, trait := range []models.TraitRecord{
		{
			Name: "Adventurous",
			Description: "Most folks are content to live a quiet life at home, but some fools actually crave the " +
				"excitement of a rollicking adventure. They live to travel to strange places and experience new " +
				"things. It can be hard to grind an adventurous spirit down — they’re always ready to see what’s over " +
				"the next hill — but sometimes they jump into a situation without looking.",
		},
		{
			Name: "Bitter",
			Description: "Some turn bitter in their travels and grow to feel all their efforts are for nothing. This " +
				"bitterness may protect them from the many and varied disappointments of life as an adventurer, but " +
				"it also burdens them. They have trouble taking the optimistic course.",
		},
		{
			Name: "Bold",
			Description: "The bold rush to action without a thought for their own safety. Daring stratagems and " +
				"reckless abandon are hallmarks of the bold. Forethought and caution are not for these souls.",
		},
		{
			Name: "Born of Earth and Stone",
			Description: "Dwarves were shaped by their maker from the bones of the earth, and they share many " +
				"qualities with earth and stone. Dwarves are steadfast and sturdy in all things, with a special " +
				"feel for metal and stone, but they can also be stubborn and unyielding.",
		},
		{
			Name: "Brave",
			Description: "The brave never hesitate to step into the unknown, but they are susceptible to wild tales " +
				"of far-off places and dangerous adventures.",
		},
		{
			Name: "Calm",
			Description: "Calm souls are difficult to anger and easily shrug off any emotion that does manage to " +
				"burden their heart. However, they risk becoming too sedate and perhaps missing a threat.",
		},
		{
			Name: "Cunning",
			Description: "The cunning are adept at deceit and at plotting traps. They are often arrogant and " +
				"underestimate their opponents.",
		},
		{
			Name: "Curious",
			Description: "Curious sorts are always eager to learn and are on the lookout for the new or the " +
				"mysterious. But there are many tales of those whose curiosity led them and their friends to an " +
				"abrupt and violent end.",
		},
		{
			Name: "Defender",
			Description: "Monsters and villains are a fact of life in the wide world. Many are the townsfolk who " +
				"have been called upon to defend their homes without training, support or even proper equipment. " +
				"These people know the value and cost of defending one’s home. Of course, when forced to choose, " +
				"these defenders often put their home above other concerns.",
		},
		{
			Name: "Early Riser",
			Description: "Early risers are up first in the morning. Awake before the sun, they’re clear-headed and " +
				"sharp while everyone is bleary-eyed and sleepy. Of course, they must be early to bed to be so " +
				"early to rise.",
		},
		{
			Name: "Extravagant",
			Description: "Some folks have expensive tastes and spend profligately. They’re experts at grand " +
				"gestures but sometimes find it difficult to show restraint when required.",
		},
		{
			Name: "Extrovert",
			Description: "Extroverts love meeting people or introducing themselves to strangers. However, they " +
				"often lead public or semipublic lives and are easy to track down, even when they don’t want " +
				"to be found.",
		},
		{
			Name: "Fearless",
			Description: "Fearless fighters hold their ground in the face of unimaginable danger. However, they " +
				"often put valor before discretion even in delicate matters.",
		},
		{
			Name: "Fiery",
			Description: "Fiery individuals have a penchant for going where they shouldn’t and doing what they " +
				"mustn’t — sticking their heads in holes or reading strange, eldritch texts. Remarkably, some manage " +
				"to skirt the consequences of their actions. But that may be because few people write stories about " +
				"dead foolhardy adventurers.",
		},
		{
			Name: "Generous",
			Description: "Generous souls always seem to have something to give to others, even if it means going " +
				"without themselves.",
		},
		{
			Name: "Heart of Battle",
			Description: "Warriors are creatures of action, able to throw their bodies into harm’s way with skill " +
				"and valor in pursuit of their goals. But often warriors too readily resort to violence when another " +
				"way might serve them better.",
		},
		{
			Name: "Hidden Depths",
			Description: "On the surface, halflings may appear to be a naive and unassuming people, but they have " +
				"hidden depths of will and character. They bear up remarkably well against pain, fatigue, and " +
				"despair that would cause other peoples to break. Sometimes the hidden depths refer to their " +
				"stomachs needing to be filled with cheese and beer.",
		},
		{
			Name: "Honorable",
			Description: "Some are known throughout the land for their upstanding character and peerless integrity. " +
				"But those who stand tallest are ripe for the greatest fall. Honor does not ensure survival.",
		},
		{
			Name: "Jaded",
			Description: "The adventuring life hardens some and makes them callous. This disposition protects them " +
				"from the folly of youthful ideals and heroism, but it also makes them dismissive of new ideas.",
		},
		{
			Name: "Loner",
			Description: "Loners have trouble working well with others. They’re at their best when no one knows " +
				"what they’re up to. Sometimes, of course, they get in over their heads and could really use a " +
				"helping hand.",
		},
		{
			Name: "Lost",
			Description: "Hard to believe, but there are those among us who have no sense of direction. They’re " +
				"useless at orienteering, but conversely, their wandering treks make them difficult to track down.",
		},
		{
			Name: "Quick-Witted",
			Description: "Quick-witted adventurers act on instinct, without need for thought or consideration. " +
				"While this attribute is clearly useful, it can lead to difficulties when patience and planning " +
				"are required.",
		},
		{
			Name: "Quiet",
			Description: "Some people are quiet in everything they do — the way they speak, the way they work, and " +
				"the way they walk. Some sink so far into quietude that they have difficulty coming out of their " +
				"shells.",
		},
		{
			Name: "Rough Hands",
			Description: "A hard life of toil toughens the hands. Such hands mark them as working folk, which can " +
				"cause them to disdain those who pursue softer arts.",
		},
		{
			Name: "Scarred",
			Description: "Survivors of terrible wars are often scarred by their experiences. They’re tough and not " +
				"easily flustered by injury or fear, but they are also maimed or traumatized by their experiences.",
		},
		{
			Name: "Sharp-Eyed",
			Description: "Sharp-eyed adventurers are always welcome in a party. They make good scouts or hunters, " +
				"but sometimes, after staring at the brush for days on end, they can get a little jumpy.",
		},
		{
			Name: "Skeptical",
			Description: "Skeptical souls are always watching for lies and deceit — a good quality — but they can be " +
				"overly cautious and have trouble trusting others.",
		},
		{
			Name: "Steady-Handed",
			Description: "Those who create fine or delicate work are often known for their steady hands. These " +
				"sorts can become fearful of damaging these wondrous instruments.",
		},
		{
			Name: "Stoic",
			Description: "Stoic souls never complain about the hardships of life on the road. They accept them and " +
				"soldier on. However, this same quality can make them emotionally remote and difficult to reach in " +
				"matters of empathy, love, and compassion.",
		},
		{
			Name: "Tall",
			Description: "Some folks are unusually straight and tall. It’s a notable attribute that grants a longer " +
				"reach, but it can be trouble when snaking one’s way through narrow tunnels in a dungeon.",
		},
		{
			Name: "Thoughtful",
			Description: "Thoughtful adventurers will ponder all options and all possible courses of action before " +
				"making a decision. This is very useful when there’s time to plan or ponder but useless in times " +
				"of haste.",
		},
		{
			Name: "Touched by the Gods",
			Description: "The Immortals speak through theurges. Moreover, when a theurge speaks to the Immortals, " +
				"the Immortals are often willing to listen. The gods guide and protect their chosen, but theurges " +
				"must be ever vigilant to ensure that the words they speak are truly those of the Immortals and not " +
				"the secret desires of their own hearts.",
		},
		{
			Name: "Wizard’s Sight",
			Description: "Magicians see more than other mortals. Whether it’s the telltale signs of a spirit or " +
				"demon or the flicker of a lie in a person’s aura, magicians can see the flow of magic and the " +
				"misty borders of the Otherworld. But they must also take care, for that which is seen with the " +
				"Wizard’s Sight cannot be unseen.",
		},
	} {
		table[trait.Name] = trait
	}

	return table
}
