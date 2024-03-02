package adventurer

import (
	"fmt"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"

	"torchbearer/pkg/models"
)

func (s *Service) OnDiscordInteractionCreate(session *discordgo.Session) {
	commands := []*discordgo.ApplicationCommand{
		{
			Name:        "tb",
			Description: "Torchbearer commands",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "adventurer",
					Description: "Adventurer management commands",
					Options: []*discordgo.ApplicationCommandOption{
						{
							Name:        "create",
							Description: "create a new adventurer",
							Type:        discordgo.ApplicationCommandOptionSubCommand,
						},
						{
							Name:        "list",
							Description: "list the adventurers",
							Type:        discordgo.ApplicationCommandOptionSubCommand,
						},
						{
							Name:        "view",
							Description: "view an adventurer by name",
							Type:        discordgo.ApplicationCommandOptionSubCommand,
							Options: []*discordgo.ApplicationCommandOption{
								{
									Name:        "name",
									Description: "the name of the adventurer you want to view",
									Type:        discordgo.ApplicationCommandOptionString,
								},
							},
						},
						{
							Name:        "delete",
							Description: "delete an adventurer by name",
							Type:        discordgo.ApplicationCommandOptionSubCommand,
						},
					},
					Type: discordgo.ApplicationCommandOptionSubCommandGroup,
				},
				//{
				//	Name:        "subcommand",
				//	Description: "Top-level subcommand",
				//	Type:        discordgo.ApplicationCommandOptionSubCommand,
				//},
			},
		},
	}

	commandHandlers := map[string]func(session *discordgo.Session, i *discordgo.InteractionCreate){
		"tb": s.tbCommandHandler,
	}

	registeredCommands := make([]*discordgo.ApplicationCommand, len(commands))

	for i, v := range commands {
		s.logger.Info("adding command %q", v.Name)
		cmd, err := session.ApplicationCommandCreate(session.State.User.ID, "", v)
		if err != nil {
			s.logger.Error("cannot create command", "name", v.Name, "error", err)
		}

		registeredCommands[i] = cmd
	}

	session.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})
}

func (s *Service) tbCommandHandler(session *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options
	content := ""

	switch options[0].Name {
	case "adventurer":
		options = options[0].Options // options of subcommand
		switch options[0].Name {
		case "create":
			s.createAdventurerModal(session, i)
		case "list":
			names := make([]string, 0)
			for _, a := range s.adventurers {
				names = append(names, fmt.Sprintf("* %s", a.Name))
			}
			content = "\n" + strings.Join(names, "\n")
		case "view":
			options = options[0].Options // options of nested subcommand
			name := fmt.Sprintf("%s", options[0].Value)

			a, err := s.GetAdventurerByName(name)
			if err != nil {
				content = fmt.Sprintf("getting adventurer by name %q: %v", name, err)
			} else {
				content = "\n" + s.adventurerMarkdownDescription(a)
			}
		case "delete":
			content = "view"
		}
	}

	session.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: content,
		},
	})
}

func (s *Service) createAdventurerModal(session *discordgo.Session, i *discordgo.InteractionCreate) {
	p := s.CreateAdventurerProcedure()

	var isFollowup bool
	var currentStepIndex int

	for {
		// terminate when no more steps
		step := p.NextStep()
		if step == nil {
			if p.OnComplete != nil {
				// call procedure completion callback if done
				p.OnComplete()
			}

			break
		}

		// procedure is either text-input of choice-selection prompt
		var component discordgo.MessageComponent
		if len(step.Choices) < 1 {
			component = &discordgo.TextInput{
				CustomID:    step.Key,
				Label:       step.Prompt,
				Style:       discordgo.TextInputShort,
				Placeholder: step.Default,
				Required:    true,
				MaxLength:   300,
				MinLength:   10,
			}
		} else {
			menu := &discordgo.SelectMenu{
				MenuType:     discordgo.StringSelectMenu,
				CustomID:     step.Key,
				ChannelTypes: []discordgo.ChannelType{discordgo.ChannelTypeDM},
				Placeholder:  step.Prompt,
			}
			component = menu

			// build up our choices from the step
			for _, choice := range step.Choices {
				menu.Options = append(menu.Options, discordgo.SelectMenuOption{
					Label:       choice.Name,
					Value:       choice.Name,
					Description: choice.Description,
				})
			}
		}

		response := &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				CustomID: "modals_survey_" + i.Interaction.Member.User.ID,
				Title:    step.Procedure.Name,
				Components: []discordgo.MessageComponent{
					discordgo.ActionsRow{
						Components: []discordgo.MessageComponent{component},
					},
				},
			},
		}

		if isFollowup {
			for p.StepIndex() == currentStepIndex {
				time.Sleep(time.Millisecond * 10)
			}

			currentStepIndex = p.CurrentStepIndex()

			component = discordgo.ActionsRow{Components: []discordgo.MessageComponent{component}}

			message, err := session.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{
				Content:    step.Prompt,
				Components: []discordgo.MessageComponent{component},
				Flags:      discordgo.MessageFlagsEphemeral,
			})

			if err != nil {
				s.logger.Error("creating FollowupMessageCreate", "error", err)
				return
			}

			_ = message
		} else {
			isFollowup = true

			if err := session.InteractionRespond(i.Interaction, response); err != nil {
				s.logger.Error("creating InterationRespond", "error", err)
				return
			}
		}
	}
}

func (s *Service) adventurerMarkdownDescription(a *models.Adventurer) string {
	const strNone = "None"

	lines := make([]string, 0)

	{
		content := fmt.Sprintf("* **Name**: %s", a.Name)
		lines = append(lines, content)
	}

	{
		content := fmt.Sprintf("* **Stock**: %s", a.Stock.Name)
		lines = append(lines, content)
	}

	{
		content := fmt.Sprintf("* **Conditions**: %s", a.Condition.String())
		lines = append(lines, content)
	}

	{
		traitNames := make([]string, 0)

		for _, trait := range a.Traits {
			traitNames = append(traitNames, fmt.Sprintf("%s (%d)", trait.RecordKey, trait.Level))
		}

		if len(traitNames) < 1 {
			traitNames = append(traitNames, strNone)
		}

		names := strings.Join(traitNames, ", ")
		content := fmt.Sprintf("* **Traits**: %s", names)
		lines = append(lines, content)
	}

	{
		skillNames := make([]string, 0)

		for _, skill := range a.Skills {
			skillNames = append(skillNames, fmt.Sprintf("%s (%d)", skill.RecordKey, skill.Level))
		}

		if len(skillNames) < 1 {
			skillNames = append(skillNames, strNone)
		}

		names := strings.Join(skillNames, ", ")
		content := fmt.Sprintf("* **Skills**: %s", names)
		lines = append(lines, content)
	}

	{
		var name string

		if w, err := s.worlds.GetWorldByID(a.World); err == nil {
			if friend, err := w.GetAdventurerByID(a.Relationships.Friend); err == nil {
				name = friend.Name
			}

			if friend, err := w.GetTownsfolkByID(a.Relationships.Friend); err == nil {
				name = friend.Name
			}
		}

		if name == "" {
			name = strNone
		}

		content := fmt.Sprintf("* **Friend**: %s", name)
		lines = append(lines, content)
	}

	{
		content := fmt.Sprintf("* **Raiment**: %s", strings.Join(a.Raiment, ", "))
		lines = append(lines, content)
	}

	{
		invLines := make([]string, 0)

		// WORN
		for slot, item := range map[string]*models.Item{
			"Head":       a.Inventory.Worn.Head,
			"Neck":       a.Inventory.Worn.Neck,
			"Left Hand":  a.Inventory.Worn.HandLeft,
			"Right Hand": a.Inventory.Worn.HandRight,
			"Torso1":     a.Inventory.Worn.Torso1,
			"Torso2":     a.Inventory.Worn.Torso2,
			"Torso3":     a.Inventory.Worn.Torso3,
			"Belt1":      a.Inventory.Worn.Belt1,
			"Belt2":      a.Inventory.Worn.Belt2,
			"Belt3":      a.Inventory.Worn.Belt3,
			"Legs":       a.Inventory.Worn.Legs,
			"Feet":       a.Inventory.Worn.Feet,
		} {
			if item == nil {
				continue
			}

			invLines = append(invLines, fmt.Sprintf("\t* %s (Worn, %s)", item.Name, slot))
		}

		// CARRIED
		for slot, item := range map[string]*models.Item{
			"Left Hand":  a.Inventory.Carried.HandLeft,
			"Right Hand": a.Inventory.Carried.HandRight,
		} {
			if item == nil {
				continue
			}

			invLines = append(invLines, fmt.Sprintf("\t* %s (carried, %s)", item.Name, slot))
		}

		var content string

		if len(invLines) < 1 {
			content = strNone
		} else {
			content = "\n" + strings.Join(invLines, "\n")
		}

		lines = append(lines, fmt.Sprintf("* **Inventory**: %s", content))
	}

	return strings.Join(lines, "\n")
}
