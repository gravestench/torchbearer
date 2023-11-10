package world

import (
	"github.com/gravestench/runtime/pkg"

	"torchbearer/pkg/services/config"
	"torchbearer/pkg/services/phase"
	"torchbearer/pkg/services/records"
)

func (s *Service) DependenciesResolved() bool {
	if s.phase == nil {
		return false
	}

	if s.cfgManager == nil {
		return false
	}

	if s.records == nil {
		return false
	}

	if !s.records.Ready() {
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
		case config.Dependency:
			s.cfgManager = candidate
		}
	}
}
