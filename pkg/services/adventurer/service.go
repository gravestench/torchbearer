package adventurer

import (
	"github.com/gravestench/runtime"
	"github.com/rs/zerolog"

	"torchbearer/pkg/models"
	"torchbearer/pkg/services/phase"
	"torchbearer/pkg/services/records"
)

type Service struct {
	logger      *zerolog.Logger
	phase       phase.PhaseManager
	records     records.Dependency
	Adventurers []models.Adventurer
}

func (s *Service) Init(rt runtime.Runtime) {
}

func (s *Service) Name() string {
	return "Adventurer"
}

func (s *Service) BindLogger(logger *zerolog.Logger) {
	s.logger = logger
}

func (s *Service) Logger() *zerolog.Logger {
	return s.logger
}
