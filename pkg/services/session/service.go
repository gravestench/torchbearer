package session

import (
	"log/slog"

	"github.com/gravestench/servicemesh"

	"torchbearer/pkg/services/phase"
)

type Service struct {
	logger *slog.Logger
	phase  phase.PhaseManager
}

func (s *Service) Init(mesh servicemesh.Mesh) {
}

func (s *Service) Name() string {
	return "Session"
}

func (s *Service) SetLogger(logger *slog.Logger) {
	s.logger = logger
}

func (s *Service) Logger() *slog.Logger {
	return s.logger
}
