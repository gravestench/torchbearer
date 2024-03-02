package webRouter

import (
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gravestench/servicemesh"

	"torchbearer/pkg/services/config"
)

type Service struct {
	mesh servicemesh.Mesh
	log  *slog.Logger

	cfgManager config.Dependency

	root *gin.Engine
	api  *gin.RouterGroup
	auth []gin.HandlerFunc // gathered from other services

	boundServices map[string]*struct{} // holds 0-size entries

	config struct {
		debug bool
	}

	reloadDebounce time.Time
}

func (s *Service) Init(mesh servicemesh.Mesh) {
	s.mesh = mesh
	gin.SetMode("release")

	go s.beginDynamicRouteBinding(s.mesh)
}

func (s *Service) Name() string {
	return "Web Router"
}

func (s *Service) SetLogger(l *slog.Logger) {
	s.log = l
}

func (s *Service) Logger() *slog.Logger {
	return s.log
}
