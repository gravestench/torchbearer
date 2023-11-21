package webRouter

import (
	"github.com/gravestench/runtime"

	"torchbearer/pkg/services/config"
)

func (s *Service) DependenciesResolved() bool {
	if s.cfgManager == nil {
		return false
	}

	return true
}

func (s *Service) ResolveDependencies(rt runtime.R) {
	for _, other := range rt.Services() {
		if cfg, ok := other.(config.Dependency); ok {
			s.cfgManager = cfg
		}
	}
}
