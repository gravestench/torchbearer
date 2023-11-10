package test

import (
	"github.com/gravestench/runtime"
	"github.com/rs/zerolog"

	"torchbearer/pkg/services/dice"
)

type Service struct {
	logger *zerolog.Logger
	dice   dice.DiceManager
}

func (s *Service) Init(rt runtime.Runtime) {
}

func (s *Service) Name() string {
	return "Test"
}

func (s *Service) BindLogger(logger *zerolog.Logger) {
	s.logger = logger
}

func (s *Service) Logger() *zerolog.Logger {
	return s.logger
}
