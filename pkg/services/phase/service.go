package phase

import (
	"github.com/gravestench/runtime"
	"github.com/rs/zerolog"
)

type Service struct {
	logger       *zerolog.Logger
	phases       map[string]Phase
	currentPhase string
}

func (s *Service) CurrentPhase() Phase {
	//TODO implement me
	panic("implement me")
}

func (s *Service) NextPhase() Phase {
	//TODO implement me
	panic("implement me")
}

func (s *Service) Init(rt runtime.Runtime) {
	s.phases = make(map[string]Phase)
	s.createAdventurePhase()
}

func (s *Service) Name() string {
	return "Phase"
}

func (s *Service) BindLogger(logger *zerolog.Logger) {
	s.logger = logger
}

func (s *Service) Logger() *zerolog.Logger {
	return s.logger
}
