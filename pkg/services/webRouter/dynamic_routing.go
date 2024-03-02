package webRouter

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gravestench/servicemesh"
	"k8s.io/utils/strings/slices"
)

func (s *Service) RouteRoot() *gin.Engine {
	return s.root
}

func (s *Service) Reload() {
	if time.Since(s.reloadDebounce) < time.Second {
		return
	}

	s.reloadDebounce = time.Now()

	s.log.Warn("reloading")

	if s.boundServices == nil {
		s.log.Info("initializing routes")
	} else {
		s.log.Info("re-initializing routes")
	}

	// forget any already-bound services
	s.boundServices = nil
	s.boundServices = make(map[string]*struct{})

	mode := gin.ReleaseMode
	if s.config.debug {
		mode = gin.DebugMode
	}

	s.api = nil

	gin.SetMode(mode)

	// set up the root and protected route groups
	s.root = gin.New()
	s.root.RemoveExtraSlash = true

	s.api = s.root.Group("api")

	s.api.Use(func(c *gin.Context) {
		Logger("gin", s.Logger())(c)
	})

	s.initStaticAssetsMiddleware()
	s.initAuthMiddleware()
}

func (s *Service) beginDynamicRouteBinding(mesh servicemesh.Mesh) {
	for {
		if s.shouldInit(mesh) {
			s.Reload()
		}

		time.Sleep(time.Second)
		s.bindNewRoutes(mesh)
	}
}

func (s *Service) shouldInit(mesh servicemesh.Mesh) bool {
	if s.boundServices == nil {
		return true // base case, happens at app start
	}

	if s.api == nil {
		return true // base case, happens at app start
	}

	// in the event that a service is removed by the
	// service manager for whatever reason, we need to check
	// if that was something that had routes. if it was, we need
	// to re-init the router (we can't actually remove routes in gin)

	// we will check if any of the services that have routes are no longer
	// in the list of the service managers services
	allExistingServiceNames := make([]string, 0)
	for _, candidate := range mesh.Services() {
		if svc, ok := candidate.(IsRouteInitializer); ok {
			allExistingServiceNames = append(allExistingServiceNames, svc.Name())
		}

		if svc, ok := candidate.(IsProtectedRouteInitializer); ok {
			allExistingServiceNames = append(allExistingServiceNames, svc.Name())
		}
	}

	// iterate over each bound service, check if its name
	// exists as a substring inside of our lookup string
	for key, _ := range s.boundServices {
		if !slices.Contains(allExistingServiceNames, key) {
			return true
		}
	}

	return false
}

func (s *Service) bindNewRoutes(mesh servicemesh.Mesh) {
	for _, candidate := range mesh.Services() {
		svcToInit, ok := candidate.(servicemesh.Service)
		if !ok {
			continue
		}

		if svc, ok := candidate.(servicemesh.HasDependencies); ok {
			if !svc.DependenciesResolved() {
				continue
			}
		}

		if _, alreadyBound := s.boundServices[svcToInit.Name()]; alreadyBound {
			continue
		}

		groupPrefix := ""
		if svc, ok := candidate.(HasRouteSlug); ok {
			groupPrefix = svc.Slug()
		}

		// handle route init
		if r, ok := candidate.(IsRouteInitializer); ok {
			r.InitRoutes(s.api.Group(groupPrefix))
			s.boundServices[svcToInit.Name()] = nil // make 0-size entry
			s.log.Info("binding routes", "router handler", svcToInit.Name())
		}

		if r, ok := candidate.(IsProtectedRouteInitializer); ok {
			g := s.api.Group(groupPrefix)
			for _, auth := range s.auth {
				g.Use(auth)
			}

			r.InitProtectedRoutes(g)
			s.boundServices[svcToInit.Name()] = nil // make 0-size entry
			s.log.Info("binding routes", "route handler", svcToInit.Name())
		}
	}
}
