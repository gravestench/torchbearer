package webRouter

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
	for _, other := range services {
		if cfg, ok := other.(config.Dependency); ok {
			s.cfgManager = cfg
		}
	}
}
