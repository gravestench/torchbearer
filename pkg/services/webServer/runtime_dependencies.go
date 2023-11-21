package webServer

import (
	"github.com/gravestench/runtime/pkg"

	"torchbearer/pkg/services/config"
	"torchbearer/pkg/services/webRouter"
)

func (s *Service) DependenciesResolved() bool {
	if s.router == nil {
		return false
	}

	if s.cfgManager == nil {
		return false
	}

	return true
}

func (s *Service) ResolveDependencies(runtime pkg.IsRuntime) {
	for _, other := range runtime.Services() {
		if router, ok := other.(webRouter.Dependency); ok {
			if router.RouteRoot() != nil {
				s.router = router
			}
		}

		if cfg, ok := other.(config.Dependency); ok && s.cfgManager == nil {
			s.cfgManager = cfg
		}
	}
}
