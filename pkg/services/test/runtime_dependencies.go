package test

import (
	"github.com/gravestench/runtime/pkg"

	"torchbearer/pkg/services/dice"
)

func (s *Service) DependenciesResolved() bool {
	if s.dice == nil {
		return false
	}

	return true
}

func (s *Service) ResolveDependencies(rt pkg.IsRuntime) {
	for _, service := range rt.Services() {
		switch candidate := service.(type) {
		case dice.DiceManager:
			s.dice = candidate
		}
	}
}
