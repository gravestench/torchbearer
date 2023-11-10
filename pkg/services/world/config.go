package world

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"

	"torchbearer/pkg/models"
	"torchbearer/pkg/services/config"
)

func (s *Service) ConfigFileName() string {
	return "world.json"
}

func (s *Service) LoadConfig(config *config.Config) {
	if s.Worlds == nil {
		s.Worlds = make([]*World, 0)
	}

	for _, worldUUID := range config.GroupKeys() {
		g := config.Group(worldUUID)

		id, err := uuid.Parse(worldUUID)
		if err != nil {
			s.logger.Error().Msgf("parsing uuid: %v", err)
			continue
		}

		settlementsData := g.GetJson(keyWorldSettlements)
		settlements := make([]models.Settlement, 0)

		if err = json.Unmarshal(settlementsData, &settlements); err != nil {
			s.logger.Error().Msgf("unmarshalling settlement data: %v", err)
			continue
		}

		w := World{
			UUID:        id,
			Name:        g.GetString(keyWorldName),
			Settlements: settlements,
		}

		s.AddWorld(w)
	}
}

func (s *Service) LoadWorlds() error {
	cfg, err := s.cfgManager.GetConfigByFileName(s.ConfigFileName())
	if err != nil {
		return fmt.Errorf("getting world config: %v", err)
	}

	s.LoadConfig(cfg)

	return nil
}

func (s *Service) SaveWorlds() error {
	cfg, err := s.cfgManager.GetConfigByFileName(s.ConfigFileName())
	if err != nil {
		return fmt.Errorf("getting world config: %v", err)
	}

	for _, world := range s.Worlds {
		g := cfg.Group(world.UUID.String())
		g.Set(keyWorldName, world.Name)
		g.Set(keyWorldSettlements, world.Settlements)
		g.Set(keyWorldStats, world.Stats)
	}

	err = s.cfgManager.SaveConfigWithFileName(s.ConfigFileName())
	if err != nil {
		return fmt.Errorf("saving world config: %v", err)
	}

	return nil
}
