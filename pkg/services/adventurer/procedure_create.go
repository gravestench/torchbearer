package adventurer

import (
	"fmt"
	"regexp"
	"strings"

	"torchbearer/pkg/models"
	"torchbearer/pkg/services/procedure"
)

const (
	keyChooseWorld                                = "world"
	keyChooseName                                 = "name"
	keyChooseStock                                = "stock"
	keyChooseHumanUpbringing                      = "human upbringing"
	keyChooseHometown                             = "hometown"
	keyChooseHometownTrait                        = "hometown trait"
	keyChooseHometownSkill                        = "hometown skill"
	keyChooseSocialGraces                         = "social graces"
	keyChooseSpecialty                            = "specialty"
	keyChooseElvenWise                            = "elven wise"
	keyChooseDwarvenWise                          = "dwarven wise"
	keyChooseHalflingWise                         = "halfling wise"
	keyChooseAdditionalWise                       = "additional wise"
	keyChooseHumanWise                            = "human wise"
	keyChooseElvenNature1                         = "elven nature 1"
	keyChooseElvenNature2                         = "elven nature 2"
	keyChooseElvenNature3                         = "elven nature 3"
	keyChooseElvenNature3Option                   = "elven nature 3 option"
	keyChooseDwarvenNature1                       = "dwarven nature 1"
	keyChooseDwarvenNature2                       = "dwarven nature 2"
	keyChooseDwarvenNature3                       = "dwarven nature 3"
	keyChooseHalflingNature1                      = "halfling nature 1"
	keyChooseHalflingNature2                      = "halfling nature 2"
	keyChooseHalflingNature3                      = "halfling nature 3"
	keyChooseHumanNature1                         = "human nature 1"
	keyChooseHumanNature2                         = "human nature 2"
	keyChooseHumanNature2option                   = "human nature 2 option"
	keyChooseHumanNature3                         = "human nature 3"
	keyChooseHumanNature3Option                   = "human nature 3 option"
	keyChooseRelationships1                       = "relationships 1"
	keyChooseRelationshipFriend1                  = "relationship friend1"
	keyChooseRelationshipFriend2                  = "relationship friend2"
	keyChooseRelationshipFriendTownsfolk1         = "relationship friend townsfolk 1"
	keyChooseRelationshipFriendTownsfolk2         = "relationship friend townsfolk 2"
	keyChooseRelationshipFriendAdventurer         = "relationship friend adventurer"
	keyChooseRelationshipFriendAdventurerLastSeen = "relationship friend adventurer lastseen"
	keyChooseRelationshipParents                  = "relationship parents"
	keyChooseRelationshipMentor                   = "relationship mentor"
	keyChooseRelationshipEnemy                    = "relationship enemy"
)

type procedureCreateAdventurer struct {
	Adventurer *models.Adventurer
	*procedure.Procedure
	service *Service
}

func (s *Service) CreateAdventurerProcedure() *procedureCreateAdventurer {
	instance := &procedureCreateAdventurer{
		service:    s,
		Adventurer: s.NewAdventurer(),
		Procedure:  procedure.New("Create Adventurer"),
	}

	// steps push additional steps onto the procedure stack
	// as they are completed. the last step won't, and is
	// implicitly terminal
	instance.PushStep(instance.chooseWorld())

	instance.OnComplete = func() {

		if err := s.AddAdventurer(instance.Adventurer); err != nil {
			s.logger.Error().Msgf("creating new adventurer: %v", err)
		}
	}

	return instance
}

func (p *procedureCreateAdventurer) chooseWorld() *procedure.Step {
	var choices []procedure.Choice

	for _, w := range p.service.worlds.GetWorlds() {
		var numAdventurers int

		for _, existingAdventurer := range p.service.adventurers {
			if existingAdventurer.World == w.UUID {
				numAdventurers++
			}
		}

		desc := fmt.Sprintf("Has %d settlements, %d adventurers", len(w.Settlements), numAdventurers)
		choices = append(choices, procedure.Choice{
			Name:        w.Name,
			Description: desc,
		})
	}

	step := &procedure.Step{
		Key:     keyChooseWorld,
		Prompt:  "Select the world to create your adventurer in:",
		Choices: choices,
	}

	step.OnComplete = func() {
		if world, err := p.service.worlds.GetWorldByName(step.Answer); err == nil {
			p.Adventurer.World = world.UUID
		}

		p.PushStep(p.chooseName())
	}

	return step
}

func (p *procedureCreateAdventurer) chooseName() *procedure.Step {
	// WORLD, choose the world that this adventurer exists in
	step := &procedure.Step{
		Key:             keyChooseName,
		Prompt:          "What is your name?",
		ValidatorRegex:  `^\w+( \w+)?$`,
		ValidatorPrompt: "first name is required, last name is optional",
	}

	step.OnComplete = func() {
		p.Adventurer.Name = step.Answer
		p.PushStep(p.chooseStock())
	}

	return step
}

func (p *procedureCreateAdventurer) chooseStock() *procedure.Step {
	// STOCK
	choices := make([]procedure.Choice, 0)

	for _, stock := range p.service.records.Stocks() {
		choices = append(choices, procedure.Choice{
			Name:        stock.String(),
			Description: stock.Description,
		})
	}

	step := &procedure.Step{
		Key:     keyChooseStock,
		Prompt:  "What is your stock?",
		Choices: choices,
	}

	step.OnComplete = func() {
		if stock, err := p.service.records.GetStockByName(step.Answer); err == nil {
			p.Adventurer.Stock.Name = stock.String()
			p.Adventurer.Stock.ChosenLevelBenefits = []string{stock.LevelBenefits[0][0].Name}
		}

		if strings.Contains(strings.ToLower(step.Answer), "human") {
			p.PushStep(p.chooseHumanUpbringing())
		} else {
			p.PushStep(p.chooseHometown())
		}
	}

	return step
}

func (p *procedureCreateAdventurer) chooseHumanUpbringing() *procedure.Step {
	// HUMAN UPBRINGING (OPTIONAL, BASED ON STOCK)
	var choices []procedure.Choice

	for _, skillName := range []string{
		"Criminal",
		"Laborer",
		"Haggler",
		"Pathfinder",
		"Peasant",
		"Survivalist",
	} {
		record, err := p.service.records.GetSkillByName(skillName)
		if err != nil {
			p.service.logger.Fatal().Msgf("selecting skill %q: %v", skillName, err)
		}

		choices = append(choices, procedure.Choice{
			Name:        record.Name,
			Description: record.Description,
		})
	}

	step := &procedure.Step{
		Key:     keyChooseHumanUpbringing,
		Prompt:  "Your human upbringing has granted you one of the following skills:",
		Choices: choices,
	}

	step.OnComplete = func() {
		if record, err := p.service.records.GetSkillByName(step.Answer); err == nil {
			p.Adventurer.Skills[step.Answer] = &models.AdventurerSkill{
				RecordKey: record.Name,
				Level:     2,
			}
		}

		p.PushStep(p.chooseHometown())
	}

	return step
}

func (p *procedureCreateAdventurer) chooseHometown() *procedure.Step {
	// HOMETOWN
	stepChooseWorld, found := p.GetStepByKey(keyChooseWorld)
	if !found {
		return nil
	}

	w, err := p.service.worlds.GetWorldByName(stepChooseWorld.Answer)
	if err != nil {
		p.service.logger.Fatal().Msgf("Creating adventurer, Step %d: %v", p.StepIndex(), err)
	}

	var choices []procedure.Choice

	for _, settlement := range w.Settlements {
		choices = append(choices, procedure.Choice{
			Name:        settlement.Name,
			Description: settlement.Description(),
		})
	}

	step := &procedure.Step{
		Key:     keyChooseHometown,
		Prompt:  "Which settlement does your adventurer call home?",
		Choices: choices,
	}

	step.OnComplete = func() {
		if settlement, errGet := w.GetSettlementByName(step.Answer); errGet == nil {
			p.Adventurer.Hometown = settlement.UUID
		}

		p.PushStep(p.chooseHometownTrait())
	}

	return step
}

func (p *procedureCreateAdventurer) chooseHometownTrait() *procedure.Step {
	// HOMETOWN - TRAIT
	stepChooseWorld, found := p.GetStepByKey(keyChooseWorld)
	if !found {
		return nil
	}

	stepHometown, found := p.GetStepByKey(keyChooseHometown)
	if !found {
		return nil
	}

	w, err := p.service.worlds.GetWorldByName(stepChooseWorld.Answer)
	if err != nil {
		p.service.logger.Fatal().Msgf("Creating adventurer, Step %d: %v", p.StepIndex(), err)
	}

	hometown := w.Settlements[0]
	for _, settlement := range w.Settlements {
		if stepHometown.Answer == settlement.Name {
			hometown = settlement
			break
		}
	}

	var choices []procedure.Choice

	for _, trait := range hometown.Culture.Traits {
		choices = append(choices, procedure.Choice{
			Name:        trait.Name,
			Description: trait.Description,
		})
	}

	step := &procedure.Step{
		Key:     keyChooseHometownTrait,
		Prompt:  "What trait best describes your adventurer?",
		Choices: choices,
	}

	step.OnComplete = func() {
		if record, errGet := p.service.records.GetTraitByName(step.Answer); errGet == nil {
			if existing, found := p.Adventurer.Traits[record.Name]; found {
				existing.Level++
			} else {
				p.Adventurer.Traits[record.Name] = &models.AdventurerTrait{
					RecordKey: record.Name,
					Level:     1,
				}
			}
		}

		p.PushStep(p.chooseHometownSkill())
	}

	return step
}

func (p *procedureCreateAdventurer) chooseHometownSkill() *procedure.Step {
	// HOMETOWN - SKILL
	stepChooseWorld, found := p.GetStepByKey(keyChooseWorld)
	if !found {
		return nil
	}

	stepHometown, found := p.GetStepByKey(keyChooseHometown)
	if !found {
		return nil
	}

	w, err := p.service.worlds.GetWorldByName(stepChooseWorld.Answer)
	if err != nil {
		p.service.logger.Fatal().Msgf("Creating adventurer, Step %d: %v", p.StepIndex(), err)
	}

	hometown := w.Settlements[0]
	for _, settlement := range w.Settlements {
		if stepHometown.Answer == settlement.Name {
			hometown = settlement
			break
		}
	}

	var choices []procedure.Choice

	for _, skill := range hometown.Culture.Skills {
		choices = append(choices, procedure.Choice{
			Name:        skill.Name,
			Description: skill.Description,
		})
	}

	step := &procedure.Step{
		Key:     keyChooseHometownSkill,
		Prompt:  "What skill did you learn in your hometown?",
		Choices: choices,
	}

	step.OnComplete = func() {
		if record, errGet := p.service.records.GetSkillByName(step.Answer); errGet == nil {
			if existing, found := p.Adventurer.Skills[record.Name]; found {
				existing.Level++
			} else {
				p.Adventurer.Skills[record.Name] = &models.AdventurerSkill{
					RecordKey: record.Name,
					Level:     2,
				}
			}
		}

		p.PushStep(p.chooseSocialGraces())
	}

	return step
}

func (p *procedureCreateAdventurer) chooseSocialGraces() *procedure.Step {
	// SOCIAL GRACES
	var choices []procedure.Choice

	for _, skill := range p.service.records.Skills() {
		switch strings.ToLower(skill.Name) {
		case "haggler", "manipulator", "orator", "persuader":
			choices = append(choices, procedure.Choice{
				Name:        skill.Name,
				Description: skill.Description,
			})
		}
	}

	step := &procedure.Step{
		Key:     keyChooseSocialGraces,
		Prompt:  "How does your adventurer convince people?",
		Choices: choices,
		ShouldSkip: func() bool {
			return len(choices) < 1
		},
	}

	step.OnComplete = func() {
		if record, errGet := p.service.records.GetSkillByName(step.Answer); errGet == nil {
			if existing, found := p.Adventurer.Skills[record.Name]; found {
				existing.Level++
			} else {
				p.Adventurer.Skills[record.Name] = &models.AdventurerSkill{
					RecordKey: record.Name,
					Level:     2,
				}
			}
		}

		p.PushStep(p.chooseSpecialty())
	}

	return step
}

func (p *procedureCreateAdventurer) chooseSpecialty() *procedure.Step {
	var choices []procedure.Choice

	for _, skill := range p.service.records.Skills() {
		switch strings.ToLower(skill.Name) {
		case "cartographer", "cook", "criminal", "dungeoneer", "haggler",
			"healer", "hunter", "manipulator", "pathfinder", "persuader",
			"orator", "rider", "sapper", "scavenger", "scout", "survivalist":
			choices = append(choices, procedure.Choice{
				Name:        skill.Name,
				Description: skill.Description,
			})
		}
	}

	step := &procedure.Step{
		Key:     keyChooseSpecialty,
		Prompt:  "What is your specialty?",
		Choices: choices,
		ShouldSkip: func() bool {
			return len(choices) < 1
		},
	}

	step.OnComplete = func() {
		if record, errGet := p.service.records.GetSkillByName(step.Answer); errGet == nil {
			if existing, found := p.Adventurer.Skills[record.Name]; found {
				existing.Level++
			} else {
				p.Adventurer.Skills[record.Name] = &models.AdventurerSkill{
					RecordKey: record.Name,
					Level:     2,
				}
			}
		}

		// depending on the chosen stock, we handle different steps from here
		stepChooseStock, found := p.GetStepByKey(keyChooseStock)
		if !found {
			return
		}

		strStock := strings.ToLower(stepChooseStock.Answer)

		if strings.Contains(strStock, "elf") || strings.Contains(strStock, "elven") {
			p.PushStep(p.chooseElvenWise())
		} else if strings.Contains(strStock, "dwarf") || strings.Contains(strStock, "dwarven") {
			p.PushStep(p.chooseDwarvenWise())
		} else if strings.Contains(strStock, "halfling") {
			p.PushStep(p.chooseHalflingWise())
		} else {
			p.PushStep(p.chooseHumanWise())
		}
	}

	return step
}

func (p *procedureCreateAdventurer) chooseElvenWise() *procedure.Step {
	choices := make([]procedure.Choice, 0)

	for _, recordName := range []string{
		"Elven Lore-wise",
		"Folly of Humanity-wise",
		"Folly of Dwarves-wise",
	} {
		if record, err := p.service.records.GetWiseByName(recordName); err == nil {
			choices = append(choices, procedure.Choice{
				Name:        record.Name,
				Description: record.Description,
			})
		}
	}

	step := &procedure.Step{
		Key:     keyChooseElvenWise,
		Prompt:  "How is your adventurer wise?",
		Choices: choices,
	}

	step.OnComplete = func() {
		if record, err := p.service.records.GetWiseByName(step.Answer); err == nil {
			p.Adventurer.Wises[step.Answer] = &models.AdventurerWise{
				Record: *record,
			}
		}

		p.PushStep(p.chooseAdditionalWise())
	}

	return step
}

func (p *procedureCreateAdventurer) chooseDwarvenWise() *procedure.Step {
	choices := make([]procedure.Choice, 0)

	for _, recordName := range []string{
		"Dwarven Chronicles-wise",
		"Shrewd Appraisal-wise",
	} {
		if record, err := p.service.records.GetWiseByName(recordName); err == nil {
			choices = append(choices, procedure.Choice{
				Name:        record.Name,
				Description: record.Description,
			})
		}
	}

	step := &procedure.Step{
		Key:     keyChooseDwarvenWise,
		Prompt:  "How is your adventurer wise?",
		Choices: choices,
	}

	step.OnComplete = func() {
		if record, err := p.service.records.GetWiseByName(step.Answer); err == nil {
			p.Adventurer.Wises[step.Answer] = &models.AdventurerWise{
				Record: *record,
			}
		}

		p.PushStep(p.chooseAdditionalWise())
	}

	return step
}

func (p *procedureCreateAdventurer) chooseHalflingWise() *procedure.Step {
	choices := make([]procedure.Choice, 0)

	for _, recordName := range []string{
		"Home-wise",
		"Needs a Little Salt-wise",
	} {
		if record, err := p.service.records.GetWiseByName(recordName); err == nil {
			choices = append(choices, procedure.Choice{
				Name:        record.Name,
				Description: record.Description,
			})
		}
	}

	step := &procedure.Step{
		Key:     keyChooseHalflingWise,
		Prompt:  "How is your adventurer wise?",
		Choices: choices,
	}

	step.OnComplete = func() {
		if record, err := p.service.records.GetWiseByName(step.Answer); err == nil {
			p.Adventurer.Wises[step.Answer] = &models.AdventurerWise{
				Record: *record,
			}
		}

		p.PushStep(p.chooseAdditionalWise())
	}

	return step
}

func (p *procedureCreateAdventurer) chooseAdditionalWise() *procedure.Step {
	step := &procedure.Step{
		Key:            keyChooseAdditionalWise,
		Prompt:         "How else is your adventurer wise? Please type in the following format: Something-wise: description here.",
		ValidatorRegex: "([a-zA-Z ]+-wise): (.*)",
	}

	step.OnComplete = func() {
		r := regexp.MustCompile(step.ValidatorRegex)
		matches := r.FindStringSubmatch(step.Answer)

		if len(matches) >= 2 {
			p.Adventurer.Wises[step.Answer] = &models.AdventurerWise{
				Record: models.WiseRecord{
					Name:        matches[0],
					Description: matches[1],
				},
			}
		}

		stepChooseStock, found := p.GetStepByKey(keyChooseStock)
		if !found {
			return
		}

		strStock := strings.ToLower(stepChooseStock.Answer)

		if strings.Contains(strStock, "elf") || strings.Contains(strStock, "elven") {
			p.PushStep(p.chooseElvenNature1())
		} else if strings.Contains(strStock, "dwarf") || strings.Contains(strStock, "dwarven") {
			p.PushStep(p.chooseDwarvenNature1())
		} else if strings.Contains(strStock, "halfling") {
			p.PushStep(p.chooseHalflingNature1())
		}
	}

	return step
}

func (p *procedureCreateAdventurer) chooseHumanWise() *procedure.Step {
	step := &procedure.Step{
		Key:            keyChooseHumanWise,
		Prompt:         "How is your adventurer wise? Please type in the following format: {Something Specific}-wise: {description here}",
		ValidatorRegex: "([a-zA-Z ]+-wise): (.*)",
	}

	step.OnComplete = func() {
		r := regexp.MustCompile(step.ValidatorRegex)
		matches := r.FindStringSubmatch(step.Answer)

		if len(matches) >= 2 {
			p.Adventurer.Wises[step.Answer] = &models.AdventurerWise{
				Record: models.WiseRecord{
					Name:        matches[0],
					Description: matches[1],
				},
			}
		}

		p.PushStep(p.chooseHumanNature1())
	}

	return step
}

func (p *procedureCreateAdventurer) chooseElvenNature1() *procedure.Step {
	step := &procedure.Step{
		Key:    keyChooseElvenNature1,
		Prompt: "Have you learned the songs of creation and do you sing them to mend hearts and calm storms? Or do you focus your ancient will into crafting works of unparalleled beauty?",
		Choices: []procedure.Choice{
			{Name: "Sing", Description: "If you sing the ancient songs, increase your Nature by one."},
			{Name: "Craft", Description: "If you bend your will to crafting Elven artifacts, replace your Singing Nature descriptor with Enchanting."},
		},
	}

	step.OnComplete = func() {
		switch step.Answer {
		case "Sing":
			p.Adventurer.Abilities.Raw.Nature++
		case "Craft":
			for idx, descriptor := range p.Adventurer.Abilities.Raw.NatureDescriptors {
				if descriptor != "Singing" {
					continue
				}

				p.Adventurer.Abilities.Raw.NatureDescriptors = append(p.Adventurer.Abilities.Raw.NatureDescriptors[:idx], p.Adventurer.Abilities.Raw.NatureDescriptors[idx+1:]...)
				break
			}

			p.Adventurer.Abilities.Raw.NatureDescriptors = append(p.Adventurer.Abilities.Raw.NatureDescriptors, "Enchanting")
		}

		p.PushStep(p.chooseElvenNature2())
	}

	return step
}

func (p *procedureCreateAdventurer) chooseElvenNature2() *procedure.Step {
	step := &procedure.Step{
		Key:    keyChooseElvenNature2,
		Prompt: "When evil stalks the world, do you confront it? Or do you retreat to the hidden places of the elves and allow time to defeat your enemies?",
		Choices: []procedure.Choice{
			{Name: "Confront", Description: "If you confront evil, increase your First Born trait to level 2."},
			{Name: "Retreat", Description: "If you retreat and hide, increase your Nature by one."},
		},
	}

	step.OnComplete = func() {
		switch step.Answer {
		case "Confront":
			p.Adventurer.Abilities.Raw.Nature++
		case "Retreat":
			for idx, descriptor := range p.Adventurer.Abilities.Raw.NatureDescriptors {
				if descriptor != "Singing" {
					continue
				}

				p.Adventurer.Abilities.Raw.NatureDescriptors = append(p.Adventurer.Abilities.Raw.NatureDescriptors[:idx], p.Adventurer.Abilities.Raw.NatureDescriptors[idx+1:]...)
				break
			}

			p.Adventurer.Abilities.Raw.NatureDescriptors = append(p.Adventurer.Abilities.Raw.NatureDescriptors, "Enchanting")
		}

		p.PushStep(p.chooseElvenNature3())
	}

	return step
}

func (p *procedureCreateAdventurer) chooseElvenNature3() *procedure.Step {
	step := &procedure.Step{
		Key:    keyChooseElvenNature3,
		Prompt: "Do you yearn to follow the gulls to the sea and journey west beyond all knowledge? Or are you prepared to live a life of struggle and grief?",
		Choices: []procedure.Choice{
			{Name: "Journey", Description: "If you yearn to journey west, increase your Nature by one."},
			{Name: "Struggle", Description: "If you are prepared to live a life of struggle, you may replace your home trait with Fiery, Curious or Restless. If you have one of these traits already, increase it by one."},
		},
	}

	step.OnComplete = func() {
		switch step.Answer {
		case "Journey":
			p.Adventurer.Abilities.Raw.Nature++
			p.PushStep(p.chooseRelationships1())
		case "Struggle":
			p.PushStep(p.chooseElvenNature3Option())
		}
	}

	return step
}

func (p *procedureCreateAdventurer) chooseElvenNature3Option() *procedure.Step {
	stepHometownTrait, found := p.GetStepByKey(keyChooseHometownTrait)
	if !found {
		return nil
	}

	choices := make([]procedure.Choice, 0)
	for _, traitName := range []string{
		"Fiery",
		"Curious",
		"Restless",
	} {
		record, err := p.service.records.GetTraitByName(traitName)
		if err != nil {
			p.service.logger.Fatal().Msgf("cant find trait %q", traitName)
		}

		choices = append(choices, procedure.Choice{
			Name:        record.Name,
			Description: record.Description,
		})
	}

	choices = append(choices, procedure.Choice{
		Name:        "Don't change",
		Description: fmt.Sprintf("keep existing trait %q", stepHometownTrait.Answer),
	})

	step := &procedure.Step{
		Key:     keyChooseElvenNature3Option,
		Prompt:  "Would you like to replace your home trait with Fiery, Curious or Restless?",
		Choices: choices,
	}

	step.OnComplete = func() {
		switch step.Answer {
		case "Don't change":
			// noop
		default:
			record, err := p.service.records.GetTraitByName(step.Answer)
			if err != nil {
				p.service.logger.Fatal().Msgf("cant find trait %q", step.Answer)
			}

			delete(p.Adventurer.Traits, stepHometownTrait.Answer)

			p.Adventurer.Traits[record.Name] = &models.AdventurerTrait{
				RecordKey: record.Name,
				Level:     1,
			}
		}

		p.PushStep(p.chooseRelationships1())
	}

	return step
}

func (p *procedureCreateAdventurer) chooseDwarvenNature1() *procedure.Step {
	step := &procedure.Step{
		Key:    keyChooseDwarvenNature1,
		Prompt: "When your kin are slain and their halls plundered, will you spend your blood avenging them? Or will you demand a blood price from the kin slayers and council your people to let sleeping dragons lie?",
		Choices: []procedure.Choice{
			{Name: "Avenge", Description: "You would take revenge at any cost, increase Nature by one."},
			{Name: "Council", Description: "You would council your people to resist their blood lust, replace the Avenging Grudges descriptor with Negotiating."},
		},
	}

	step.OnComplete = func() {
		switch step.Answer {
		case "Avenge":
			p.Adventurer.Abilities.Raw.Nature++
		case "Council":
			for idx, descriptor := range p.Adventurer.Abilities.Raw.NatureDescriptors {
				if descriptor == "Avenging Grudges" {
					p.Adventurer.Abilities.Raw.NatureDescriptors[idx] = "Negotiating"
				}
			}
		}

		p.PushStep(p.chooseDwarvenNature2())
	}

	return step
}

func (p *procedureCreateAdventurer) chooseDwarvenNature2() *procedure.Step {
	step := &procedure.Step{
		Key:    keyChooseDwarvenNature2,
		Prompt: "Would you plunge ever deeper into the bones of the earth looking for treasures untold? Or do you fear what you would uncover should you dig too deep?",
		Choices: []procedure.Choice{
			{Name: "Delve", Description: "You dig ever deeper, increase your Nature by one."},
			{Name: "Fear", Description: "You fear what lies beneath, increase your Born of Earth and Stone trait to level 2."},
		},
	}

	step.OnComplete = func() {
		switch step.Answer {
		case "Delve":
			p.Adventurer.Abilities.Raw.Nature++
		case "Fear":
			if trait, found := p.Adventurer.Traits["Born of Earth and Stone"]; found {
				trait.Level = 2
			}
		}

		p.PushStep(p.chooseDwarvenNature3())
	}

	return step
}

func (p *procedureCreateAdventurer) chooseDwarvenNature3() *procedure.Step {
	step := &procedure.Step{
		Key:    keyChooseDwarvenNature3,
		Prompt: "Do you yearn to spend your days crafting wondrous objects from silver and gold? Or do you prefer to spend gold, preferably other people’s?",
		Choices: []procedure.Choice{
			{Name: "Craft", Description: "You were born to craft wondrous objects, increase your Nature by one."},
			{Name: "Spend", Description: "If you yearn to spend gold, set your starting Resources to 1."},
		},
	}

	step.OnComplete = func() {
		switch step.Answer {
		case "Craft":
			p.Adventurer.Abilities.Raw.Nature++
		case "Spend":
			p.Adventurer.Abilities.Town.Resources = 1
		}

		p.PushStep(p.chooseRelationships1())
	}

	return step
}

func (p *procedureCreateAdventurer) chooseHalflingNature1() *procedure.Step {
	step := &procedure.Step{
		Key:    keyChooseHalflingNature1,
		Prompt: "Do you make the most out of every meal, slathering it with butter, lavishing it with syrup, worshipping it with wine? Or do you tighten your belt, shoo away guests and make fast the locks at night?",
		Choices: []procedure.Choice{
			{Name: "Indulge", Description: "If you make the most out of each meal, increase your Nature by one."},
			{Name: "Tighten", Description: "If you tighten your belt with a grim face, replace your Merrymaking descriptor with Hoarding."},
		},
	}

	step.OnComplete = func() {
		switch step.Answer {
		case "Indulge":
			p.Adventurer.Abilities.Raw.Nature++
		case "Tighten":
			for idx, descriptor := range p.Adventurer.Abilities.Raw.NatureDescriptors {
				if descriptor == "Merrymaking" {
					p.Adventurer.Abilities.Raw.NatureDescriptors[idx] = "Hoarding"
				}
			}
		}

		p.PushStep(p.chooseHalflingNature2())
	}

	return step
}

func (p *procedureCreateAdventurer) chooseHalflingNature2() *procedure.Step {
	step := &procedure.Step{
		Key:    keyChooseHalflingNature2,
		Prompt: "When confronted by bullying big folk, do you put them in their place with a witty riddle? Or do you roll up your sleeves and show them you’re ready to teach them a lesson?",
		Choices: []procedure.Choice{
			{Name: "Riddle", Description: "If you offer up a clever riddle, increase your Nature by one."},
			{Name: "Teach", Description: "If you roll up your sleeves, increase your Hidden Depths trait to level 2."},
		},
	}

	step.OnComplete = func() {
		switch step.Answer {
		case "Riddle":
			p.Adventurer.Abilities.Raw.Nature++
		case "Teach":
			if trait, found := p.Adventurer.Traits["Hidden Depths"]; found {
				trait.Level = 2
			}
		}

		p.PushStep(p.chooseHalflingNature3())
	}

	return step
}

func (p *procedureCreateAdventurer) chooseHalflingNature3() *procedure.Step {
	step := &procedure.Step{
		Key:    keyChooseHalflingNature3,
		Prompt: "Do you sneak into dragons’ lairs just to see what all the fuss is about? Or do you prefer to announce your intentions and have a frank conversation about your concerns?",
		Choices: []procedure.Choice{
			{Name: "Sneak", Description: "If you sneak into dragons’ lairs, increase your Nature by one."},
			{Name: "Announce", Description: "If you announce your intentions to have a frank discussion, replace your Sneaking Nature descriptor with Demanding."},
		},
	}

	step.OnComplete = func() {
		switch step.Answer {
		case "Sneak":
			p.Adventurer.Abilities.Raw.Nature++
		case "Announce":
			for idx, descriptor := range p.Adventurer.Abilities.Raw.NatureDescriptors {
				if descriptor == "Sneaking" {
					p.Adventurer.Abilities.Raw.NatureDescriptors[idx] = "Demanding"
				}
			}
		}

		p.PushStep(p.chooseRelationships1())
	}

	return step
}

func (p *procedureCreateAdventurer) chooseHumanNature1() *procedure.Step {
	step := &procedure.Step{
		Key:    keyChooseHumanNature1,
		Prompt: "Do you sit by the hearth at night drinking and boasting of your great deeds? Or do you spend those chill nights quietly preparing for the dark times to come?",
		Choices: []procedure.Choice{
			{Name: "Boast", Description: "If you boast of your exploits, real or imagined, increase your Nature by one."},
			{Name: "Prepare", Description: "If you quietly prepare, increase your class trait to level 2."},
		},
	}

	step.OnComplete = func() {
		switch step.Answer {
		case "Boast":
			p.Adventurer.Abilities.Raw.Nature++
		case "Prepare":
			stepHumanUpbringing, found := p.GetStepByKey(keyChooseHumanUpbringing)
			if !found {
				return
			}

			if trait, found := p.Adventurer.Traits[stepHumanUpbringing.Answer]; found {
				trait.Level = 2
			}
		}

		p.PushStep(p.chooseHumanNature2())
	}

	return step
}

func (p *procedureCreateAdventurer) chooseHumanNature2() *procedure.Step {
	step := &procedure.Step{
		Key:    keyChooseHumanNature2,
		Prompt: "When the elves and dwarves voice their concerns, do you demand to be heard as an equal? Or do you bow your head and listen to the wisdom of your elders?",
		Choices: []procedure.Choice{
			{Name: "Demand", Description: "If you demand your rights, increase Nature by one."},
			{Name: "Listen", Description: "If you listen to the wisdom of the elder ones, take a second wise: Elf-wise, Dwarf-wise or Politics-wise."},
		},
	}

	step.OnComplete = func() {
		switch step.Answer {
		case "Demand":
			p.Adventurer.Abilities.Raw.Nature++
		case "Listen":
			stepHumanUpbringing, found := p.GetStepByKey(keyChooseHumanUpbringing)
			if !found {
				return
			}

			if trait, found := p.Adventurer.Traits[stepHumanUpbringing.Answer]; found {
				trait.Level = 2
			}
		}

		p.PushStep(p.chooseHumanNature3())
	}

	return step
}

func (p *procedureCreateAdventurer) chooseHumanNature2option() *procedure.Step {
	step := &procedure.Step{
		Key:    keyChooseHumanNature2option,
		Prompt: "Choose another wise:",
		Choices: []procedure.Choice{
			{Name: "Elf-wise", Description: "Elf-wise is a wise that represents a character's deep knowledge and understanding of the ways, customs, and culture of the elves. Characters with this wise have spent significant time among the elegant and enigmatic elves, learning their language, traditions, and history. This wise can be invaluable when dealing with elves, as it allows the character to navigate elven society, negotiate with elven leaders, or decipher elven artifacts and mysteries. It may also help the character gain insight into elven motivations, making it a valuable asset in situations involving elves."},
			{Name: "Dwarf-wise", Description: "Dwarf-wise reflects a character's expertise in all things related to dwarves, their culture, and their underground realms. Characters with this wise have delved deep into dwarven strongholds, learned their language, and forged bonds with these sturdy and proud people. This wise can prove essential when navigating the intricacies of dwarven politics, negotiating with dwarf merchants, or seeking the aid of dwarven craftsmen. It can also help the character identify valuable minerals, understand ancient dwarf runes, and navigate the treacherous tunnels and caverns of dwarven territories."},
			{Name: "Politics-wise", Description: "Politics-wise represents a character's knowledge and insight into the complex world of politics and power dynamics. Characters with this wise are well-versed in the art of diplomacy, intrigue, and governance. They understand the motivations of rulers, the workings of bureaucracies, and the intricacies of courtly etiquette. This wise can be an invaluable tool for characters involved in political maneuvering, whether they are negotiating treaties, vying for positions of authority, or trying to uncover hidden agendas. It can also help the character identify potential allies and adversaries in the political arena, making it a crucial asset in navigating the challenges of a politically-driven campaign."},
		},
	}

	step.OnComplete = func() {
		for _, choice := range step.Choices {
			if choice.Name != step.Answer {
				continue
			}

			p.Adventurer.Wises[choice.Name] = &models.AdventurerWise{
				Record: models.WiseRecord{
					Name:        choice.Name,
					Description: choice.Description,
				},
			}
		}

		p.PushStep(p.chooseHumanNature3())
	}

	return step
}

func (p *procedureCreateAdventurer) chooseHumanNature3() *procedure.Step {
	step := &procedure.Step{
		Key:    keyChooseHumanNature3,
		Prompt: "Would you flee from the hordes of goblins, beasts, and monsters who prey on civilization? Or will you plunge into their midst, questing for glory?",
		Choices: []procedure.Choice{
			{Name: "Flee", Description: "If you would flee and hide inside the walls of tall citadels, increase your Nature by one."},
			{Name: "Plunge", Description: "If you do not fear those who prey on civilization, you may replace your home trait with Loner, Foolhardy or Defender. If you have one of these traits already, increase it by one."},
		},
	}

	step.OnComplete = func() {
		switch step.Answer {
		case "Flee":
			p.Adventurer.Abilities.Raw.Nature++
			p.PushStep(p.chooseRelationships1())
		case "Plunge":
			p.PushStep(p.chooseHumanNature3Option())
		}
	}

	return step
}

func (p *procedureCreateAdventurer) chooseHumanNature3Option() *procedure.Step {
	stepHometownTrait, found := p.GetStepByKey(keyChooseHometownTrait)
	if !found {
		return nil
	}

	choices := make([]procedure.Choice, 0)
	for _, traitName := range []string{
		"Loner",
		"Foolhardy",
		"Defender",
	} {
		record, err := p.service.records.GetTraitByName(traitName)
		if err != nil {
			p.service.logger.Fatal().Msgf("cant find trait %q", traitName)
		}

		choices = append(choices, procedure.Choice{
			Name:        record.Name,
			Description: record.Description,
		})
	}

	choices = append(choices, procedure.Choice{
		Name:        "Don't change",
		Description: fmt.Sprintf("keep existing trait %q", stepHometownTrait.Answer),
	})

	step := &procedure.Step{
		Key:     keyChooseHumanNature3Option,
		Prompt:  "Would you like to replace your home trait with Fiery, Curious or Restless?",
		Choices: choices,
	}

	step.OnComplete = func() {
		switch step.Answer {
		case "Don't change":
			// noop
		default:
			record, err := p.service.records.GetTraitByName(step.Answer)
			if err != nil {
				p.service.logger.Fatal().Msgf("cant find trait %q", step.Answer)
			}

			delete(p.Adventurer.Traits, stepHometownTrait.Answer)

			p.Adventurer.Traits[record.Name] = &models.AdventurerTrait{
				RecordKey: record.Name,
				Level:     1,
			}
		}

		p.PushStep(p.chooseRelationships1())
	}

	return step
}

func (p *procedureCreateAdventurer) chooseRelationships1() *procedure.Step {
	step := &procedure.Step{
		Key:    keyChooseRelationships1,
		Prompt: "Do you have friends who enjoy your occasional visits or are you a loner, tough and cool?",
		Choices: []procedure.Choice{
			{Name: "Friends", Description: "If you have a friend, add +1 Circles. Some friends will help on the road or in the wild; others will help in towns. See the Starting Friend rules."},
			{Name: "Loner", Description: "If you are a loner, tough and cool, your Circles starts at 1, and you have an enemy. Write down the name of your nemesis or mortal enemy on your character sheet and see the Starting Enemy rules. Skip the rest of the Circles and Relationships questions and take the Loner trait at level 1 or increase it by one if you already have it. Also, go get snacks for the rest of the group while they finish answering the Circles questions."},
		},
	}

	step.OnComplete = func() {
		switch step.Answer {
		case "Friends":
			p.Adventurer.Abilities.Town.Circles++
			p.PushStep(p.chooseRelationshipFriend1())
		case "Loner":
			p.Adventurer.Abilities.Town.Circles = 1
			p.PushStep(p.chooseRelationshipEnemy())
		}
	}

	return step
}

func (p *procedureCreateAdventurer) chooseRelationshipFriend1() *procedure.Step {
	step := &procedure.Step{
		Key:             keyChooseRelationshipFriend1,
		Prompt:          "What is your friend's name?",
		ValidatorRegex:  `^\w+( \w+)?$`,
		ValidatorPrompt: "first name is required, last name is optional",
	}

	step.OnComplete = func() {
		p.PushStep(p.chooseRelationshipFriend2())
	}

	return step
}

func (p *procedureCreateAdventurer) chooseRelationshipFriend2() *procedure.Step {
	step := &procedure.Step{
		Key:    keyChooseRelationshipFriend2,
		Prompt: "Is your friend townsfolk or an adventurer?",
		Choices: []procedure.Choice{
			{Name: "Townsfolk", Description: "If your friend is town-bound, choose in which settlement they live and choose a profession from your hometown’s skill list for them."},
			{Name: "Adventurer", Description: "If an adventurer, choose their class and specialty for them. Their level is equal to yours, leveling up as you do. Determine the last place you saw your friend."},
		},
	}

	step.OnComplete = func() {
		switch step.Answer {
		case "Townsfolk":
			p.PushStep(p.chooseRelationshipFriendTownsfolk1())
		case "Adventurer":
			p.PushStep(p.chooseRelationshipFriendAdventurer())
		}
	}

	return step
}

func (p *procedureCreateAdventurer) chooseRelationshipFriendTownsfolk1() *procedure.Step {
	stepChooseWorld, found := p.GetStepByKey(keyChooseWorld)
	if !found {
		return nil
	}

	w, err := p.service.worlds.GetWorldByName(stepChooseWorld.Answer)
	if err != nil {
		p.service.logger.Fatal().Msgf("Creating adventurer, Step %d: %v", p.StepIndex(), err)
	}

	var choices []procedure.Choice

	for _, settlement := range w.Settlements {
		choices = append(choices, procedure.Choice{
			Name:        settlement.Name,
			Description: settlement.Description(),
		})
	}

	step := &procedure.Step{
		Key:     keyChooseRelationshipFriendTownsfolk1,
		Prompt:  "Which settlement does your friend currently reside?",
		Choices: choices,
	}

	step.OnComplete = func() {
		p.PushStep(p.chooseRelationshipFriendTownsfolk2())
	}

	return step
}

func (p *procedureCreateAdventurer) chooseRelationshipFriendTownsfolk2() *procedure.Step {
	step := &procedure.Step{
		Key:    keyChooseRelationshipFriendTownsfolk2,
		Prompt: "Decide if your friend is townsfolk or an adventurer. Write your friend’s name on your character sheet.",
		Choices: []procedure.Choice{
			{Name: "Townsfolk", Description: "If your friend is town-bound, choose in which settlement they live and choose a profession from your hometown’s skill list for them."},
			{Name: "Adventurer", Description: "If an adventurer, choose their class and specialty for them. Their level is equal to yours, leveling up as you do. Determine the last place you saw your friend."},
		},
	}

	step.OnComplete = func() {
		// You can define the next step or action here.
	}

	return step
}

func (p *procedureCreateAdventurer) chooseRelationshipFriendAdventurer() *procedure.Step {
	step := &procedure.Step{
		Key:    keyChooseRelationshipFriendAdventurer,
		Prompt: "Decide if your friend is townsfolk or an adventurer. Write your friend’s name on your character sheet.",
		Choices: []procedure.Choice{
			{Name: "Townsfolk", Description: "If your friend is town-bound, choose in which settlement they live and choose a profession from your hometown’s skill list for them."},
			{Name: "Adventurer", Description: "If an adventurer, choose their class and specialty for them. Their level is equal to yours, leveling up as you do. Determine the last place you saw your friend."},
		},
	}

	step.OnComplete = func() {
		// You can define the next step or action here.
	}

	return step
}

func (p *procedureCreateAdventurer) chooseRelationshipFriendAdventurerLastSeen() *procedure.Step {
	step := &procedure.Step{
		Key:    keyChooseRelationshipFriendAdventurerLastSeen,
		Prompt: "Decide if your friend is townsfolk or an adventurer. Write your friend’s name on your character sheet.",
		Choices: []procedure.Choice{
			{Name: "Townsfolk", Description: "If your friend is town-bound, choose in which settlement they live and choose a profession from your hometown’s skill list for them."},
			{Name: "Adventurer", Description: "If an adventurer, choose their class and specialty for them. Their level is equal to yours, leveling up as you do. Determine the last place you saw your friend."},
		},
	}

	step.OnComplete = func() {
		// You can define the next step or action here.
	}

	return step
}

func (p *procedureCreateAdventurer) chooseRelationshipParents() *procedure.Step {
	step := &procedure.Step{
		Key:    keyChooseRelationshipParents,
		Prompt: "Do you have parents you can stomach talking to or are you an orphan?",
		Choices: []procedure.Choice{
			{Name: "Parents", Description: "If you have parents, add +1 Circles. Note your family name or parents’ names on your character sheet. Choose a trade for your parents from your hometown’s skill list."},
			{Name: "Orphan", Description: "If you’re an orphan, you have a keepsake from your parents that is worn around your neck or on one hand (worn/neck or worn/hand). Describe its sentimental value. It is worth 1D of treasure. Put it in your inventory."},
		},
	}

	step.OnComplete = func() {
		p.PushStep(p.chooseRelationshipMentor())
	}

	return step
}

func (p *procedureCreateAdventurer) chooseRelationshipMentor() *procedure.Step {
	step := &procedure.Step{
		Key:    keyChooseRelationshipMentor,
		Prompt: "Did you have a mentor or did you make your own way in this rough life?",
		Choices: []procedure.Choice{
			{Name: "Mentor", Description: "If you have a mentor, add +1 Circles. Your mentor is a 7th level adventurer of the same class. Note your mentor’s name on your character sheet. Magicians must select a mentor."},
			{Name: "Own Way", Description: "If you made your own way in life, you start with a pouch of gold coins worth 2D of treasure (belt 1). Put it in your inventory."},
		},
	}

	step.OnComplete = func() {
		p.PushStep(p.chooseRelationshipEnemy())
	}

	return step
}

func (p *procedureCreateAdventurer) chooseRelationshipEnemy() *procedure.Step {
	step := &procedure.Step{
		Key:    keyChooseRelationshipEnemy,
		Prompt: "Have you made an enemy in your life or have your dubious deeds managed to escape notice?",
		Choices: []procedure.Choice{
			{Name: "Made Enemy", Description: "If you have made an enemy, add +1 Circles. Note your enemy’s name and see the Starting Enemy rules."},
			{Name: "No Enemy", Description: "The benefit for not having an enemy is not having an enemy."},
		},
	}

	step.OnComplete = func() {
		// You can define the next step or action here.
	}

	return step
}
