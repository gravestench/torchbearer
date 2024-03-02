package phase

import (
	"log/slog"

	"github.com/gravestench/servicemesh"
)

type Service struct {
	logger       *slog.Logger
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

func (s *Service) Init(mesh servicemesh.Mesh) {
	s.phases = make(map[string]Phase)
	s.createAdventurePhase()
}

func (s *Service) Name() string {
	return "Phase"
}

func (s *Service) Ready() bool {
	return s.phases != nil
}

func (s *Service) SetLogger(logger *slog.Logger) {
	s.logger = logger
}

func (s *Service) Logger() *slog.Logger {
	return s.logger
}
