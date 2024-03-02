package records

import (
	"encoding/json"

	"torchbearer/pkg/models"
)

func (s *Service) initConfigFiles() {
	s.initSkillRecordsConfig()
	s.initStockRecordConfig()
	s.initTraitRecordConfig()
	s.initWisesRecordConfig()
}

func (s *Service) ConfigFileNames() []string {
	return []string{
		cfgFileNameSkillRecords,
		cfgFileNameStockRecords,
		cfgFileNameTraitRecords,
		cfgFileNameWiseRecords,
	}
}

const (
	cfgFileNameSkillRecords = "records.skills.json"

	keySkillName          = "Name"
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
			s.logger.Error("creating skill records config file", "error", err)
			panic(err)
		}
	}

	defaults := s.defaultSkillTable()

	for _, skill := range defaults {
		g := cfg.Group(skill.Name)

		g.Set(keySkillName, skill.Name)
		g.Set(keySkillDescription, skill.Description)
		g.Set(keySkillBeginnersLuck, skill.BeginnersLuck)
		g.Set(keySkillTools, skill.Tools)
		g.Set(keySkillConsumables, skill.Consumables)
		g.Set(keySkillSupportSkills, skill.SupportSkills)
		g.Set(keySkillFactors, skill.Factors)
	}

	// save the config (ensure to overwrite changes to the defaults)
	if err = s.cfgManager.SaveConfigWithFileName(cfgFileNameSkillRecords); err != nil {
		s.logger.Error("saving config", "filename", cfgFileNameSkillRecords, "error", err)
		panic(err)
	}

	// iterate through the config, with defaults and any custom stuff
	for _, name := range cfg.GroupKeys() {
		var skill models.SkillRecord

		g := cfg.Group(name)

		skill.Name = name
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
			s.logger.Error("creating stock records config file", "error", err)
			panic(err)
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
		s.logger.Error("saving config", "filename", cfgFileNameStockRecords, "error", err)
		panic(err)
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
			s.logger.Error("creating stock records config file", "error", err)
			panic(err)
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
		s.logger.Error("saving config", "filename", cfgFileNameTraitRecords, "error", err)
		panic(err)
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

const (
	cfgFileNameWiseRecords = "records.wise.json"
	keyWiseDescription     = "Description"
)

func (s *Service) initWisesRecordConfig() {
	// load or create the config
	cfg, err := s.cfgManager.GetConfigByFileName(cfgFileNameWiseRecords)
	if err != nil {
		cfg, err = s.cfgManager.CreateConfigWithFileName(cfgFileNameWiseRecords)
		if err != nil {
			s.logger.Error("creating stock records config file", "error", err)
			panic(err)
		}
	}

	// ensure the defaults are used
	defaults := s.initWisesTable()
	for _, wise := range defaults {
		g := cfg.Group(wise.Name)
		g.Set(keyWiseDescription, wise.Description)
	}

	// save the config (ensure to overwrite changes to the defaults)
	if err = s.cfgManager.SaveConfigWithFileName(cfgFileNameWiseRecords); err != nil {
		s.logger.Error("saving config", "filename", cfgFileNameWiseRecords, "error", err)
		panic(err)
	}

	// iterate through the config, with defaults and any custom stuff
	for _, name := range cfg.GroupKeys() {
		var wise models.WiseRecord

		g := cfg.Group(name)

		wise.Name = name
		wise.Description = g.GetString(keyWiseDescription)

		s.WisesRecords[name] = wise
	}
}
