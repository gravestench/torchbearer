package world

import (
	"fmt"

	"github.com/google/uuid"

	"torchbearer/pkg/models"
)

const (
	keyWorldName        = "Name"
	keyWorldSettlements = "Settlements"
	keyWorldStats       = "Stats"
)

type World struct {
	uuid.UUID
	*Service
	Name        string
	Settlements []models.Settlement
	Stats       struct {
		SessionsPlayed int
		TestsRolled    int
	}
}

func (s *Service) NewWorld(name string) (*World, error) {
	if existing, _ := s.GetWorld(name); existing != nil {
		return nil, fmt.Errorf("world with name %q already exists", name)
	}

	w := &World{
		Service: s,
		UUID:    uuid.New(),
		Name:    name,
	}

	s.generateNewWorldSettlements(w)

	return w, nil
}

func (s *Service) AddWorld(w World) {
	var exists bool

	for _, existing := range s.Worlds {
		if existing.UUID == w.UUID {
			exists = true
			break
		}
	}

	if exists {
		return
	}

	s.Worlds = append(s.Worlds, &w)
}

func (s *Service) GetWorld(name string) (*World, error) {
	for _, world := range s.Worlds {
		if world.Name != name {
			continue
		}

		return world, nil
	}

	return nil, fmt.Errorf("world %q not found", name)
}

func (s *Service) DeleteWorld(name string) error {
	for idx, world := range s.Worlds {
		if world.Name != name {
			continue
		}

		s.Worlds = append(s.Worlds[:idx], s.Worlds[idx+1:]...)

		if cfg, err := s.cfgManager.GetConfigByFileName(s.ConfigFileName()); err != nil {
			return fmt.Errorf("getting world config: %v", err)
		} else {
			cfg.Delete(world.UUID.String())
		}

		return nil
	}

	return fmt.Errorf("world %q not found", name)
}
