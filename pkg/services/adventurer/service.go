package adventurer

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/gravestench/runtime"
	"github.com/rs/zerolog"

	"torchbearer/pkg/models"
	chatgpt_agent "torchbearer/pkg/services/chatgpt-agent"
	"torchbearer/pkg/services/config"
	"torchbearer/pkg/services/phase"
	"torchbearer/pkg/services/records"
	"torchbearer/pkg/services/world"
)

const (
	cfgFileNameAdventurers     = "adventurers.json"
	keyAdventurerWorld         = "World UUID"
	keyAdventurerHometown      = "Hometown UUID"
	keyAdventurerParty         = "Party UUID"
	keyAdventurerName          = "Name"
	keyAdventurerStock         = "Stock"
	keyAdventurerRaiment       = "Raiment"
	keyAdventurerCondition     = "Condition"
	keyAdventurerGoals         = "Goals"
	keyAdventurerRewards       = "Rewards"
	keyAdventurerRelationships = "Relationships"
	keyAdventurerAbilities     = "Abilities"
	keyAdventurerTraits        = "Traits"
	keyAdventurerSkills        = "Skills"
	keyAdventurerWises         = "Wises"
)

type Service struct {
	logger      *zerolog.Logger
	config      config.Dependency
	phase       phase.Dependency
	records     records.Dependency
	worlds      world.Dependency
	adventurers []*models.Adventurer
}

func (s *Service) Init(rt runtime.Runtime) {
	s.adventurers = make([]*models.Adventurer, 0)

	if err := s.LoadAdventurers(); err != nil {
		s.logger.Error().Msgf("loading adventurers: %v", err)
	}

	go func() {
		agent := &chatgpt_agent.Service{}
		rt.Add(agent).Wait()

		for !agent.DependenciesResolved() {
			time.Sleep(time.Second)
		}

		const initContext = "You are creating a character for the torchbearer RPG system. I will be asking you a series of questions which are either text input or choice selection. i want you to respond tersely and only respond with what is being asked. If you dont know best answer, just make a guess. Do not respond with anything except for an exact match to the validator regex."
		agent.SetContext(initContext)

		p := s.CreateAdventurerProcedure()

		for {
			step := p.NextStep()
			if step == nil {
				break
			}

			question := step.Prompt
			if step.Default != "" {
				question = fmt.Sprintf("%s. A suggested default is '%s'.", question, step.Default)
			}

			if step.ValidatorRegex != "" {
				question = fmt.Sprintf("%s. Valid responses will match the regex pattern '%s'.", question, step.ValidatorRegex)
			}

			if step.ValidatorPrompt != "" {
				question = fmt.Sprintf("%s. A hint for passing validation is '%s'.", question, step.ValidatorPrompt)
			}

			if len(step.Choices) > 0 {
				choices := make([]string, 0)
				for _, choice := range step.Choices {
					choices = append(choices, choice.Name)
				}

				question = fmt.Sprintf("%s Valid choices are [%s]", question, strings.Join(choices, ", "))
			}

			s.logger.Warn().Msgf("Question: %s", step.Prompt)

			res, err := agent.Ask(question)
			if err != nil {
				s.logger.Error().Msgf("asking GPT to answer a question: %v", err)
				continue
			}

			step.Answer = res

			if err := step.Complete(); err != nil {
				continue
			}

			s.logger.Warn().Msgf("Response: %s", res)
			s.logger.Warn().Msgf("Context: %s", agent.Context())
		}
	}()
}

func (s *Service) Name() string {
	return "Adventurer"
}

func (s *Service) BindLogger(logger *zerolog.Logger) {
	s.logger = logger
}

func (s *Service) Logger() *zerolog.Logger {
	return s.logger
}

func (s *Service) LoadAdventurers() error {
	cfg, err := s.config.GetConfigByFileName(cfgFileNameAdventurers)
	if err != nil {
		return fmt.Errorf("loading adventurers: %v", err)
	}

	for _, id := range cfg.GroupKeys() {
		g := cfg.Group(id)
		advID, err := uuid.Parse(id)
		if err != nil {
			continue
		}

		a := models.Adventurer{ID: advID}
		{
			data := g.GetJson(keyAdventurerWorld)
			json.Unmarshal(data, &a.World)
		}
		{
			data := g.GetJson(keyAdventurerHometown)
			json.Unmarshal(data, &a.Hometown)
		}
		{
			data := g.GetJson(keyAdventurerParty)
			json.Unmarshal(data, &a.Party)
		}
		{
			data := g.GetJson(keyAdventurerName)
			json.Unmarshal(data, &a.Name)
		}
		{
			data := g.GetJson(keyAdventurerStock)
			json.Unmarshal(data, &a.Stock)
		}
		{
			data := g.GetJson(keyAdventurerRaiment)
			json.Unmarshal(data, &a.Raiment)
		}
		{
			data := g.GetJson(keyAdventurerCondition)
			json.Unmarshal(data, &a.Condition)
		}
		{
			data := g.GetJson(keyAdventurerGoals)
			json.Unmarshal(data, &a.Goals)
		}
		{
			data := g.GetJson(keyAdventurerRewards)
			json.Unmarshal(data, &a.Rewards)
		}
		{
			data := g.GetJson(keyAdventurerRelationships)
			json.Unmarshal(data, &a.Relationships)
		}
		{
			data := g.GetJson(keyAdventurerAbilities)
			json.Unmarshal(data, &a.Abilities)
		}
		{
			data := g.GetJson(keyAdventurerTraits)
			json.Unmarshal(data, &a.Traits)
		}
		{
			data := g.GetJson(keyAdventurerSkills)
			json.Unmarshal(data, &a.Skills)
		}
		{
			data := g.GetJson(keyAdventurerWises)
			json.Unmarshal(data, &a.Wises)
		}

		s.adventurers = append(s.adventurers, &a)
	}

	return nil
}

func (s *Service) SaveAdventurers() error {
	cfg, err := s.config.GetConfigByFileName(cfgFileNameAdventurers)
	if err != nil {
		cfg, err = s.config.CreateConfigWithFileName(cfgFileNameAdventurers)
		if err != nil {
			s.logger.Fatal().Msgf("creating skill records config file: %v", err)
		}
	}

	for _, a := range s.adventurers {
		g := cfg.Group(a.ID.String())
		g.Set(keyAdventurerWorld, a.World)
		g.Set(keyAdventurerHometown, a.Hometown)
		g.Set(keyAdventurerParty, a.Party)
		g.Set(keyAdventurerName, a.Name)
		g.Set(keyAdventurerStock, a.Stock)
		g.Set(keyAdventurerRaiment, a.Raiment)
		g.Set(keyAdventurerCondition, a.Condition)
		g.Set(keyAdventurerGoals, a.Goals)
		g.Set(keyAdventurerRewards, a.Rewards)
		g.Set(keyAdventurerRelationships, a.Relationships)
		g.Set(keyAdventurerAbilities, a.Abilities)
		g.Set(keyAdventurerTraits, a.Traits)
		g.Set(keyAdventurerSkills, a.Skills)
		g.Set(keyAdventurerWises, a.Wises)
	}

	return s.config.SaveConfigWithFileName(cfgFileNameAdventurers)
}

func (s *Service) Adventurers() ([]*models.Adventurer, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *Service) NewAdventurer() *models.Adventurer {
	a := &models.Adventurer{
		ID:      uuid.New(),
		Raiment: make(models.Raiment, 0),
		Traits:  make(map[string]*models.AdventurerTrait),
		Skills:  make(map[string]*models.AdventurerSkill),
		Wises:   make(map[string]*models.AdventurerWise),
	}

	return a
}

func (s *Service) AddAdventurer(a *models.Adventurer) error {
	s.adventurers = append(s.adventurers, a)

	return s.SaveAdventurers()
}

func (s *Service) RemoveAdventurer(name string) error {
	return fmt.Errorf("not implemented")
}
