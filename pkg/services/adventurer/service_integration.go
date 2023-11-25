package adventurer

import (
	"torchbearer/pkg/models"
)

type Dependency = AdventurerManager

type AdventurerManager interface {
	LoadAdventurers() error
	SaveAdventurers() error
	Adventurers() ([]*models.Adventurer, error)
	NewAdventurer() *models.Adventurer
	AddAdventurer(*models.Adventurer) error
	RemoveAdventurer(name string) error
}
