package session

import (
	"github.com/gravestench/runtime"
	"github.com/rs/zerolog"

	"torchbearer/pkg/services/phase"
)

type Service struct {
	logger *zerolog.Logger
	phase  phase.PhaseManager
}

func (s *Service) Init(rt runtime.Runtime) {
}

func (s *Service) Name() string {
	return "Session"
}

func (s *Service) BindLogger(logger *zerolog.Logger) {
	s.logger = logger
}

func (s *Service) Logger() *zerolog.Logger {
	return s.logger
}
