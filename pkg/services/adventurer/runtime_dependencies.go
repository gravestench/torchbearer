package adventurer

import (
	"github.com/gravestench/runtime/pkg"

	"torchbearer/pkg/services/config"
	"torchbearer/pkg/services/phase"
	"torchbearer/pkg/services/records"
	"torchbearer/pkg/services/world"
)

func (s *Service) DependenciesResolved() bool {
	if s.phase == nil {
		return false
	}

	if s.records == nil {
		return false
	}

	if !s.records.Ready() {
		return false
	}

	if s.worlds == nil {
		return false
	}

	if !s.worlds.Ready() {
		return false
	}

	if s.config == nil {
		return false
	}

	return true
}

func (s *Service) ResolveDependencies(rt pkg.IsRuntime) {
	for _, service := range rt.Services() {
		switch candidate := service.(type) {
		case phase.PhaseManager:
			s.phase = candidate
		case records.Dependency:
			s.records = candidate
		case world.Dependency:
			s.worlds = candidate
		case config.Dependency:
			s.config = candidate
		}
	}
}
