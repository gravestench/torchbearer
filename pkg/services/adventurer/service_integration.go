package adventurer

import (
	"torchbearer/pkg/models"
)

type AdventurerManager interface {
	LoadAdventurers() error
	SaveAdventurers() error
	Adventurers() (map[string]models.Adventurer, error)
	NewAdventurer(name string) models.Adventurer
	RemoveAdventurer(name string) error
	EndAdventurer()
}
