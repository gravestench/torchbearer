package world

import (
	"sort"

	"github.com/gravestench/runtime"
	"github.com/rs/zerolog"

	"torchbearer/pkg/services/config"
	"torchbearer/pkg/services/phase"
	"torchbearer/pkg/services/records"
)

type Service struct {
	logger     *zerolog.Logger
	phase      phase.PhaseManager
	cfgManager config.Dependency
	records    records.Dependency
	Worlds     []*World
	ready      bool
}

func (s *Service) Init(rt runtime.Runtime) {
	s.Worlds = make([]*World, 0)
	if err := s.LoadWorlds(); err != nil {
		s.logger.Error().Msgf("loading world config: %v", err)
	}
	s.ready = true
}

func (s *Service) Ready() bool {
	return s.ready
}

func (s *Service) Name() string {
	return "World"
}

func (s *Service) BindLogger(logger *zerolog.Logger) {
	s.logger = logger
}

func (s *Service) Logger() *zerolog.Logger {
	return s.logger
}

func (s *Service) GetSortedWorlds() []*World {
	worlds := make([]*World, 0)

	for _, world := range s.Worlds {
		worlds = append(worlds, world)
	}

	sort.Slice(worlds, func(i, j int) bool {
		return worlds[i].UUID.String() < worlds[i].UUID.String()
	})

	return worlds
}

func (s *Service) GetWorlds() []*World {
	return s.GetSortedWorlds()
}
