package world

import (
	"github.com/google/uuid"
	"github.com/gravestench/servicemesh"

	"torchbearer/pkg/models"
	"torchbearer/pkg/services/config"
	"torchbearer/pkg/services/phase"
	"torchbearer/pkg/services/records"
)

type AdventurerManager interface {
	LoadAdventurers() error
	SaveAdventurers() error
	Adventurers() ([]*models.Adventurer, error)
	NewAdventurer() *models.Adventurer
	AddAdventurer(*models.Adventurer) error
	RemoveAdventurer(name string) error
	GetAdventurerByName(name string) (*models.Adventurer, error)
	GetAdventurerByID(id uuid.UUID) (*models.Adventurer, error)
}

func (s *Service) DependenciesResolved() bool {
	if s.phase == nil {
		return false
	}

	if s.cfgManager == nil {
		return false
	}

	if s.records == nil {
		return false
	}

	if !s.records.RecordsLoaded() {
		return false
	}

	if s.adventurer == nil {
		return false
	}

	return true
}

func (s *Service) ResolveDependencies(services []servicemesh.Service) {
	for _, service := range services {
		switch candidate := service.(type) {
		case phase.PhaseManager:
			s.phase = candidate
		case records.Dependency:
			s.records = candidate
		case config.Dependency:
			s.cfgManager = candidate
		case AdventurerManager:
			s.adventurer = candidate
		}
	}
}
