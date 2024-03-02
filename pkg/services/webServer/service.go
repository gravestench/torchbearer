package webServer

import (
	"log/slog"
	"net/http"

	"github.com/gravestench/servicemesh"

	"torchbearer/pkg/services/config"
	"torchbearer/pkg/services/webRouter"
)

type Service struct {
	log        *slog.Logger
	router     webRouter.Dependency
	cfgManager config.Dependency
	server     *http.Server
	lastConfig string
}

func (s *Service) Init(mesh servicemesh.Mesh) {
	s.StartServer()
}

func (s *Service) Name() string {
	return "Web Server"
}

func (s *Service) SetLogger(l *slog.Logger) {
	s.log = l
}

func (s *Service) Logger() *slog.Logger {
	return s.log
}
