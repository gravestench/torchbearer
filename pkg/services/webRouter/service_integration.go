package webRouter

import (
	"github.com/gin-gonic/gin"
	"github.com/gravestench/servicemesh"

	"torchbearer/pkg/services/config"
)

var (
	_ servicemesh.Service         = &Service{}
	_ servicemesh.HasLogger       = &Service{}
	_ servicemesh.HasDependencies = &Service{}
	_ config.HasDefaultConfig     = &Service{}
	_ IsWebRouter                 = &Service{}
)

type Dependency = IsWebRouter

// IsWebRouter is just responsible for yielding the root route handler.
// Services will use this in order to set up their own routes.
type IsWebRouter interface {
	RouteRoot() *gin.Engine
	Reload()
}

// IsRouteInitializer is a type of service that will
// set up its own web routes using a base route group
type IsRouteInitializer interface {
	servicemesh.Service
	InitRoutes(*gin.RouterGroup)
}

// IsProtectedRouteInitializer is a type of service that will
// set up its own web routes using the protected `api` route group
type IsProtectedRouteInitializer interface {
	servicemesh.Service
	InitProtectedRoutes(*gin.RouterGroup)
}

// HasRouteSlug describes a service that has an identifier that is used
// as a prefix for its subroutes
type HasRouteSlug interface {
	Slug() string
}

// ProvidesAuthMiddleware describes a service that can act as auth middleware
// for the router.
type ProvidesAuthMiddleware interface {
	AuthMiddleware() gin.HandlerFunc
}
