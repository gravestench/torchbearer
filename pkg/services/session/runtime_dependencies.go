package session

import (
	"github.com/gravestench/runtime/pkg"

	"torchbearer/pkg/services/phase"
)

func (s *Service) DependenciesResolved() bool {
	if s.phase == nil {
		return false
	}

	return true
}

func (s *Service) ResolveDependencies(rt pkg.IsRuntime) {
	for _, service := range rt.Services() {
		switch candidate := service.(type) {
		case phase.PhaseManager:
			s.phase = candidate
		}
	}
}
