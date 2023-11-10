package records

import (
	"encoding/json"

	"torchbearer/pkg/models"
)

func (s *Service) initConfigFiles() {
	s.initSkillRecordsConfig()
	s.initStockRecordConfig()
	s.initTraitRecordConfig()
}

func (s *Service) ConfigFileNames() []string {
	return []string{cfgFileNameSkillRecords, cfgFileNameStockRecords, cfgFileNameTraitRecords}
}

const (
	cfgFileNameSkillRecords = "records.skills.json"

	keySkillDescription   = "Description"
	keySkillBeginnersLuck = "Beginners Luck Ability"
	keySkillTools         = "Tools"
	keySkillConsumables   = "Consumables"
	keySkillSupportSkills = "Support Skills"
	keySkillFactors       = "Factors"
)

func (s *Service) initSkillRecordsConfig() {
	// load or create the config
	cfg, err := s.cfgManager.GetConfigByFileName(cfgFileNameSkillRecords)
	if err != nil {
		cfg, err = s.cfgManager.CreateConfigWithFileName(cfgFileNameSkillRecords)
		if err != nil {
			s.logger.Fatal().Msgf("creating skill records config file: %v", err)
		}
	}

	defaults := s.defaultSkillTable()

	for _, skill := range defaults {
		g := cfg.Group(skill.Name)

		g.Set(keySkillDescription, skill.Description)
		g.Set(keySkillBeginnersLuck, skill.BeginnersLuck)
		g.Set(keySkillTools, skill.Tools)
		g.Set(keySkillConsumables, skill.Consumables)
		g.Set(keySkillSupportSkills, skill.SupportSkills)
		g.Set(keySkillFactors, skill.Factors)
	}

	// save the config (ensure to overwrite changes to the defaults)
	if err = s.cfgManager.SaveConfigWithFileName(cfgFileNameSkillRecords); err != nil {
		s.logger.Fatal().Msgf("saving config %q: %v", cfgFileNameSkillRecords, err)
	}

	// iterate through the config, with defaults and any custom stuff
	for _, name := range cfg.GroupKeys() {
		var skill models.SkillRecord

		g := cfg.Group(name)

		skill.Description = g.GetString(keySkillDescription)
		skill.BeginnersLuck = g.GetString(keySkillBeginnersLuck)
		skill.Tools = g.GetStrings(keySkillTools)
		skill.Consumables = g.GetStrings(keySkillConsumables)

		skill.SupportSkills = g.GetStrings(keySkillSupportSkills)
		skill.Factors = g.Get(keySkillFactors).(map[string][]models.SkillFactor)

		s.SkillRecords[name] = skill
	}
}

const (
	cfgFileNameStockRecords = "records.stocks.json"

	keyStockRace          = "Race"
	keyStockClass         = "Class"
	keyStockDescription   = "Description"
	keyStockLevelBenefits = "Level Benefits"
)

func (s *Service) initStockRecordConfig() {
	// load or create the config
	cfg, err := s.cfgManager.GetConfigByFileName(cfgFileNameStockRecords)
	if err != nil {
		cfg, err = s.cfgManager.CreateConfigWithFileName(cfgFileNameStockRecords)
		if err != nil {
			s.logger.Fatal().Msgf("creating stock records config file: %v", err)
		}
	}

	// ensure the defaults are used
	defaults := s.initStockTable()

	for _, stock := range defaults {
		g := cfg.Group(stock.String())

		g.Set(keyStockRace, stock.Race)
		g.Set(keyStockClass, stock.Class)
		g.Set(keyStockDescription, stock.Description)
		g.Set(keyStockLevelBenefits, stock.LevelBenefits)
	}

	// save the config (ensure to overwrite changes to the defaults)
	if err = s.cfgManager.SaveConfigWithFileName(cfgFileNameStockRecords); err != nil {
		s.logger.Fatal().Msgf("saving config %q: %v", cfgFileNameStockRecords, err)
	}

	// iterate through the config, with defaults and any custom stuff
	for _, key := range cfg.GroupKeys() {
		var stock models.Stock

		g := cfg.Group(key)

		stock.Race = g.Get(keyStockRace).(models.Race)
		stock.Class = g.Get(keyStockClass).(models.Class)
		stock.Description = g.GetString(keyStockDescription)

		benefitsJson := g.GetJson(keyStockLevelBenefits)
		_ = json.Unmarshal(benefitsJson, &stock.LevelBenefits)

		s.StockRecords[key] = stock
	}
}

const (
	cfgFileNameTraitRecords = "records.traits.json"
	keyTraitDescription     = "Description"
)

func (s *Service) initTraitRecordConfig() {
	// load or create the config
	cfg, err := s.cfgManager.GetConfigByFileName(cfgFileNameTraitRecords)
	if err != nil {
		cfg, err = s.cfgManager.CreateConfigWithFileName(cfgFileNameTraitRecords)
		if err != nil {
			s.logger.Fatal().Msgf("creating stock records config file: %v", err)
		}
	}

	// ensure the defaults are used
	defaults := s.initTraitTable()
	for _, trait := range defaults {
		g := cfg.Group(trait.Name)
		g.Set(keyTraitDescription, trait.Description)
	}

	// save the config (ensure to overwrite changes to the defaults)
	if err = s.cfgManager.SaveConfigWithFileName(cfgFileNameTraitRecords); err != nil {
		s.logger.Fatal().Msgf("saving config %q: %v", cfgFileNameTraitRecords, err)
	}

	// iterate through the config, with defaults and any custom stuff
	for _, name := range cfg.GroupKeys() {
		var trait models.TraitRecord

		g := cfg.Group(name)

		trait.Name = name
		trait.Description = g.GetString(keyTraitDescription)

		s.TraitRecords[name] = trait
	}
}
