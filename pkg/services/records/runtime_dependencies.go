package records

import (
	"github.com/gravestench/runtime/pkg"

	"torchbearer/pkg/services/config"
)

func (s *Service) DependenciesResolved() bool {
	if s.cfgManager == nil {
		return false
	}

	return true
}

func (s *Service) ResolveDependencies(rt pkg.IsRuntime) {
	for _, service := range rt.Services() {
		switch candidate := service.(type) {
		case config.Dependency:
			s.cfgManager = candidate
		}
	}
}
