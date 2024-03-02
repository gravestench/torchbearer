package world

import (
	"fmt"
	"log/slog"
	"math/rand"
	"sort"
	"time"

	"github.com/google/uuid"
	"github.com/gravestench/servicemesh"

	"torchbearer/pkg/services/config"
	"torchbearer/pkg/services/phase"
	"torchbearer/pkg/services/records"
)

type Service struct {
	logger     *slog.Logger
	phase      phase.PhaseManager
	cfgManager config.Dependency
	records    records.Dependency
	adventurer AdventurerManager
	Worlds     []*World
	ready      bool
}

func (s *Service) Init(mesh servicemesh.Mesh) {
	s.Worlds = make([]*World, 0)
	if err := s.LoadWorlds(); err != nil {
		s.logger.Error("loading world config", "error", err)
	}
	s.ready = true
}

func (s *Service) IsLoaded() bool {
	return s.ready
}

func (s *Service) Name() string {
	return "World"
}

func (s *Service) SetLogger(logger *slog.Logger) {
	s.logger = logger
}

func (s *Service) Logger() *slog.Logger {
	return s.logger
}

func (s *Service) GetSortedWorlds() []*World {
	worlds := make([]*World, 0)

	for _, world := range s.Worlds {
		worlds = append(worlds, world)
	}

	sort.Slice(worlds, func(i, j int) bool {
		return worlds[i].WorldID.String() < worlds[i].WorldID.String()
	})

	return worlds
}

func (s *Service) GetWorlds() []*World {
	return s.GetSortedWorlds()
}

func (s *Service) NewWorld(name string) (*World, error) {
	if existing, _ := s.GetWorldByName(name); existing != nil {
		return nil, fmt.Errorf("world with name %q already exists", name)
	}

	w := &World{
		Service: s,
		WorldID: uuid.New(),
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
		if existing.WorldID == w.WorldID {
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

func (s *Service) GetWorldByID(id uuid.UUID) (*World, error) {
	for _, world := range s.Worlds {
		if world.WorldID != id {
			continue
		}

		return world, nil
	}

	return nil, fmt.Errorf("world %q not found", id)
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
			cfg.Delete(world.WorldID.String())
			s.cfgManager.SaveConfigWithFileName(s.ConfigFileName())
		}

		return nil
	}

	return fmt.Errorf("world %q not found", name)
}
