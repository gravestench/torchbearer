package account

import (
	"github.com/gravestench/servicemesh"

	"torchbearer/pkg/services/config"
	"torchbearer/pkg/services/email"
	"torchbearer/pkg/services/webRouter"
)

func (s *Service) DependenciesResolved() bool {
	if s.cfgManager == nil {
		return false
	}

	if s.email == nil {
		return false
	}

	return true
}

func (s *Service) ResolveDependencies(services []servicemesh.Service) {
	for _, service := range services {
		switch candidate := service.(type) {
		case config.Dependency:
			s.cfgManager = candidate
		case email.Dependency:
			s.email = candidate
		case webRouter.Dependency:
			s.router = candidate
		}
	}
}
