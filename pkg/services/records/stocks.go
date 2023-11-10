package records

import (
	"torchbearer/pkg/models"
)

type StockRecords map[string]models.Stock

func (s *Service) initStockTable() StockRecords {
	table := make(StockRecords)

	for _, stock := range []models.Stock{
		{
			Race:  models.Human,
			Class: models.ClassWarrior,
			Description: "Mercenaries, soldiers, defenders, marauders — warriors have been a part of civilization " +
				"since its birth. Warriors are comfortable in armor and handy with every type of weapon. They can " +
				"grow into great commanders or cunning combatants.",
			LevelBenefits: [][]models.Benefit{
				{
					{
						Name: "Warrior",
						Description: "Warriors may wield any weapon and wear any armor, wear helmets and use " +
							"shields. When in camp, warriors may keep watch and lead the effort to prevent a " +
							"disaster for free provided they have no conditions (except fresh).",
					},
				},
				{
					{
						Name: "Brawler",
						Description: "When you are unarmed in a fight, you do not suffer the -1D penalty to all " +
							"actions. If disarmed in an ongoing kill, capture or drive off conflict, you " +
							"automatically switch to your bare hands and do not count as disarmed. Bare hands do " +
							"not provide a bonus to any actions.",
						Selectable: true,
					},
					{
						Name: "Skirmisher",
						Description: "When you’re wearing leather armor, roll an additional die to deflect a blow. " +
							"If either die comes up a success for the armor type, you absorb one point of damage.",
						Selectable: true,
					},
				},
				{
					{
						Name: "Cool Headed",
						Description: "Your veins are ice. When the game master assigns the angry condition, you do " +
							"not mark it. However, if you earn the condition through the grind, you still become " +
							"angry and suffer its effects as per the standard rules.",
						Selectable: true,
					},
					{
						Name: "Endurance",
						Description: "You’ve suffered much hardship and it has given you a certain toughness. " +
							"Endurance grants +1D to recover from injury.",
						Selectable: true,
					},
				},
				{
					{
						Name: "Stubborn",
						Description: "You’re always last out of a fight. If your hit points are depleted to zero " +
							"but one of your teammates is still in the fight, instead reduce your hit points to 1 " +
							"and deduct the remainder of the damage from another character of your choice. If there " +
							"is another character with this benefit, to see who goes out break the tie using this " +
							"rubric: higher level, higher Fighter skill, higher Nature, higher Might. If you’re the " +
							"same in all respects, you have a doppelgänger. Beware.",
						Selectable: true,
					},
					{
						Name: "Duelist",
						Description: "When using a sword, between rounds you may reassign the action type " +
							"(Attack, Defend, Feint or Maneuver) it benefits. Also, a shield, dagger or hand axe " +
							"carried in your off hand grants +1D to Defend against hand-to-hand weapons in addition " +
							"to your main weapon’s quality for this conflict action.",
						Selectable: true,
					},
				},
				{
					{
						Name: "Squire",
						Description: "A hopeful young warrior-in-training swears their service to you. Your squire " +
							"stands by you and help you with Fighter, Hunter, Commander, Rider and Mentor tests. " +
							"Add +1D to your roll when your squire helps. To gain their help in a conflict, you " +
							"must assign them a point of disposition. Your squire also has three available " +
							"inventory slots and requires a portion of food and water in camp.",
						Selectable: true,
					},
					{
						Name:        "Agile",
						Description: "In fights and battles, your help grants +2D instead of the standard +1D.",
						Selectable:  true,
					},
				},
				{
					{
						Name: "Shrug It Off",
						Description: "You have suffered many slings and arrows and are numb to the mortal coil. " +
							"Once per adventure, you may remove the injured condition. No roll is necessary. You " +
							"can do this any time before you seek help from a healer. However, once you get help " +
							"from a healer, you must abide by the standard rules for that injury.",
						Selectable: true,
					},
					{
						Name: "War Captain",
						Description: "When acting as a hero in an ambush, skirmish or battle, and leading from the " +
							"front, you double your hero bonus to +2s for Attack actions.",
						Selectable: true,
					},
				},
				{
					{
						Name: "Focus",
						Description: "Your steely focus is legendary. Choose one action type: Attack, Defend, Feint " +
							"or Maneuver. For kill, drive off, ambush, skirmish and battle conflicts, you gain +1s " +
							"to that action type. The focus benefit combines with other bonuses from weapons, " +
							"spells, invocations, armor, etc.",
						Selectable: true,
					},
					{
						Name: "Expert",
						Description: "Increase your rating cap for Fighter, Commander or Strategist from 6 to 7. " +
							"Choose two. You may advance these ratings to 7 with six passed tests and five " +
							"failed tests.",
						Selectable: true,
					},
				},
				{
					{
						Name: "Giant Slayer",
						Description: "Increase your Might by one when you are the conflict captain. Combine this " +
							"bonus with other bonuses from spells, invocations and magical items.",
						Selectable: true,
					},
					{
						Name: "Veteran",
						Description: "Your experience makes you a wily opponent. Use your level instead of your Will " +
							"or Health rating as your base disposition for kill, capture and drive off conflicts.",
						Selectable: true,
					},
				},
				{
					{
						Name: "Heroic Ability",
						Description: "Choose Fighter, Health, Commander or Survivalist. The chosen ability or skill " +
							"becomes “heroic.” When rolling the heroic ability or skill, a 3-6 indicates success " +
							"(rather than the standard 4-6).",
						Selectable: true,
					},
					{
						Name: "Transformed",
						Description: "Change one Nature descriptor to Fighting, Drinking or Looting. You may use " +
							"this benefit to replace a lost Nature descriptor.",
						Selectable: true,
					},
				},
				{
					{
						Name: "Presence",
						Description: "When you’re in a kill, drive off or battle conflict, your side gets +2 to " +
							"disposition in addition to whatever is rolled. You do not need to lead the conflict.",
						Selectable: true,
					},
					{
						Name: "Fame and Glory",
						Description: "Tales of your deeds resound throughout the lands. Increase your Precedence " +
							"by 1. Your Circles becomes heroic when looking for warriors and persons of quality to " +
							"add to your retinue (when rolling, 3-6 indicates success rather than the standard 4-6). " +
							"In addition, 2d3×10 soldiers flock to your banner, eager to act as your honor guard and " +
							"share in your glory. Their ranks swell during each respite. Roll 1d6. On a 4-6, 10 more " +
							"soldiers join your cause.",
						Selectable: true,
					},
				},
			},
		},
		{
			Race:  models.Human,
			Class: models.ClassTheurge,
			Description: "Theurges are possessed by a shard of Immortal power. This possession allows them to " +
				"invoke powers beyond the capabilities of other mortals, but such power carries with it terrible " +
				"consequences.",
			LevelBenefits: [][]models.Benefit{
				{
					{
						Name: "Theurge",
						Description: "Theurges channel divine might through relics and their bodies. They have a " +
							"capacity for divine punishment called Urðr. Their Urðr starts at 1, and they begin the " +
							"game with two minor relics. They may wield any one type of weapon except bows or " +
							"crossbows. They may not wear armor at first level, but they may use shields.",
					},
				},
				{
					{
						Name: "Militant",
						Description: "You may use one additional weapon type (any except bows or crossbows), and you " +
							"can wear helmets and use any type of armor.",
						Selectable: true,
					},
					{
						Name: "Vow to the Lords of Justice",
						Description: "The Immortals of Law and Justice have cursed you with a sight that penetrates " +
							"the heart. Once per phase, you may make a Will test versus a target’s Nature. If " +
							"successful, the game master reveals one of the following: creed, a trait or the last " +
							"evil deed committed by the target. This test does not cost a turn.",
						Selectable: true,
					},
				},
				{
					{
						Name:        "Acolyte",
						Description: "Increase your Urðr by one.",
						Selectable:  true,
					},
					{
						Name: "Vow to the Lords of Healing and Fire",
						Description: "By laying hands on the sick or exhausted, you grant the unfortunate soul a " +
							"free recovery test for the sick or exhausted conditions. The power can be used once " +
							"per camp or town phase. Conditions must be alleviated according to the recovery rules.",
						Selectable: true,
					},
				},
				{
					{
						Name: "Vow to the Lords of Light and Darkness",
						Description: "With a supplication to the Lords of Light and Dark, you cause your weapon to " +
							"glow with a holy dweomer. The glow provides light equivalent to a torch and lasts for " +
							"three turns. At 6th level this benefit acts as a lantern for the phase. You may " +
							"extinguish the light when you so desire, but it will also fail if your faith, resolve " +
							"or vows waiver. You may use this benefit once per adventure phase.",
						Selectable: true,
					},
					{
						Name: "Feared in Hell",
						Description: "If made afraid, you may make an Ob 3 Will test to channel your fear into " +
							"anger. If successful, change the afraid condition to angry. This test does not cost a " +
							"turn and may be done in any phase once per session (and you may only make this attempt " +
							"once per condition). You may not engage this benefit if already angry. If a spirit, " +
							"demon or undead creature witnesses your transformation, they must pass an Ob 3 Nature " +
							"test or flee from your presence.",
						Selectable: true,
					},
				},
				{
					{
						Name:        "Exorcist",
						Description: "Increase your Urðr by one.",
						Selectable:  true,
					},
					{
						Name: "Immortal Incarnation",
						Description: "Your Immortal patrons grant you a mount and companion of immense power. " +
							"Choose an appropriate form: warhorse, reindeer, great wolf or great boar. Your mount " +
							"has a Nature 6 and Might 4. The creature’s Nature/Might will increase to 7/5 at 8th " +
							"level.",
						Selectable: true,
					},
				},
				{
					{
						Name:        "Grand Master",
						Description: "Increase your Urðr by one.",
						Selectable:  true,
					},
					{
						Name: "Demonslayer",
						Description: "Increase your Might by one when contesting with spirit, demon or undead type " +
							"creatures.",
						Selectable: true,
					},
				},
				{
					{
						Name: "Immortal Armor",
						Description: "Absorb one point of damage from spells, invocations, demons, spirits and " +
							"undead (even when not targeted directly). In addition, demons, spirits and the undead " +
							"suffer -1s to all actions directed against you. You may combine this benefit with " +
							"other bonuses conferred by spells, invocations and magical armor.",
						Selectable: true,
					},
					{
						Name: "Vow to the Lady of Battle",
						Description: "Increase your rating cap for Fighter, Commander, Ritualist or Theologian from " +
							"6 to 7. You may advance the chosen skill to rating 7 with six passed tests and five " +
							"failed tests.",
						Selectable: true,
					},
				},
				{
					{
						Name: "Heroic Ability",
						Description: "Choose Commander, Fighter, Ritualist or Theologian. The chosen skill becomes " +
							"“heroic.” When rolling this skill, 3-6 indicates a success " +
							"(rather than the standard 4-6).",
						Selectable: true,
					},
					{
						Name: "Transformed",
						Description: "Change a Nature descriptor to one of the following: Questing, Judging or " +
							"Slaying. You may use this benefit to replace a lost Nature descriptor.",
						Selectable: true,
					},
				},
				{
					{
						Name: "Cult",
						Description: "Your deeds rise to the level of legend. A cult forms around your living, " +
							"holy works. Increase your Precedence by one. Found a temple to your cult in a " +
							"settlement of your choice and add it as a location to that town. Roll on the guild " +
							"table to determine its founding members. The temple acts as a home for you and your " +
							"followers. When you visit your temple, you may recruit 2d6 cultists to accompany you. " +
							"You may also use your Circles ability to add members to your cult. Take a +1D Circles " +
							"bonus when using the ability within your cult. Cult members will undertake independent " +
							"action for you or follow you, carrying your entire inventory. Cultists must be fed. " +
							"Cultists come with gear appropriate to their guild or, if found with Circles, as " +
							"appropriate to their class or station. In combat, cultists act as minions. If they are " +
							"assigned hit points, you may use the Boss Monster rule for yourself (see the Denizens " +
							"chapter of the Scholar’s Guide).",
						Selectable: true,
					},
					{
						Name: "Holy Avenger",
						Description: "Your weapon becomes imbued with your power and fury. Name your weapon and, " +
							"so long as you wield it, it grants +1D to all Fighter and Orator tests and increases " +
							"your Might by one. So long as you carry the holy avenger, you cannot be made afraid. " +
							"Upon your death, if the weapon is interred with you (and left for a suitable period), " +
							"it retains its magic. It can only be wielded by a group you designate—requiring a " +
							"certain Nature descriptor, trait, belief or creed of your choice. If grasped by any " +
							"other, the holy avenger burns their hand, injuring them and forcing them to drop the " +
							"weapon.",
						Selectable: true,
					},
				},
			},
		},
		{
			Race:  models.Human,
			Class: models.ClassMagician,
			Description: "Magicians have tapped into a special part of themselves called the memory palace. This " +
				"part of their mind allows them to capture and store magical spells and release them upon command. " +
				"This art is called arcanism, and it is a dangerous road to travel.",
			LevelBenefits: [][]models.Benefit{
				{
					{
						Name: "Magician",
						Description: "Magicians have the ability to cast arcane spells. They possess a special " +
							"ability called the memory palace. Their memory palace starts with a capacity of 1. " +
							"Magicians begin the game with three first circle spells. They may wield daggers or " +
							"staves as weapons, and they may not wear armor of any type or use shields."},
				},
				{
					{
						Name: "Gifted",
						Description: "You are suffused with raw magical energy. In a time of crisis, once per " +
							"session, you may use the Arcanist skill in place of any skill or ability. However, " +
							"failure brings an additional twist on top of the standard result assigned by the game " +
							"master. The game master chooses from the magical twists list or invents their own. The " +
							"result leaves a permanent effect.",
						Selectable: true,
					},
					{
						Name: "Needful Things",
						Description: "With a gesture of your hand, you can momentarily turn scraps and debris into " +
							"tools or useful items for use in a single skill test. This test costs a turn or check. " +
							"After the test, the scraps return to their original state. Test Arcanist with the " +
							"following factors:",
						Selectable: true,
					},
				},
				{
					{Description: "Receive a scroll containing a new spell from your mentor. The game master chooses " +
						"a spell that the magician is capable of casting."},
					{
						Name:        "Sorcerer",
						Description: "Add one spell slot to your memory palace.",
						Selectable:  true,
					},
					{
						Name: "Bind Familiar",
						Description: "Your familiar emerges from the corner of your vision and adopts you as its " +
							"own. You are bound together by fate. Your familiar will act to guide and protect you " +
							"and can understand simple commands and intentions, but it does not fully understand " +
							"your language and thus may interpret your commands to suit its personality. To gain " +
							"their help in a conflict, you must assign them a point of disposition.",
						Selectable: true,
					},
				},
				{
					{
						Name: "Mage Light",
						Description: "Once per phase, you may summon to your hand a pure light, equivalent to a " +
							"candle. The light may be transferred to a staff or crystal. At 5th level, the " +
							"illumination provided by this benefit increases to the equivalent of a torch. At 6th " +
							"level, it increases to the equivalent of a lantern. At each level thereafter, it " +
							"provides light and dim light to one additional character. Conjuring the mage light " +
							"requires an Ob 3 Arcanist test and it does not cost a turn. The light cannot be doused " +
							"by wind or weather. The magician can kindle or extinguish it at will, but it cannot be " +
							"ignited if the magician is angry or afraid and will blink out if those conditions are " +
							"earned while it is active.",
						Selectable: true,
					},
					{
						Name: "Inhuman Endurance",
						Description: "Ignore exhausted, injured or sick for one turn per session. If you pass the " +
							"test or win the conflict for which you ignored the condition, the condition dissipates. " +
							"Otherwise, it reemerges in addition to the twist or condition results from the failure.",
						Selectable: true,
					},
				},
				{
					{Description: "Receive a scroll containing a new spell from your mentor. The game master chooses " +
						"a spell that the magician is capable of casting."},
					{
						Name:        "Adept",
						Description: "Add one spell slot to your memory palace.",
						Selectable:  true,
					},
					{
						Name: "Apprentice",
						Description: "You have acquired an apprentice to aid you in your tasks. Your apprentice will " +
							"stick by you and offers help on Arcanist, Lore Master, Scholar and Alchemist tests, as " +
							"well as Will and Health tests, but not Beginner’s Luck tests (because what can an " +
							"apprentice teach a master?). Add +1D to your roll when your apprentice helps. To gain " +
							"their help in a conflict, you must assign them a point of disposition. Your apprentice " +
							"also has three inventory slots available and requires a portion of food and water in camp.",
						Selectable: true,
					},
				},
				{
					{
						Name: "Magister",
						Description: "If helping another magician with arcana, enchanting, lore or scholarship, you " +
							"provide +2D when you help instead of +1D.",
						Selectable: true,
					},
					{
						Name: "Unbridled Power",
						Description: "When casting spells like Swarm, Balefire or Hammer of Heaven, you are " +
							"granted +1s. However, such power cannot be easily contained. If you fail the test to " +
							"cast or fail a versus test, you are temporarily afflicted with one of the following:",
						Selectable: true,
					},
				},
				{
					{Description: "Receive one final scroll containing a new spell from your mentor. The game master " +
						"chooses a spell that the magician is capable of casting."},
					{
						Name:        "Magus",
						Description: "Add one spell slot to your memory palace.",
						Selectable:  true,
					},
					{
						Name:        "True Nature Unfolds",
						Description: "Increase your Might by one.",
						Selectable:  true,
					},
				},
				{
					{
						Name: "Master",
						Description: "Increase your rating cap for Arcanist, Lore Master, Alchemist or Scholar from " +
							"6 to 7. You may advance the chosen skill to rating 7 with six passed tests and five " +
							"failed tests.",
						Selectable: true,
					},
					{
						Name: "Transformed",
						Description: "Change one Nature descriptor to Creating, Destroying or Preserving. You may " +
							"use this benefit to replace a lost Nature descriptor.",
						Selectable: true,
					},
				},
				{
					{
						Name:        "Wizard",
						Description: "Add one spell slot to your memory palace.",
						Selectable:  true,
					},
					{
						Name: "Heroic Ability",
						Description: "Choose Arcanist, Lore Master, Alchemist, Enchanter, Scholar or Will. The " +
							"chosen skill or ability becomes “heroic.” When rolling, 3-6 indicates a success " +
							"(rather than the standard 4-6).",
						Selectable: true,
					},
				},
				{
					{
						Name: "Ageless",
						Description: "Suffused with aetherial magic, you age imperceptibly and for all intents will " +
							"never die from old age. In addition, the magic sustains you through trials that would " +
							"fell another being. You accumulate conditions and their effects but cannot die from " +
							"the grind. Poisons have no bite upon your dweomer-decayed form and may be consumed " +
							"with impunity.",
						Selectable: true,
					},
					{
						Name: "The Lawbreaker",
						Description: "You may instill life into inanimate matter, be it once living or forever " +
							"dead. Choose a target to animate and test Arcanist. Meeting the obstacle animates it " +
							"with one Nature descriptor. Spend margin of success to animate additional creations " +
							"(up to a maximum equal to your Will) or to add additional Nature descriptors. " +
							"Once animated, the creation will follow your orders, though each of its actions costs " +
							"a turn.",
						Selectable: true,
					},
				},
			},
		},
		{
			Race:  models.Dwarf,
			Class: models.ClassOutcast,
			Description: "These are dwarves who could not tolerate the rigid life in the holds and halls of their " +
				"ancestors. Tough and resourceful, outcasts can grow into mighty heroes or greed - mad tyrants.",
			LevelBenefits: [][]models.Benefit{
				{
					{
						Name: "Outcast",
						Description: "Outcasts may wield any weapon except great swords. They may wear any armor, " +
							"wear a helmet and use shields. Outcasts grant +1 to camp event rolls in dungeons and " +
							"dwarven-made structures.",
					},
				},
				{
					{
						Name: "Hauberk",
						Description: "When wearing chain armor, you reduce the chance that your armor is damaged " +
							"when hit (from the usual 1-3 to 1-2 to remain undamaged). In addition, your chain " +
							"armor will protect against weapons that bypass the protection of chain armor.",
						Selectable: true,
					},
					{
						Name: "Shrewd",
						Description: "Dwarves are shrewd hagglers. You may make the Haggler test during the town " +
							"phase without raising your lifestyle cost. You also ignore Precedence requirements " +
							"when haggling!",
						Selectable: true,
					},
				},
				{
					{
						Name: "Hardy Stock",
						Description: "You come from hardy stock. Add +1D to recover from the sick condition or to " +
							"any tests to resist poison or drugs.",
						Selectable: true,
					},
					{
						Name: "Miner",
						Description: "When below ground, you are granted +1D to your Nature when using your " +
							"Delving descriptor to detect deadfalls, pits, unstable tunnels, collapsing features, " +
							"bad air, seams, rock quality, sloping passages and Dwarf-made structures.",
						Selectable: true,
					},
				},
				{
					{
						Name: "Greed",
						Description: "You may invoke greed when you earn the angry condition and have a belief, " +
							"goal or instinct about an object or person. Greed grants a free test or conflict to " +
							"purchase, steal, bargain or murder for the item or person in question (but you cannot " +
							"be helped unless your companions have Greed or an appropriate instinct). If you acquire " +
							"the object in question (or murder the person), you receive a free test to recover from " +
							"angry or afraid in the next camp or town phase.",
						Selectable: true,
					},
					{
						Name: "Stubborn",
						Description: "You’re always last out of a fight. If your hit points are depleted to zero " +
							"but one of your teammates is still in the fight, instead reduce your hit points to 1 " +
							"and deduct the remainder of the damage from another character of your choice. If there " +
							"is another character with this benefit, to see who goes out break the tie using this " +
							"rubric: higher level, higher Fighter skill, higher Nature, higher Might. If you’re the " +
							"same in all respects, you have a doppelgänger. Beware.",
						Selectable: true,
					},
				},
				{
					{
						Name: "Cousin",
						Description: "Your cousin forsakes your clan to come join you or ventures from the clan " +
							"hold to try to convince you to return home. Your cousin remains at your side and helps " +
							"you with your default class skills (Fighter, Dungeoneer, Armorer, Laborer, Orator and " +
							"Scout). Add +1D to your roll when your cousin helps. To gain their help in a conflict, " +
							"you must assign them a point of disposition. Your cousin also has three available " +
							"inventory slots and requires a portion of food and water in camp.",
						Selectable: true,
					},
					{
						Name: "Tinker",
						Description: "You have a knack for dwarven craft. Reduce by one the factors for Building and " +
							"Mending and Truthtelling in the Peasant skill and the Repairing factors in the Armorer " +
							"skill. Minimum obstacle is 1.",
						Selectable: true,
					},
				},
				{
					{
						Name: "Great Endurance",
						Description: "Hunger and thirst have little effect on you. You do not suffer the -1 " +
							"disposition penalty for the hungry and thirsty condition. In addition, you cannot be " +
							"made exhausted as the result of a failed test or through the grind. You can be given " +
							"the condition by an invocation, spell or curse.",
						Selectable: true,
					},
					{
						Name: "Secret Destiny",
						Description: "In your next haul of loot, you find a dwarven or elven artifact of rare " +
							"beauty and strength that was destined for your hands. Choose:",
						Selectable: true,
					},
				},
				{
					{
						Name: "Troll Slayer",
						Description: "Increase your Might by one when fighting troll type opponents. Combine this " +
							"bonus with other bonuses from spells, invocations and magic weapons.",
						Selectable: true,
					},
					{
						Name: "War Captain",
						Description: "When acting as a hero in an ambush, skirmish or battle, and leading from the " +
							"front, you double your hero bonus to +2s for Attack actions.",
						Selectable: true,
					},
				},
				{
					{
						Name: "Make Light of Heavy Burdens",
						Description: "Increase your rating cap for Health from 6 to 7. You may advance Health to " +
							"rating 7 with six passed and five failed tests.",
						Selectable: true,
					},
					{
						Name: "Strength",
						Description: "Use your level instead of your Will or Health rating as the base for " +
							"disposition for kill, capture and drive off conflicts.",
						Selectable: true,
					},
				},
				{
					{
						Name: "Heroic Ability",
						Description: "Choose Fighter, Health, Orator or Armorer. The chosen skill or ability " +
							"becomes “heroic.” When rolling, 3-6 indicates a success (rather than the standard 4-6).",
						Selectable: true,
					},
					{
						Name: "Transformed",
						Description: "Change one Nature descriptor to Hoarding, Boasting or Ruling. You may use " +
							"this benefit to replace a lost Nature descriptor.",
						Selectable: true,
					},
				},
				{
					{
						Name:        "Made of Sterner Stuff",
						Description: "Increase your Might by one.",
						Selectable:  true,
					},
					{
						Name: "The Madness Within",
						Description: "When using greed (as in the level 4 benefit), all tests count as if they are " +
							"within your Nature. Also, you may take a second free test to satisfy your greed.",
						Selectable: true,
					},
				},
			},
		},
		{
			Race:  models.Elf,
			Class: models.ClassRanger,
			Description: "The elves spend ages wandering the dreamlands, pondering the greater mysteries of life. " +
				"Rangers are those few elves who awake from their slumber and find themselves unable to return to " +
				"sleep. Forced to wander wakeful, rangers manifest the spell songs of their people, growing more " +
				"and more powerful until they either create their own Elfhome or forsake the dream forever.",
			LevelBenefits: [][]models.Benefit{
				{
					{
						Name: "Ranger",
						Description: "Rangers may wield bows, swords, spears and daggers, wear leather and chain " +
							"armor and wear helmets. Rangers grant +1 to camp event rolls in the wilderness.",
					},
				},
				{
					{
						Name: "Born Under Silver Stars",
						Description: "Elves were born before the sun, thus they learn to orient themselves by the " +
							"moon and stars: You are granted +1D to Pathfinder, Navigator and Sailor tests.",
						Selectable: true,
					},
					{
						Name: "Essence of the Earth",
						Description: "Elves are woven from supple strands of the Skein of Destiny. They are granted " +
							"+1D to recover from exhaustion.",
						Selectable: true,
					},
				},
				{
					{
						Name: "Remember the Ancient Forms",
						Description: "You may use your Remembering Nature when repairing or rebuilding an item or " +
							"structure. The test must relate to your parents’ or mentor’s skill or the skills of " +
							"your hometown.",
						Selectable: true,
					},
					{
						Name: "Song of Soothing",
						Description: "You may use your Singing or Enchanting Nature to perform ancient spells of " +
							"healing. Use the Healer factors.",
						Selectable: true,
					},
				},
				{
					{
						Name: "Fearless",
						Description: "Some elves know no fear. When the game master assigns you the afraid " +
							"condition, you do not mark it. However, if you earn afraid through the grind, you " +
							"suffer the effects of the condition as per the standard rules.",
						Selectable: true,
					},
					{
						Name: "Athelas",
						Description: "You have been taught powerful wards against disease and poison. You do not " +
							"need to spend a check in camp or add to your lifestyle cost in town to make a test to " +
							"recover from sickness.",
						Selectable: true,
					},
				},
				{
					{
						Name: "Burning Bright",
						Description: "Reveal thine terrifying true inner spirit. Once per session, use Singing or " +
							"Enchanting Nature to instill fear in your opponents. To set the obstacle, total the " +
							"Might of the creatures the ranger wishes to target up to a total value equal to Might " +
							"plus one. If the ranger passes the test, their targets are made afraid and must " +
							"retreat or cower before them. If they wish to affect a conflict, use this benefit " +
							"before rolling for disposition. Burning Bright does not cost a turn to perform.",
						Selectable: true,
					},
					{
						Name: "Master of Dreams",
						Description: "You may use your Singing or Enchanting Nature to weave a soothing somnific " +
							"strain and put all who hear it into a deep slumber. To set the obstacle, total the " +
							"Might of the creatures the ranger wishes to target. They can target a total value " +
							"equal to twice their Might—but cannot affect any target whose Might is greater than " +
							"their own. Success indicates the victims fall into a deep sleep to be awakened only " +
							"after some hours of rest or if shaken violently. This benefit requires one turn to " +
							"perform and may not be done while fighting or in other chaotic circumstances. It must " +
							"be performed quietly or from a secret place from which the victim can hear the soft " +
							"strains of the song.",
						Selectable: true,
					},
				},
				{
					{
						Name: "Prescience",
						Description: "The past becomes the future. Once per phase in a kill, capture, drive off or " +
							"battle conflict, you may swap out one of your unrevealed actions for a different " +
							"action. You may do so at any point in the round before the opponent has revealed their " +
							"opposing action.",
						Selectable: true,
					},
					{
						Name: "Grief",
						Description: "The toll of an immortal life manifests itself as a deep well of grief within " +
							"the elf. Use your level as the base disposition (instead of your Will or Health rating) " +
							"when engaged in conflicts involving tragedy, loss or grief. In addition, if you write a " +
							"belief or goal regarding tragedy, loss or grief, all tests in service of assuaging your " +
							"grieving heart are considered to be within your Nature—even if your broken heart pushes " +
							"you beyond the bounds of reason.",
						Selectable: true,
					},
				},
				{
					{
						Name: "Elven Steed",
						Description: "You are adopted by a loyal elven steed with Nature 7 and the descriptors of " +
							"Galloping, Guiding and Defending.\n\nShe has the weapons: Sleek form: +1D Maneuver in " +
							"kill, capture, drive off, flee and pursue conflicts; and Sharp hooves: +1s Attack in " +
							"drive off conflicts.\n\nShe is trained for war and riding as per the Rider rules.\n\n" +
							"She is Might 4, but she also has a special quality that increases her rider’s Might by " +
							"1 in kill, capture, drive off and flee/pursue conflicts if her rider’s Might is equal to" +
							" or greater than hers.\n\nWhen in the wild, you may summon your steed with an Ob 3 " +
							"Singing Nature test. This test costs a turn.",
						Selectable: true,
					},
					{
						Name: "Shrug It Off",
						Description: "Pain means little to you. You may remove the injured condition once per " +
							"adventure. No roll is necessary. You can do this any time before you seek help from " +
							"a healer. However, once you get help from a healer for an injury, you must abide by " +
							"the regular rules for that injury.",
						Selectable: true,
					},
				},
				{
					{
						Name: "Sons and Daughters of Arda",
						Description: "Increase your Might by one. Increase your Might by an additional one when " +
							"using grief (the level 6 benefit).",
						Selectable: true,
					},
					{
						Name: "Transformed",
						Description: "Change one Nature descriptor to Avenging, Perceiving or Judging. You may use " +
							"this level benefit to replace a lost Nature descriptor.",
						Selectable: true,
					},
				},
				{
					{
						Name: "Heroic Ability",
						Description: "Choose Scout, Nature, Health or Survivalist. The chosen skill or ability " +
							"becomes “heroic.” When rolling this skill or ability, 3-6 indicates a success " +
							"(rather than the standard 4-6).",
						Selectable: true,
					},
					{
						Name: "Voice of Ages",
						Description: "Speak with the chthonic spirits of rock, tree, stream and field. You may ask " +
							"a spirit’s name and provenance; you may ask what it contains or once held; and you " +
							"may ask it to perform a service—guide your path, open the way, provide water, hinder " +
							"opponents, etc. Test your Will versus the spirit’s Nature:\n\nStone — Nature 3\nTree — " +
							"Nature 4\nStream or cave — Nature 5\nPond — Nature 6\nHill — Nature 7\nRiver — Nature 8",
						Selectable: true,
					},
				},
				{
					{
						Name: "Maze of Mists and Shadows",
						Description: "You know a powerful spell song that weaves a fence of mists and shadows " +
							"around a patch of land, forest or even a city. Casting this spell song requires an " +
							"Ob 4 Nature test during the town phase. You see all who enter the maze of mists and " +
							"shadows. Once alerted to the presence of an interloper, you may grant them permission " +
							"to pass through the mists without incident or you may force them to test: To pass " +
							"through the concealed area, a creature or character must succeed at a Nature test " +
							"against an obstacle equal to your max Nature rating. Failure indicates they’ve been " +
							"led astray—or can enter but with the exhausted or afraid condition (your choice). " +
							"Characters or creatures with a Might greater than yours may pass through the Maze of " +
							"Mists and Shadows against your will without being affected by it. This forced entry " +
							"breaks the enchantment.",
						Selectable: true,
					},
					{
						Name:        "Dream Vision",
						Description: "Summoning the pure spirit of your true nature, you can see impossible distances, hidden auras, truth behind lies or even project a vision into the mind of another. Test your Will against the appropriate:\n\nImmortal Vision Factors:\n\nSee impossible distances (2)\nSee an aura (3)\nDetect truth and lies (4)\nSee a future event (5)\n\nIn Dreams Factors:\n\nRemembering an event experienced by another (2)\nRecovering from all conditions (4)\nRetrieving a lost object (4)\nBreaking a curse (5)",
						Selectable:  true,
					},
				},
			},
		},
		{
			Race:  models.Halfling,
			Class: models.ClassBurglar,
			Description: "Halflings simply aren’t adventurers. It’s just not done. I mean, how can one eat second " +
				"breakfast during an adventure? However, should the unthinkable happen, halfling adventurers make " +
				"excellent burglars. They’re oft underestimated by their companions, but they possess deep wells of " +
				"patience and endurance.",
			LevelBenefits: [][]models.Benefit{
				{
					{
						Name: "Burglar",
						Description: "Burglars are useful jacks-of-all-trades. They may wield any weapon except " +
							"crossbows, great swords, halberds, and polearms. They may wear leather armor, helmets, " +
							"and use shields."},
				},
				{
					{
						Name: "Soft Step",
						Description: "When concealed in or sneaking through underbrush or scrub, opponents suffer " +
							"a -1s penalty to detect you.",
						Selectable: true,
					},
					{
						Name: "Abstemious",
						Description: "At any point you choose, you may tighten your belt and push off your hunger " +
							"or thirst. Remove the hungry and thirsty condition, but check off the angry condition.",
						Selectable: true,
					},
				},
				{
					{
						Name: "Second Breakfast",
						Description: "You may eat an extra ration of food now to hedge against hunger later. This " +
							"extra ration essentially counts as armor against one future hungry and thirsty " +
							"condition. Note this counts only when eating food, not drinking water, wine, or " +
							"whatever. If you leave town fresh, you automatically begin the adventure phase with a " +
							"second breakfast in your belly.",
						Selectable: true,
					},
					{
						Name: "Mathom",
						Description: "When given a gift of gratitude, you may demonstrate sincere thanks to the " +
							"giver and automatically recover from one of the following conditions: hungry and " +
							"thirsty, angry, afraid, or exhausted. No test or check is required, and recovery " +
							"happens regardless of recovery order. This benefit can be used once per person you " +
							"know. Gifts must be given in sincerity.",
						Selectable: true,
					},
				},
				{
					{
						Name:        "Plucky",
						Description: "When angry, you may use your Hidden Depths trait to your benefit.",
						Selectable:  true,
					},
					{
						Name: "Oft-Overlooked",
						Description: "Halflings are often overlooked by the other peoples of the world. If your " +
							"group is captured, you will be left behind. If your companions are targeted, you will " +
							"always be picked last. This goes for finding work, too!",
						Selectable: true,
					},
				},
				{
					{
						Name: "Vest Pockets",
						Description: "You’ve learned how to stow a lot in a little space. Take three additional " +
							"torso inventory slots. These special slots can hold only pack 1 items like potions, " +
							"jewelry, and daggers. What have you got in your pockets?",
						Selectable: true,
					},
					{
						Name: "Clever",
						Description: "Add +1 to your team’s disposition when you are embroiled in conflicts " +
							"involving riddles, fleeing, or sneaking. You do not have to lead the conflict to gain " +
							"this benefit.",
						Selectable: true,
					},
				},
				{
					{
						Name: "Companion",
						Description: "One of your friends hears of your exploits and defies the conservative " +
							"conventions of the Shire to come join you in the life on the road—or journeys out to " +
							"convince you to come to your senses and come home. Your companion will stick by you and " +
							"help you with your default class skills (Cook, Criminal, Fighter, Hunter, Scout, and " +
							"Scavenger). Add +1D to your roll when your companion helps. To gain their help in a " +
							"conflict, you must assign them a point of disposition. Your companion also has three " +
							"available inventory slots and requires a portion of food and water in camp.",
						Selectable: true,
					},
					{
						Name: "Helpful",
						Description: "When acting in service of your goal, you grant +2D when you help instead of " +
							"the standard +1D.",
						Selectable: true,
					},
				},
				{
					{
						Name: "It Could Be Worse",
						Description: "You may make one test per session as if you have no conditions " +
							"(barring dead). This test does not cost a turn or a check. You may be helped by your " +
							"companion (as in the L6 benefit) or other halflings but no others.",
						Selectable: true,
					},
					{
						Name: "Shoulder the Burden",
						Description: "You are a boon to your friends. During the adventure phase, if one of your " +
							"friends suffers a condition, you may take that condition on yourself so that they are " +
							"spared. If you already suffer from this condition, you cannot take it on for your friend.",
						Selectable: true,
					},
				},
				{
					{
						Name: "Supreme Confidence",
						Description: "Use your level instead of your Will or Nature rating as the base for " +
							"disposition for conflicts involving riddling, sneaking, and fleeing.",
						Selectable: true,
					},
					{
						Name: "Friend of the Powerful",
						Description: "Halflings have a knack for charming the powerful. Upon meeting a neutral or " +
							"allied potentate, if you put on charming or respectful airs, you may note the NPC as a " +
							"friend on your character sheet. No roll or test is required. This may be done in any " +
							"phase and does not cost a turn or a check.",
						Selectable: true,
					},
				},
				{
					{
						Name: "Heroic Ability",
						Description: "Choose Will, Scout, or Cook. The chosen ability or skill becomes “heroic.” " +
							"When rolling, 3-6 indicates a success (rather than the standard 4-6).",
						Selectable: true,
					},
					{
						Name: "Transformed",
						Description: "Change one Nature descriptor to Cooking, Eating, or Adventuring. You may use " +
							"this benefit to replace a lost Nature descriptor.",
						Selectable: true,
					},
				},
				{
					{
						Name: "Stalwart",
						Description: "When helping another character and they fail the test, you may invoke this " +
							"benefit. Instead of the standard failure procedure, the result becomes a twist " +
							"dictating that you step in and make the test yourself. No other twist or condition is " +
							"given until after you have made your test. Your friend cannot help, but your success " +
							"or failure counts for both of you. You may use this benefit once per session.",
						Selectable: true,
					},
					{
						Name: "Humble Home",
						Description: "You return home to find you have inherited a fine home from a long-forgotten " +
							"distant relative. The home is located in your hometown and is equivalent to an inn. " +
							"This turn of events makes you somewhat of a respectable person. Increase Precedence by " +
							"1. In addition, increase your hometown Resources bonus and your reputation Circles " +
							"bonus to +2D.",
						Selectable: true,
					},
				},
			},
		},
	} {
		table[stock.String()] = stock
	}

	return table
}
