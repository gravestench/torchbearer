package session

import (
	"github.com/gravestench/servicemesh"

	"torchbearer/pkg/services/phase"
)

func (s *Service) DependenciesResolved() bool {
	if s.phase == nil {
		return false
	}

	return true
}

func (s *Service) ResolveDependencies(services []servicemesh.Service) {
	for _, service := range services {
		switch candidate := service.(type) {
		case phase.PhaseManager:
			s.phase = candidate
		}
	}
}
