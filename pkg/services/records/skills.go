package records

import (
	"torchbearer/pkg/models"
)

type SkillRecords map[string]models.Record

func (st SkillRecords) AddSkill(s models.Record) {
	st[s.Name] = s
}

func (st SkillRecords) GetSkillByName(name string) (models.Record, bool) {
	skill, exists := st[name]

	return skill, exists
}

func (s *Service) defaultSkillTable() SkillRecords {
	table := make(SkillRecords)

	for _, skill := range []models.Record{
		{
			Name: "Beggar",
			Description: "Beggars are seen as worthless and talentless " +
				"by rich fools and impudent snobs, but any who have lived on " +
				"the streets understand that it takes real wits, skill and " +
				"luck to survive. Certainly a lot more than those daffodils " +
				"have. Use the Beggar skill to ply your trade on the " +
				"streets of towns and cities, appealing to the generosity, " +
				"pity or disgust of passersby to gain free food, drink or " +
				"lodging. Note that in some places, use of this skill is a " +
				"crime. Begging counts as personal business in town, but it " +
				"can be done before finding lodging.",
			BeginnersLuck: "Will",
			Tools:         []string{"Begging bowls", "bells or chimes"},
			SupportSkills: []string{"Orator", "Manipulator"},
			Factors: map[string][]models.SkillFactor{
				"Settlement Type": {
					{1, "Remote village or forgotten temple"},
					{2, "Religious bastion, prosperous wayhouse or busy crossroads"},
					{3, "Dilapidated port, caravan or Shire"},
					{4, "Bustling metropolis, walled town or borderland fortress"},
					{5, "Elfhome or dwarven halls"},
					{6, "Wizard’s tower or steading"},
				},
				"+Ask": {
					{1, "One portion of food or 1D of copper coins"},
					{2, "Stables or flophouse lodging"},
					{3, "2D of copper coins"},
				},
			},
		},
		{
			Name:          "Butcher",
			Description:   "Butchers slaughter live animals and other sentient creatures to convert them to food and supplies.",
			BeginnersLuck: "Health",
			Tools:         []string{"Cleavers", "carving knives"},
			SupportSkills: []string{"Cook", "Laborer"},
			Factors: map[string][]models.SkillFactor{
				"Type": {
					{1, "Human, sheep, rabbit or chicken"},
					{2, "Pig, ox, goat, cat, dog or cow"},
					{3, "Stag, aurochs, dwarf or horse"},
					{4, "Monster or supernatural creature"},
				},
				"+Resource": {
					{0, "Supplies for Cook, Tanner, Armorer and similar skills"},
					{1, "Portions of meat (portions listed in Animals section of Equipment)"},
				},
			},
		},
		{
			Name:          "Enchanter",
			Description:   "Truly gifted magicians and alchemists use their powers to conjure magic into mortal vessels, crafting enchanted vessels, creating potions and artifacts. Use the Enchanter skill to imbue a specially manufactured vessel with magical effects. Select a number of times the enchantment effect will work, and how long the item itself holds the magical energy before it dissipates (whether it was used or not), then choose an effect to imbue. Expended enchanted items may be recharged using the recharge factors below.",
			BeginnersLuck: "Will",
			Tools:         []string{"Gems", "gold", "jewels"},
			SupportSkills: []string{"Alchemist", "Arcanist"},
			Factors: map[string][]models.SkillFactor{
				"Vessel": {
					{1, "Paper or superb quality foodstuff"},
					{2, "Crystal, wood, alchemical solution, glass or gem"},
					{3, "Stone, bone or metal"},
					{4, "Living flesh"},
				},
				"+Number of Uses": {
					{0, "One charge"},
					{1, "1d3 charges"},
					{2, "2d6 charges"},
					{3, "Limitless"},
				},
				"+Duration of Enchantment": {
					{1, "One phase"},
					{2, "One adventure"},
					{3, "Permanent"},
				},
				"+Imbue Effect": {
					{1, "Imbue a wise, +1D to one action type, candle light, or improve leather armor"},
					{2, "Imbue item with a level 1 trait, +1s to one action type, torch light or improve chain armor, a helmet or shield"},
					{3, "Imbue the item with a level 2 trait, increase Might or Precedence for specific purpose, lantern light or improve plate armor"},
					{4, "Imbue the item with a level 3 trait, increase Precedence by one in general, increase memory palace or Urðr by one"},
					{5, "Increase Might by one in general"},
				},
				"Recharge": {
					{-1, "Obstacle equal to number of charges to be restored, up to the maximum number of original charges."},
				},
			},
		},
		{
			Name:          "Fisher",
			Description:   "Using spears, bows, hooks, lines, nets, and weirs, fisherfolk harvest fish from ponds, lakes, streams, rivers, and even the sea.",
			BeginnersLuck: "Will",
			Tools:         []string{"Hook and line", "nets", "spears", "bows", "weirs"},
			SupportSkills: []string{"Laborer", "Sailor"},
			Factors: map[string][]models.SkillFactor{
				"Environment": {
					{1, "Lake or river"},
					{2, "Pond, stream or coast"},
					{3, "Open sea"},
				},
				"Crafting Gear": {
					{1, "Makeshift spear"},
					{2, "Hook and line"},
					{3, "Net"},
				},
				"+Amount": {
					{0, "One fish"},
					{1, "A handful of fish"},
					{3, "Many fish"},
				},
				"+For Intended Prey": {
					{1, "Fish"},
					{2, "Sharks, marlin or tuna"},
					{3, "Whales and large creatures"},
				},
			},
		},
		{
			Name:          "Jeweler",
			Description:   "It is natural to wish to take the bounty of the earth and shape it into forms that highlight its beauty. Jewelers utilize gems, stones, silver, and gold to manufacture wondrous and highly prized items. To craft an item, the character must have materials equivalent to the base value of the item they wish to make.",
			BeginnersLuck: "Will",
			Consumables:   []string{"Silver", "gold", "gems", "precious stones, etc."},
			SupportSkills: []string{"Laborer", "Alchemist"},
			Factors: map[string][]models.SkillFactor{
				"Item": {
					{1, "Bracelet, armband, buckle or gold tooth"},
					{2, "Necklace, pendant, jeweled pin, gilding, silverware or coins"},
					{3, "Brooch, metal plate or tableware"},
					{4, "Tiara, diadem or crown"},
					{5, "Kolty or clockworks"},
					{6, "Jeweled headdress"},
				},
				"+Quality": {
					{0, "Simple"},
					{1, "Dazzling [+1D sale price]"},
					{2, "Intricate [+2D sale price]"},
				},
			},
		},
		{
			Name:          "Smith",
			Description:   "These stalwart craftsfolk are considered the backbone of civilization by some. They provide us the necessary tools to cook, build, and survive in this damned world.",
			BeginnersLuck: "Health",
			Tools:         []string{"Metal", "coal"},
			SupportSkills: []string{"Laborer", "Armorer"},
			Factors: map[string][]models.SkillFactor{
				"Jewelry: Item": {
					{1, "Bracelet, armband, buckle or gold tooth"},
					{2, "Necklace, pendant, jeweled pin, gilding, silverware or coins"},
					{3, "Brooch, metal plate or tableware"},
					{4, "Tiara, diadem or crown"},
					{5, "Kolty or clockworks"},
					{6, "Jeweled headdress"},
				},
				"+Jewelry: Quality": {
					{0, "Simple"},
					{1, "Dazzling [+1D sale price]"},
					{2, "Intricate [+2D sale price]"},
				},
				"Hardware: Type": {
					{1, "Supplies like horseshoes or nails"},
					{2, "Tools"},
					{3, "Fine work or delicate tools"},
					{4, "Gate"},
					{5, "Iron fence"},
					{6, "Portcullis"},
				},
				"+Hardware: Appearance": {
					{0, "Functional"},
					{2, "Appealing [+1D sale price]"},
					{3, "Beautiful [+2D sale price]"},
				},
			},
		},
		{
			Name:          "Strategist",
			Description:   "Civilization is always at war. It is an ugly truth. Strategists study the battles of the past to adopt traditions and create new ways to defeat opponents and conquer cities. Use the Strategist skill in warfare conflicts to set disposition and to roll for certain actions.",
			BeginnersLuck: "Will",
			Tools:         []string{"Maps", "historical accounts", "training manuals"},
			SupportSkills: []string{"Commander", "Scholar", "Steward"},
			Factors: map[string][]models.SkillFactor{
				"Strategy": {
					{1, "Create an order of battle"},
					{2, "Create a command structure"},
					{3, "Create a marching order"},
				},
				"Form Unit": {
					{2, "Infantry"},
					{3, "Archers, mounted infantry"},
					{4, "Cavalry, arbalesters"},
				},
				"Intelligence: Historical Knowledge": {
					{1, "Recall a famous ancient battle or general"},
					{2, "Recall a famous recent battle or general"},
					{3, "Recall an obscure battle or general"},
				},
				"Intelligence: +Settlement of Origin": {
					{1, "Walled town, religious bastion"},
					{2, "Bustling metropolis, borderland fortress"},
					{3, "Busy crossroads, dilapidated port"},
				},
				"Engagements": {
					{0, "Choose a place of battle and draw your opponent into an engagement there, test Strategist vs Strategist."},
				},
			},
		},
		{
			Name:          "Tanner",
			Description:   "Tanners cure hides to make leather and craft leather goods like belts, boots, scabbards, water skins, book covers, scroll cases, satchels, and saddles. They purchase hides from peasants and hunters, tools from smiths, and chemicals from alchemists.",
			BeginnersLuck: "Health",
			Consumables:   []string{"Hides", "unsavory chemicals"},
			SupportSkills: []string{"Laborer", "Armorer"},
			Factors: map[string][]models.SkillFactor{
				"Hide Type": {
					{1, "Cow, sheep or horse"},
					{2, "Cat, dog, goat or rabbit"},
					{3, "Aurochs, human or pig"},
					{4, "Monsters or supernatural creatures"},
				},
				"+Function": {
					{0, "Typical leather for boots or gloves"},
					{1, "Tough leather for armor"},
					{2, "Supple leather for book covers"},
				},
			},
		},
		{
			Name:          "Alchemist",
			Description:   "Alchemists are learned souls who use hermetic knowledge to create tinctures, balms, acids, transmutations, and incendiaries.",
			BeginnersLuck: "Will",
			Tools:         []string{"Formulæ"},
			SupportSkills: []string{"Lore Master", "Laborer"},
			Factors: map[string][]models.SkillFactor{
				"Vapors & Poisons: Impose Condition": {
					{1, "Afraid"},
					{2, "Angry"},
					{3, "Exhausted"},
					{4, "Sick"},
				},
				"+Vapors & Poisons: Duration": {
					{0, "One test"},
				},
				"Elixirs & Salves: Remove Condition": {
					{0, "One test"},
					{1, "Elixir/salve counts as successful recovery"},
				},
				"+Elixirs & Salves: Duration of Relief": {
					{0, "One test"},
					{1, "Effect lasts until recovered"},
					{1, "Afraid"},
					{2, "Angry"},
					{3, "Exhausted"},
				},
				"Alchemical Process:": {
					{0, "One test"},
					{1, "Elixir/salve counts as successful recovery"},
				},
				"+Alchemical Process: Process": {
					{1, "Purify with fire"},
					{2, "Distill liquid or make dye"},
					{3, "Magnetize or crystallize"},
				},
				"Inflammables & Explosives: Substance": {
					{1, "Flash powder, smoke pot, or firework"},
					{2, "Incendiary"},
					{3, "Explosive"},
				},
				"+Inflammables & Explosives: Charges": {
					{0, "One test"},
					{1, "Charges equal to skill rating"},
				},
				"+Inflammables & Explosives: Effect": {
					{1, "+1D to skill test"},
					{2, "+1s weapon (choose action type)"},
					{3, "+1 Might"},
				},
			},
		},
		// TODO :: Arcanist
		{
			Name:          "Armorer",
			Description:   "Armorers forge armor and weapons for adventurers, knights, and other soldiers. Use this skill to craft new weapons and repair damaged armor.",
			BeginnersLuck: "Health",
			Tools:         []string{"Steel rivets", "leather straps", "metal plates"},
			SupportSkills: []string{"Smith", "Laborer"},
			Factors: map[string][]models.SkillFactor{
				"Weapon Crafting": {
					// Hand-to-hand
					{1, "Dagger or hand axe"},
					{2, "Mace or battle axe"},
					{3, "Flail, sword or warhammer"},

					// polearms
					{1, "Spear or lance"},
					{2, "Halberd or polearm"},
					{3, "Great sword"},

					// missiles
					{1, "Sling"},
					{2, "Bow"},
					{3, "Crossbow"},
				},
				"Armor Crafting": {
					{1, "Shield or leather armor"},
					{2, "Helmet"},
					{3, "Chain armor"},
					{4, "Plate armor"},
				},
				"Repairing": {
					{1, "Helmet"},
					{2, "Chain armor"},
					{3, "Plate armor"},
				},
			},
		},
		{
			Name:          "Carpenter",
			Description:   "Carpenters make useful items out of wood, like chairs, doors, cabinets, ladders, joints, pulleys, levers, and even boats.",
			BeginnersLuck: "Health",
			Tools:         []string{"Lumber", "nails", "glue", "rope"},
			SupportSkills: []string{"Alchemist", "Laborer"},
			Factors: map[string][]models.SkillFactor{
				"Woodworking": {
					{1, "Cup, bowl or board"},
					{2, "Pulley, ladder, ramp or furniture"},
					{3, "Small structure like a shack"},
					{4, "Small house, bridge"},
					{5, "Long house"},
				},
				"Boats": {
					{2, "Raft"},
					{3, "Skiff with 6 sack capacity"},
					{4, "Faering with 8 sack capacity"},
				},
			},
		},
		{
			Name:          "Cartographer",
			Description:   "Cartographers create and interpret maps. This skill is essential to adventurers, explorers, and caravan masters. Cartographer is used to navigate safely through known areas without needless danger or wasted time. During the game, log the locations and features you have visited while exploring. Then test Cartographer to synthesize these notes into an accurate map. So long as you possess the map, have light to read it and you’re “on the map,” you can travel to a location without needing to test, provided no other obstacles have emerged since your last trip.",
			BeginnersLuck: "Will",
			Tools:         []string{"Ink", "paper", "parchment", "cloth", "charcoal"},
			SupportSkills: []string{"Scholar", "Pathfinder"},
			Factors: map[string][]models.SkillFactor{
				"Area": {
					{1, "Small area (a few rooms or a few terrain features)"},
					{2, "Moderate-sized area (half a level or a few days' travel)"},
					{3, "Large area (dungeon level or a region of countryside)"},
				},
				"+Information": {
					{1, "Area personally surveyed"},
					{2, "Information transmitted by notes"},
					{3, "Information transmitted by word of mouth"},
				},
			},
		},
		{
			Name:          "Commander",
			Description:   "Commanders understand how to organize, supply, and command a force of soldiers for battle. Mercenary captains and conquerors utilize this skill to lead forces in warfare. The skill is used in ambush, battle, and skirmish conflicts described in the Lore Master’s Manual, as well to give direct orders to a unit of soldiers.",
			BeginnersLuck: "Will",
			Tools:         []string{"Ammunition", "carts and wagons"},
			SupportSkills: []string{"Steward", "Orator"},
			Factors: map[string][]models.SkillFactor{
				"Command unit": {
					{1, "Veteran elite unit"},
					{2, "Experienced unit"},
					{3, "Trained unit"},
					{4, "Inexperienced untrained unit"},
				},
				"+Order": {
					{0, "March in good order"},
					{1, "Engage a superior foe"},
					{2, "Stand ground under fire or fall back"},
					{3, "Prevent troops from pillage or rout"},
				},
			},
		},
		{
			Name:          "Cook",
			Description:   "Like magicians, cooks can dispel hunger with a little kitchen magic, even when you’re out in the wild, far from home. They can make bread from a handful of grains or stew from a brace of coneys and wild taters. Essentially, a cook can stretch a single portion of rations to feed the whole group. Thus, every adventuring party needs a cook. Cook is used to prepare meals or preserve food, both from a wide range of ingredients.",
			BeginnersLuck: "Will",
			Tools:         []string{"Salt", "garlic", "spices, etc."},
			SupportSkills: []string{"Alchemist", "Laborer"},
			Factors: map[string][]models.SkillFactor{
				"Meal Prep: Ingredients": {
					{1, "Fresh food or fish"},
					{2, "Preserved food or game"},
					{3, "Forage"},
					{4, "Moldering sacks of grain or other rotten food"},
				},
				"+Meal Prep: Amount of Food to Be Preserved": {
					{1, "Preserve 1 fresh ration or game"},
					{2, "Preserve 2 fresh rations or game"},
					{3, "Preserve 4 fresh rations or game"},
				},
				"Preservation: Amount of Meals": {
					{0, "For one"},
					{1, "Small group"},
					{2, "Large group"},
					{4, "Preserve 8 fresh rations or game"},
					{5, "Preserve 16 fresh rations or game"},
					{6, "Preserve 32 fresh rations or game"},
				},
			},
		},
		{
			Name:          "Criminal",
			Description:   "Criminals are experts in questionable enterprises like smuggling, counterfeiting, picking pockets or picking locks. This skill is used to commit various crimes and detect crimes being committed.",
			BeginnersLuck: "Health",
			Tools:         []string{"Lockpicks", "silken handkerchiefs", "paint"},
			SupportSkills: []string{"Scout", "Scholar"},
			Factors: map[string][]models.SkillFactor{
				"Lock Picking": {
					{1, "Simple lock"},
					{2, "Decent lock or rusted simple lock"},
					{3, "Rusted decent lock"},
					{4, "Complex lock"},
					{5, "Rusted complex lock"},
					{6, "Masterwork lock"},
				},
				"Escaping Bonds": {
					{1, "Leather restraints"},
					{2, "Rope restraints"},
					{3, "Metal restraints"},
				},
				"Counterfeiting": {
					{1, "Coin clipping"},
					{2, "Simple die stamp"},
					{3, "Complex die stamp"},
					{4, "Counterfeit coins"},
					{5, "Simple printed material"},
					{6, "Complex printed material"},
				},
			},
		},
		{
			Name:          "Dungeoneer",
			Description:   "Dungeoneers are experts at exploring caves, dungeons, and the ruins of lost civilizations—getting themselves and their companions out alive. The Dungeoneer skill is used to traverse dangerous underground environments and to detect hazards therein.",
			BeginnersLuck: "Health",
			Tools:         []string{"Rope", "spikes", "candles", "chalk", "twine"},
			SupportSkills: []string{"Sapper", "Survivalist"},
			Factors: map[string][]models.SkillFactor{
				"Delving: Terrain": {
					{1, "Steep slopes"},
					{2, "Vertical pitches"},
					{3, "Narrow squeezes"},
					{4, "Water-filled chambers"},
					{5, "Water-filled chambers with a swift current"},
				},
				"+Delving: Adventures": {
					{0, "One person"},
					{1, "Two people"},
					{2, "Small group"},
				},
				"Detection: Environmental Detail": {
					{1, "Bad/good air"},
					{2, "Slope"},
					{3, "Direction"},
				},
			},
		},
		{
			Name:          "Fighter",
			Description:   "Fighters are trained to use weapons to slay people, beasts, and monsters. This is the skill of knights, soldiers, bandits, reavers, monster hunters, and adventurers. In kill, capture, and drive off conflicts, the Fighter skill is used for Attack and Feint actions.",
			BeginnersLuck: "Health",
			Tools:         []string{"Torches", "lanterns", "tables", "bottles", "stones", "doors"},
			SupportSkills: []string{"Hunter"},
			Factors: map[string][]models.SkillFactor{
				"Fighting": {
					{0, "One person"},
					{1, "Two people"},
					{2, "Small group"},
				},
				"Brawling": {
					{0, "For simple altercations or brawls, use Fighter in versus tests against another character’s Fighter skill (or Beginner’s Luck Health)."},
				},
			},
		},
		{
			Name:          "Haggler",
			Description:   "Hagglers live to make bargains and score the best deals at the market. The Haggler skill is used to negotiate with other characters and to gain rolls on the Market Haggling Events table in town.",
			BeginnersLuck: "Will",
			Tools:         []string{"Abacus", "scales"},
			SupportSkills: []string{"Manipulator"},
			Factors: map[string][]models.SkillFactor{
				"Bargaining in the Market": {
					{2, "Busy crossroads"},
					{3, "Religious Bastion, remote village, dwarven halls"},
					{4, "Bustling metropolis, wizard’s tower"},
					{5, "Elfhome"},
				},
			},
		},
		{
			Name:          "Healer",
			Description:   "Midwives, sergeants, and physickers all know the importance of the healing arts. The Healer skill is used to keep adventurers whole and healthy by treating wounds and illnesses and to create poultices to aid recovery. Poultices grant +1D to recovery tests when applied in camp or town.",
			BeginnersLuck: "Will",
			Tools:         []string{"Medicinal herbs", "minerals", "bandages"},
			SupportSkills: []string{"Survivalist", "Alchemist"},
			Factors: map[string][]models.SkillFactor{
				"Healing / Treating": {
					{2, "Bruises or bumps"},
					{3, "Sword cuts or broken bones"},
					{4, "Illness or burns"},
					{5, "Virulent disease"},
					{6, "Poison"},
				},
				"Mixing Poultices for...": {
					{1, "Exhausted"},
					{2, "Afraid"},
					{3, "Injured"},
					{4, "Sick"},
					{5, "Angry"},
				},
			},
		},
		{
			Name:          "Hunter",
			Description:   "Nobles, their hunters and, of course, poachers use this skill to lure, stalk, trap and slay beasts for food in forested preserves and in the wild. The Hunter skill is used to examine the nature of beasts as well as to stalk, trap or capture prey. For bringing down a beast, make a versus test using Hunter versus the beast’s Nature, or use a kill or capture conflict. Use the Hunter skill for Feint and Defend in capture conflicts. Hunters may read the aspects of a creature’s Nature. Make a versus test between Hunter and the creature’s Nature. Success indicates that they may discover a descriptor, weapon or their Might. Otherwise, use the factors below.",
			BeginnersLuck: "Health",
			Tools:         []string{"Bows", "javelins", "boar spears", "dogs", "horses", "birds of prey"},
			SupportSkills: []string{"Survivalist", "Laborer"},
			Factors: map[string][]models.SkillFactor{
				"Hunting: Ambushing": {
					{0, "Hunter vs Nature tests to ambush your prey"},
				},
				"Hunting: Stalking / Bestial Habits": {
					{1, "Trails"},
					{2, "Food or prey"},
					{3, "Weapons or nests"},
				},
				"Hunting: Trapping": {
					{1, "Snares and nets"},
					{2, "Pit"},
					{3, "Deadfall"},
				},
				"+Hunting: Rarity": {
					{1, "Common animal"},
					{2, "Rare or shy beast"},
					{3, "Exotic or magical creature"},
				},
				"+Hunting: Prey": {
					{1, "Small animals and monsters"},
					{2, "Human-size animals and monsters"},
					{3, "Large animals and monsters"},
					{4, "Massive animals and monsters"},
					{5, "Giant animals and monsters"},
				},
			},
		},
		{
			Name:          "Laborer",
			Description:   "Laborers are the bulk of the workforce in villages, towns and cities. They gather wood for the carpenters, stone for the masons and metal for the smiths. They dig ditches, carry loads and generally just do what they are told (until they get rum brave and riot). But when you hit the big score in a dungeon, laborers are who you need. The Laborer skill is used to perform manual labor, to carry more than you should, and to help trade and craft skills like Alchemist, Armorer, Carpenter, Cook, Hunter, Peasant, Sailor, Sapper, Stonemason, and Weaver.",
			BeginnersLuck: "Health",
			Tools:         []string{"Rope", "pulleys", "shims", "blocks"},
			SupportSkills: []string{"Peasant"},
			Factors: map[string][]models.SkillFactor{
				"Activity": {
					{1, "Gathering wood or tying on an extra pack 1 item"},
					{2, "Hauling water, cutting or digging"},
					{3, "Mucking out stables, burning charcoal or breaking rocks"},
					{4, "Pulling a cart, cutting salt blocks or mining ore"},
				},
				"Carrying": {
					{0, "Hands full [carried 2]"},
					{1, "+1 carried inventory slot (+1 per)"},
				},
				"Hauling": {
					{1, "Candlesticks, dinnerware or lamps"},
					{2, "Chests or boxes"},
					{3, "Rugs, tapestries or bodies"},
					{4, "Armoires or wardrobes"},
					{5, "Pianos, paintings or small statues"},
					{6, "Glass panes or mirrors"},
					{7, "Thrones"},
					{8, "Arches or large statues"},
					{9, "Quasqueton’s wall paneling"},
					{10, "Columns"},
				},
			},
		},
		{
			Name:          "Lore Master",
			Description:   "Lore masters are the keepers of arcane knowledge and deep mysteries. The Lore Master skill is used to understand the workings of magician spells, recall forgotten lore and read auras (in combination with the Supernal Vision spell). Magicians use this skill to memorize spells and place them in their memory palaces. The obstacles for memorization are described with each spell. In riddle conflicts, use Lore Master to Defend and Maneuver. Lastly, it is used in abjure and bind conflicts against spirits and demons (as described in the Lore Master’s Manual).",
			BeginnersLuck: "Will",
			Tools:         []string{"Tome of ancient lore", "book of folklore", "cyclopedia of beasts", "codex of angels", "libram of demons", "grimoire of games"},
			SupportSkills: []string{"Arcanist", "Theologian"},
			Factors: map[string][]models.SkillFactor{
				"Ephemera": {
					{1, "Fairy or folk tales"},
					{2, "Curses"},
					{3, "Enchanted places, magical phenomena"},
					{4, "Magical or arcane symbols"},
				},
				"Aura reading / Supernal Vision": {
					{1, "Reading a spell’s aura"},
					{2, "Reading an invocation’s aura"},
					{3, "Reading a potion or scroll’s aura"},
					{4, "Reading magical weapons or armor’s aura"},
				},
			},
		},
		{
			Name:          "Manipulator",
			Description:   "Manipulators use lies, half-truths, ugly truths, soothing platitudes, seduction, and intimidation to get what they want. The Manipulator skill is used to intimidate, defuse, bluff, or trick someone and to Feint and Maneuver in riddle, convince crowd, and convince conflicts.",
			BeginnersLuck: "Will",
			Tools:         []string{"Wine", "drugs", "coin", "disguises"},
			SupportSkills: []string{"Haggler", "Persuader"},
			Factors: map[string][]models.SkillFactor{
				"Intimidate or Trick": {
					{0, "Manipulator vs Manipulator (or BL Will)"},
				},
				"Grift or Bluff": {
					{0, "Manipulator vs Persuader, Haggler, or Manipulator (or BL Will)"},
				},
			},
		},
		{
			Name:          "Mentor",
			Description:   "Mentors know how to teach. Using this skill, a mentor can offer their students advancement tests or teach them new spells. In order to teach, a mentor must have the skill they’re teaching at a higher rating than their student (or the spell being taught, of course). When teaching skills, a passed Mentor test allows the instructor to offer a pass or fail for advancement for the skill being taught. It’s their choice. This process can also be used to grant a student a test toward a skill they’re learning. When teaching spells, a successful test puts a new spell in the student’s memory palace.",
			BeginnersLuck: "Will",
			Tools:         []string{"Instruction manuals", "diagrams", "equipment"},
			SupportSkills: []string{"Persuader"},
			Factors: map[string][]models.SkillFactor{
				"MENTOR FACTORS Teaching": {
					{0, "The obstacle for instruction is the student’s current Nature rating."},
				},
			},
		},
		{
			Name:          "Orator",
			Description:   "Orators make speeches to sway crowds. The Orator skill is used to convince a crowd, wrest control of a crowd from another speaker and to Attack and Defend in convince crowd conflicts.",
			BeginnersLuck: "Will",
			Tools:         []string{"Pre-written speeches", "parables", "myths"},
			SupportSkills: []string{"Manipulator"},
			Factors: map[string][]models.SkillFactor{
				"Swaying Opinion": {
					{0, "To sway a crowd, use Orator versus Orator or crowd’s Beginner’s Luck Will if they are unskilled. Use the lowest Will rating or the highest Orator rating."},
				},
			},
		},
		{
			Name:          "Pathfinder",
			Description:   "Pathfinders make and mark paths through the wilderness (use the Cartographer and Dungeoneer skills to find your way underground). The Pathfinder skill is used to make a journey overland and to Feint and Maneuver in outdoor flee and pursue conflicts.",
			BeginnersLuck: "Health",
			Tools:         []string{"Maps", "sunstones", "travelogues", "diaries"},
			SupportSkills: []string{"Scout", "Cartographer"},
			Factors: map[string][]models.SkillFactor{
				"Destination": {
					{1, "Nearby"},
					{2, "Short journey"},
					{3, "Long journey"},
					{4, "Remote or isolated"},
				},
				"+Route Conditions": {
					{1, "Well-traveled"},
					{2, "Infrequently used"},
					{3, "Overgrown or washed out"},
					{4, "Blazing a new trail"},
				},
			},
		},
		{
			Name:          "Peasant",
			Description:   "Peasants are the backbone of society. These people grow the food and crops that enable civilization to exist. Use the Peasant skill to build or mend, undertake farm work and complain.",
			BeginnersLuck: "Health",
			Tools:         []string{"Draft animals", "baskets", "sacks", "your lazy son"},
			SupportSkills: []string{"Laborer"},
			Factors: map[string][]models.SkillFactor{
				"Work": {
					{1, "Herding or tending"},
					{2, "Ditch digging, planting or harvesting"},
					{3, "Calving or foaling"},
				},
				"Crafting: Building and Mending": {
					{1, "Socks or clothing"},
					{2, "Fences or tools"},
					{3, "Wheels or axles"},
					{4, "Huts or sheds"},
					{5, "Friendships"},
				},
				"+Crafting: Season": {
					{0, "Summer"},
					{1, "Autumn"},
					{2, "Spring"},
					{3, "Winter"},
				},
				"Complaining About...": {
					{1, "Taxes or roleplaying games"},
					{2, "Spouses or weather"},
					{3, "Priests"},
					{4, "Strange neighbors"},
					{5, "Complaining"},
				},
			},
		},
		{
			Name:          "Persuader",
			Description:   "Persuade your friends to act on your behalf. Explain to them why it’s in their best interest. The Persuader skill is used to assuage someone’s doubts, to enlist them to your side and to Attack and Defend in convince conflicts.",
			BeginnersLuck: "Will",
			Tools:         []string{"Wine", "coin", "gifts"},
			SupportSkills: []string{"Manipulator"},
			Factors: map[string][]models.SkillFactor{
				"Persuading": {
					{0, "Persuader vs Persuader or Manipulator (or Beginner’s Luck Will)"},
				},
			},
		},
		{
			Name:          "Rider",
			Description:   "Knights, horse archers and goblin wolf-riders are trained to control their mounts in chaotic situations. Use the Rider skill to train mounts and to make use of them during conflicts. Also, while mounted, use Rider to Attack and Feint in pursue and flee conflicts. To urge a mount into combat, use the Mounted Combat factors listed and test Rider when equipping weapons. Success indicates you may benefit from your mount. Failure indicates you must use your own abilities, and you count as unarmed for the first round. A properly trained and controlled mount grants the benefits of its Nature, Might, and weapons to its rider. In kill, capture, drive off, pursue, and flee conflicts, you may use your mount’s Might and weapons instead of your own. To fully train a beast to ride or for war, requires a number of camp or town phases equal to their Nature rating.",
			BeginnersLuck: "Health",
			Tools:         []string{"Food", "treats", "pets", "Rope", "reins", "muzzles"},
			SupportSkills: []string{"Peasant"},
			Factors: map[string][]models.SkillFactor{
				"TRAINING FACTORS Nature": {
					{0, "Obstacle equal to beast’s Nature"},
				},
				"Combat: Trained for...": {
					{1, "War (1)"},
					{2, "Riding (2)"},
					{3, "Untrained or domesticated (3)"},
				},
				"+Combat: Use": {
					{1, "Domestication (1)"},
					{2, "Riding (2)"},
					{3, "War (3)"},
				},
				"+Combat: Conflict": {
					{0, "Flee or pursue (0)"},
					{1, "Capture (1)"},
					{2, "Kill or drive off (2)"},
				},
			},
		},
		{
			Name:          "Ritualist",
			Description:   "Ritualists use prayers and chants to draw upon the power of the Immortal Lords and other primal forces to impose their will upon the world. It is the domain of theurges, shamans, and other creatures who know the secrets of beseeching the Immortal Lords. The Ritualist skill is used to officiate ceremonies, to perform invocations, and to Defend and Maneuver in banish conflicts.",
			BeginnersLuck: "Will",
			Tools:         []string{"Ritual supplies (sacramentals)"},
			SupportSkills: []string{"Theologian"},
			Factors: map[string][]models.SkillFactor{
				"OFFICIANT FACTORS Ritual": {
					{1, "Wedding (1)"},
					{2, "Coming of age ceremony (2)"},
					{3, "Funeral (3)"},
					{4, "Consecration (4)"},
				},
				"INVOCATIONS FACTORS Invocations": {
					// TODO :: Finish ritualist
					{0, "Rules for performing invocations are found in the Ritual chapter. The obstacles are listed with each invocation in the Invocations chapter of the Reference section."},
				},
			},
		},
		{
			Name:          "Sailor",
			Description:   "Sailors use wisdom, practice, and hard-won experience to navigate coasts and waterways. Use this skill to safely pilot a vessel along waterways. To use this skill, one must have a craft or be part of a crew on someone else’s craft.",
			BeginnersLuck: "Health",
			Tools:         []string{"Tar, rope, lumber", "Maps, charts, sunstones"},
			SupportSkills: []string{"Survivalist", "Laborer"},
			Factors: map[string][]models.SkillFactor{
				"Craft Type": {
					{1, "Skeid or karve (1)"},
					{2, "Raft or skiff (2)"},
					{3, "Knarr (3)"},
					{4, "Drakar or cog (4)"},
				},
				"+Season": {
					{0, "Summer (0)"},
					{1, "Spring (1)"},
					{2, "Fall (2)"},
					{3, "Winter (3)"},
				},
				"+Location": {
					{0, "Known waters (0)"},
					{1, "Charted waters (1)"},
					{3, "Uncharted waters (3)"},
				},
				"+Weather": {
					{1, "Calm, brisk or blustery (1)"},
					{2, "Winds, rain or fog (2)"},
					{3, "Storms (3)"},
					{4, "Ragnarok (4)"},
				},
			},
		},
		{
			Name:          "Sapper",
			Description:   "Sappers are experts in the unpredictable dynamics of digging and defending in the dark tunnels below. Use the Sapper skill to dig tunnels, collapse them, and set traps for the unwary.",
			BeginnersLuck: "Health",
			Tools:         []string{"Sulphur, lumber, grease"},
			SupportSkills: []string{"Alchemist", "Laborer"},
			Factors: map[string][]models.SkillFactor{
				"Tunneling: Type": {
					{1, "Crawl-way (1)"},
					{2, "Shaft (2)"},
					{3, "Tunnel (3)"},
				},
				"+Tunneling: Length": {
					{1, "Short (1)"},
					{2, "Long (2)"},
				},
				"+Tunneling: Material": {
					{0, "Earth (0)"},
					{1, "Clay (1)"},
					{2, "Stone (2)"},
					{3, "Sand (3)"},
				},
				"Setting Traps": {
					{1, "Pit (1) + Material factor above"},
					{2, "Tripwire alarm (2)"},
					{3, "Deadfall (3)"},
					{4, "Spear or crossbow mechanism (4)"},
					{5, "Gas or smoke mechanism (5)"},
					{6, "Explosives (6)"},
				},
				"Disarming Traps": {
					{1, "Tripwire or open pit (1)"},
					{2, "False floor (2)"},
					{3, "Pressure plate (3)"},
					{4, "Complex, multipart mechanism (4)"},
					{5, "Explosive (5)"},
					{6, "Sigils or runes (6)"},
				},
			},
		},
		{
			Name:          "Scavenger",
			Description:   "Scavengers scrounge up useful bits and materials. They can also forage for sustenance no matter where they are in the wild. Use this skill to find what you need when there’s no market in sight.",
			BeginnersLuck: "Health",
			SupportSkills: []string{"Scout"},
			Tools:         []string{"10-foot poles, dowsing rods, lodestones"},
			Factors: map[string][]models.SkillFactor{
				"Scavenging: Material/Item": {
					{1, "Cast-off material or supplies (1)"},
					{2, "Tools or gear (2)"},
					{3, "Weapons (3)"},
					{4, "Loot [roll on Loot Table 1] (4)"},
				},
				"+Scavenging: Location": {
					{0, "Near town (0)"},
					{1, "Dungeon (1)"},
					{2, "Ruins (2)"},
					{3, "Wilderness (3)"},
				},
				"+Scavenging: Rarity": {
					{0, "Common item (0)"},
					{1, "Unusual item (1)"},
					{2, "Rare item (2)"},
				},
				"Foraging: Terrain": {
					{1, "Forests or fields (1)"},
					{2, "Ruins (2)"},
					{3, "Near town (3)"},
					{4, "Dungeons (4)"},
					{5, "Blasted wasteland (5)"},
				},
				"+Foraging: Amount": {
					{0, "One portion of forage (0)"},
					{1, "Two portions of forage (1)"},
					{2, "Three portions of forage (2)"},
				},
			},
		},
		{
			Name:          "Scholar",
			Description:   "Scholars specialize in writing accounts of events for historical records as well as plumbing the depths of libraries and archives for their secrets. Use the Scholar skill to write, read, research a subject and to scribe spell scrolls and spell books.",
			BeginnersLuck: "Will",
			SupportSkills: []string{"Lore Master, Steward"},
			Tools:         []string{"Ink, quills, brushes"},
			Factors: map[string][]models.SkillFactor{
				"Reading": {
					{1, "Journals or letters (1)"},
					{2, "Weird inscriptions (2)"},
					{3, "Histories (3)"},
					{4, "Playtest docs (4)"},
				},
				"Writing": {
					{1, "Copying weird inscriptions (1)"},
					{2, "Writing instructions or directions (2)"},
					{3, "Composing essays or letters (3)"},
					{4, "Memoirs or personal treatises (4)"},
				},
				"Scrolls and Spell Books": {
					{0, "Each individual spell is listed with its scribing obstacle"},
					{0, "If scribing multiple spells at once, use the highest circle spell as the base obstacle and add 1 per additional spell"},
				},
				"Researching a Subject": {
					{1, "Common misconceptions (1)"},
					{2, "Superficial knowledge of subject (2)"},
					{3, "Expertise about a particular aspect (3)"},
					{4, "Comprehensive understanding (4)"},
					{5, "Complete understanding of subject (5)"},
				},
				"Researching an Area": {
					{1, "Geography (1)"},
					{2, "Recent history (2)"},
					{3, "Famous lineages or heraldry (3)"},
					{4, "Ancient history (4)"},
					{5, "Obscure history or obscure lineages (5)"},
				},
			},
		},
		{
			Name:          "Scout",
			Description:   "Scouts are adept at spotting and tracking monsters on the prowl, sneaking behind enemy lines, trailing targets and finding hidden things. Use the Scout skill to detect traps, find hidden things or people, to hide yourself and to Attack and Feint in flee and pursue conflicts.",
			BeginnersLuck: "Will",
			SupportSkills: []string{"Pathfinder, Hunter"},
			Tools:         []string{"10-foot pole"},
			Consumables:   []string{"Chalk dust, candles"},
			Factors: map[string][]models.SkillFactor{
				"Sneaking or Spotting": {
					{1, "Scout versus Scout or Nature (or Beginner’s Luck Will)"},
				},
				"Searching: Size of Target": {
					{1, "Large (cave mouth, house) (1)"},
					{2, "Moderate (person, secret door) (2)"},
					{3, "Small (helmet, hammer) (3)"},
					{4, "Tiny (needle) (4)"},
				},
				"+Searching: Location of target": {
					{1, "Location known (1)"},
					{2, "Location roughly known (2)"},
					{3, "Location guessed (3)"},
				},
				"Detecting traps": {
					{2, "Obvious trap (e.g., a lock inside a stone demon’s maw with a hinged jaw and gleaming serrated teeth) (2)"},
					{3, "Concealed trap (e.g., a tripwire or pit) (3)"},
					{4, "Subtle trap (e.g., a pressure plate) (4)"},
					{5, "Devious trap (e.g., a chamber filled with odorless poisonous gas) (5)"},
				},
			},
		},
		{
			Name:          "Steward",
			Description:   "Stewards manage estates, businesses, towns, baronies and even kingdoms. This skill is used by judges, stewards, guild masters, merchants, sempstress, abbots and powerful nobles to administer domains or oversee organizations. Use this skill to write laws, account for taxes, rents and tithes, allocate funds for projects and file reports. The difficulty increases with the complexity of each settlement type.",
			BeginnersLuck: "Will",
			SupportSkills: []string{"Scholar, Theologian"},
			Tools:         []string{"Accounting ledgers, census rolls"},
			Factors: map[string][]models.SkillFactor{
				"Maintaining: Things": {
					{1, "Books or logs"},
					{2, "Rents or tithes"},
					{3, "Manifests or budgets"},
					{4, "Works and projects [grants +1 to Town Events Table]"},
					{5, "Organizing fetes and balls"},
				},
				"+Maintaining: Settlement type": {
					{1, "Sumptuary Laws"},
					{2, "Religious Laws"},
					{3, "Criminal Laws"},
					{4, "Civil Laws"},
				},
				"Laws: Writing Law": {
					{0, "Base"},
					{1, "Steading or shire"},
					{2, "Remote village or wizard’s tower"},
					{3, "Borderland fortress or religious bastion"},
					{4, "Busy crossroads, walled town"},
					{5, "Dilapidated port"},
					{6, "Bustling metropolis"},
				},
				"+Laws: Settlement Type": {
					{0, "Steading"},
					{1, "Remote village"},
					{2, "Borderland fortress or wizard’s tower"},
					{3, "Busy crossroads, dilapidated port"},
					{4, "Walled town or religious bastion"},
					{5, "Bustling metropolis"},
				},
			},
		},
		{
			Name:          "Stonemason",
			Description:   "One might say that stonemasons are the foundation of society. Use this skill to cut stone and use it to make structures like walls, bridges, arches, and buildings.",
			BeginnersLuck: "Health",
			SupportSkills: []string{"Laborer"},
			Tools:         []string{"Gloves, scaffolding, pry bar, hammer, spikes"},
			Factors: map[string][]models.SkillFactor{
				"Building: Complexity": {
					{1, "Simple structure like a wall"},
					{2, "Moderately complex like a column"},
					{3, "Complex like an arch"},
					{4, "Reinforced like fortifications"},
				},
				"+Building: Size": {
					{1, "Small like a bench"},
					{2, "Moderately-sized like a wall"},
					{3, "Large like a foundation or a house"},
				},
				"Surveying": {
					{1, "Age"},
					{2, "Sturdiness or quality"},
					{3, "Maker"},
				},
			},
		},
		{
			Name:          "Survivalist",
			Description:   "Competent survivalists know how to make their environment work for them. They are adept at making shelters, finding water, building fires, and jury-rigging tools. Survivalists also know how to read the weather and judge when it will be safe to travel.",
			BeginnersLuck: "Health",
			SupportSkills: []string{"Peasant"},
			Tools:         []string{"Poles, tarpaulins, gut, twine, stakes, dowsing rods"},
			Factors: map[string][]models.SkillFactor{
				"Making Camp: Terrain type": {
					{0, "Near town"},
					{1, "Ruins"},
					{2, "Wilderness or caves"},
					{3, "Dungeons"},
					{4, "Blasted wasteland"},
				},
				"+Making Camp: Adventurers": {
					{0, "Small group"},
					{1, "Large group"},
					{2, "Caravan or small company"},
				},
				"+Making Camp: Amenities (optional, can select multiple)": {
					{1, "Shelter"},
					{1, "Water source"},
					{1, "Concealment"},
				},
				"Actions": {
					{1, "Locate emergency shelter or a potable water source"},
					{2, "Start a fire in bad conditions"},
					{3, "Emergency tool-making"},
					{4, "Jury-rigging a boat"},
				},
				"Waiting Out Weather": {
					{2, "Rain showers or fog"},
					{3, "Snow or blustery winds"},
					{4, "Storm or gale"},
					{5, "Blizzard or thundersnow"},
					{6, "Apocalyptic weather"},
				},
			},
		},
		{
			Name:          "Theologian",
			Description:   "Theologians are masters of doctrine and of the secrets of the Immortal Lords. They know the cosmology of the heavens and the hells, the ranks of the saints and demons, and even their hidden names. Use the Theologian skill to recall doctrines and hidden names, pray at shrines and temples, and to purify your Immortal burden. The skill is also used in banish and bind conflicts found in the Lore Master’s Manual.",
			BeginnersLuck: "Will",
			SupportSkills: []string{"Scholar, Ritualist"},
			Tools:         []string{"Spirit candles, holy water, gold dust"},
			Factors: map[string][]models.SkillFactor{
				"Renown: Entity Type": {
					{1, "Young Immortal"},
					{2, "Spirits"},
					{3, "Old Immortal"},
					{4, "Ancient Immortal"},
					{5, "Jotunn"},
				},
				"+Renown: Action": {
					{1, "Offer gratitude"},
					{2, "Beg favor"},
					{3, "Propitiate"},
				},
				"Doctrine": {
					{1, "Common doctrine"},
					{2, "Obscure doctrine"},
					{3, "Foreign doctrine"},
					{4, "Secret doctrine"},
				},
				"Cosmological Knowledge": {
					{2, "Near Reaches [Aether, Faerie, Dry Lands]"},
					{3, "Middle Reaches [Astral Ocean, City of Pearl, Terminus, Citadel of Law]"},
					{4, "Far Reaches [Heavens, Hells, the Abyss]"},
				},
			},
		},
		{
			Name:          "Weaver",
			Description:   "Weavers manufacture useful textiles by hand and on the loom. Use the Weaver skill to manufacture fabric and create cloaks, aprons, blankets, sheets, and tapestries.",
			BeginnersLuck: "Will",
			SupportSkills: []string{"Laborer", "Peasant"},
			Tools:         []string{"Fur, yarn, wool, dye"},
			Factors: map[string][]models.SkillFactor{
				"Type": {
					{1, "Nets"},
					{2, "Baskets, small sacks or hats"},
					{3, "Cloaks, pockets, blankets, large sacks, sails or bedding"},
					{4, "Clothing or large basket"},
					{5, "Tapestries [requires loom]"},
				},
				"+ Style": {
					{0, "Plain [1D tapestry sale price]"},
					{1, "Ornamental [3D tapestry sale price]"},
					{2, "Fashionable [6D tapestry sale price]"},
					{3, "Stunning [12D tapestry sale price]"},
				},
			},
		},
	} {
		table.AddSkill(skill)
	}

	return table
}
