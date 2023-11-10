package tui

import (
	"github.com/gravestench/runtime"

	"torchbearer/pkg/services/config"
)

func (s *Service) DependenciesResolved() bool {
	if s.cfg == nil {
		return false
	}

	return true
}

func (s *Service) ResolveDependencies(rt runtime.Runtime) {
	for _, service := range rt.Services() {
		switch candidate := service.(type) {
		case config.Dependency:
			s.cfg = candidate
		}
	}
}
