package email

import (
	"github.com/gravestench/servicemesh"

	"torchbearer/pkg/services/config"
)

func (s *Service) DependenciesResolved() bool {
	if s.cfgManager == nil {
		return false
	}

	return true
}

func (s *Service) ResolveDependencies(services []servicemesh.Service) {
	for _, service := range services {
		switch candidate := service.(type) {
		case config.Dependency:
			s.cfgManager = candidate
		}
	}
}
