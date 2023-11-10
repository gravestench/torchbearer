package phase

func (s *Service) createAdventurePhase() {
	s.phases["Adventure"] = Phase{
		Name:        "Adventure",
		Description: "The adventure phase of the game.",
		SubPhase: []Phase{
			{
				Name: "Listen",
				SubPhase: []Phase{
					{Description: "The game master describes the environment, situation, and immediate problem."},
				},
			},
			{
				Name:        "Explore",
				Description: "Ask questions, discuss, repeat.",
				SubPhase: []Phase{
					{Description: "The players discuss their options."},
					{Description: "The players ask the game master about their position (or the rules)."},
				},
			},
			{
				Name:        "Decide",
				Description: "The group makes a decision",
				SubPhase: []Phase{
					{Description: "The group decides on a course of action."},
					{Description: "If the players are divided between options, the party-leader makes the call."},
				},
			},
			{
				Name:        "Act",
				Description: "The group acts out the decision they made.",
				SubPhase: []Phase{
					{
						Name: "Describe to live",
						SubPhase: []Phase{
							{Description: "Each player describes the actions of their character."},
							{Description: "Players may not cite skills, abilities, or conflicts."},
						},
					},
					{
						Name: "Good Idea",
						SubPhase: []Phase{
							{Description: "Do not roll if the game master deems an action to be a good idea."},
						},
					},
					{
						Name: "Risk & Danger",
						SubPhase: []Phase{
							{Description: "If the action is risky or dangerous, then test a skill or ability."},
						},
					},
				},
			},
			{
				Name: "Test",
				SubPhase: []Phase{
					{
						Name: "Tally bonus dice before test",
						SubPhase: []Phase{
							{Name: "Fresh", Description: "+1D"},
							{Name: "Trait L1 or L2", Description: "+1D"},
							{Name: "Spend Persona", Description: "+1D to +3D"},
							{Name: "Channel Nature", Description: "+1D x Nature"},
							{Name: "Help from skill or ability", Description: "+1D"},
							{Name: "Aid from wise", Description: "+1D"},
							{Name: "Gear", Description: "+1D"},
							{Name: "Supplies", Description: "+1D"},
							{Name: "Level benefit effect", Description: "+1D"},
							{Name: "Enchanted item effect", Description: "+1D"},
							{Name: "Invocation Effect", Description: "+1D"},
							{Name: "Spell effect", Description: "+1D"},
						},
					},
					{
						Name: "Secondary characters",
						SubPhase: []Phase{
							{Name: "Help with skill or ability", Description: "+1D"},
							{Name: "Aid with a wise", Description: "+1D"},
							{Name: "Synergy", Description: "spend a fate point to learn from the test"},
						},
					},
					{
						Name: "Tally disadvantages before the test",
						SubPhase: []Phase{
							{Name: "Use a trait against yourself", Description: "-1D; earns a check; once per session"},
							{Name: "Use a trait against yourself in a versus test", Description: "+2D to opponent; earns 2 checks"},
							{Name: "Conditions", Description: "Sick -1D; Injured -1D"},
						},
					},
					{
						Name: "Successful post-roll bonuses",
						SubPhase: []Phase{
							{Name: "For applicable weapons", Description: "+1s"},
							{Name: "For applicable L3 trait", Description: "+1s"},
							{Name: "Per positive Might differential", Description: "+1s"},
							{Name: "Per positive Precedence differential", Description: "+1s"},
						},
					},
					{
						Name: "Rerolls",
						SubPhase: []Phase{
							{Name: "Luck", Description: "Spend Fate to reroll 6s"},
							{Name: "Wise", Description: "Spend Fate (one Wyrm); Spend persona (all Wyrms)"},
						},
					},
					{
						Name: "After the roll penalties",
						SubPhase: []Phase{
							{Name: "Dungeoneer or Fighter wearing a backpack", Description: "-1s"},
							{Name: "Acting in dim light or darkness", Description: "-1s"},
							{Name: "Penalty modifiers and factors", Description: "-1s (each)"},
							{Name: "Disposition for 'Exhausted' and 'Hungry and Thirsty' conditions", Description: "-1s (each)"},
						},
					},
					{
						Name: "Tied versus tests",
						SubPhase: []Phase{
							{Description: "Break tie with L3 trait"},
							{Description: "Break tie with Might in kill, capture, and drive-off conflicts"},
							{Description: "Break tie with Precedence in convince conflicts"},
							{Description: "Break tie in GM's favor (earns 2 checks)"},
							{Description: "Tie-breaker roll"},
						},
					},
				},
			},
			{
				Name: "Result",
				SubPhase: []Phase{
					{
						Name: "Passed test",
						SubPhase: []Phase{
							{Name: "Success", Description: "Player achieves their intent"},
						},
					},
					{
						Name:        "Failed test",
						Description: "Game master chooses",
						SubPhase: []Phase{
							{Name: "Twist", Description: "A new obstacle"},
							{Name: "Players gain a Condition", Description: "player intent not achieved"},
						},
					},
					{
						Name: "Advancement",
						SubPhase: []Phase{
							{Description: "Log pass or fail for advancement"},
							{Description: "Do not mark an advance if Ob 0"},
						},
					},
					{
						Name: "Loot",
						SubPhase: []Phase{
							{Description: "Loot seeded in area by GM"},
							{Description: "Loot after conflict according to might of opponents"},
							{Description: "Loot scavenged from area"},
						},
					},
				},
			},
		},
	}
}
