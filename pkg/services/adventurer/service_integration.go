package adventurer

import (
	"github.com/google/uuid"

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
	GetAdventurerByName(name string) (*models.Adventurer, error)
	GetAdventurerByID(id uuid.UUID) (*models.Adventurer, error)
}
