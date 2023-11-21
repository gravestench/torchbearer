package world

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/google/uuid"

	"torchbearer/pkg/models"
)

type World struct {
	UUID        uuid.UUID
	Name        string
	Seed        int64
	AsciiMap    string
	rng         *rand.Rand
	Settlements []*models.Settlement
	Stats       struct {
		SessionsPlayed int
		TestsRolled    int
	}
	*Service `json:"-"`
}

func (s *Service) NewWorld(name string) (*World, error) {
	if existing, _ := s.GetWorldByName(name); existing != nil {
		return nil, fmt.Errorf("world with name %q already exists", name)
	}

	w := &World{
		Service: s,
		UUID:    uuid.New(),
		Name:    name,
		Seed:    time.Now().UnixNano(),
	}

	w.rng = rand.New(rand.NewSource(w.Seed))

	w.generateNewWorldSettlements()
	w.generateAsciiMap()

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

func (s *Service) GetWorldByName(name string) (*World, error) {
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
			s.cfgManager.SaveConfigWithFileName(s.ConfigFileName())
		}

		return nil
	}

	return fmt.Errorf("world %q not found", name)
}
