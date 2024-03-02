package world

import (
	"fmt"
	"math/rand"

	"github.com/google/uuid"

	"torchbearer/pkg/models"
)

type World struct {
	WorldID     uuid.UUID
	Name        string
	Seed        int64
	AsciiMap    string
	rng         *rand.Rand
	Settlements []*models.Settlement
	Stats       struct {
		SessionsPlayed int
		TestsRolled    int
	}
	*Service `json:"-"`
}

func (w *World) GetSettlementByName(name string) (*models.Settlement, error) {
	for _, settlement := range w.Settlements {
		if settlement.Name == name {
			return settlement, nil
		}
	}

	return nil, fmt.Errorf("settlement with name %q not found", name)
}

func (w *World) GetSettlementByID(id uuid.UUID) (*models.Settlement, error) {
	for _, settlement := range w.Settlements {
		if settlement.SettlementID == id {
			return settlement, nil
		}
	}

	return nil, fmt.Errorf("settlement with ID %q not found", id.String())
}

func (w *World) GetAdventurerByName(name string) (*models.Adventurer, error) {
	return w.adventurer.GetAdventurerByName(name)
}

func (w *World) GetAdventurerByID(id uuid.UUID) (*models.Adventurer, error) {
	return w.adventurer.GetAdventurerByID(id)
}

func (w *World) GetTownsfolkByName(name string) (*models.Townsfolk, error) {
	for _, settlement := range w.Settlements {
		if t, err := settlement.GetTownsfolkByName(name); err == nil {
			return t, nil
		}
	}

	return nil, fmt.Errorf("townsfolk with name %q not found", name)
}

func (w *World) GetTownsfolkByID(id uuid.UUID) (*models.Townsfolk, error) {
	for _, settlement := range w.Settlements {
		if t, err := settlement.GetTownsfolkByID(id); err == nil {
			return t, nil
		}
	}

	return nil, fmt.Errorf("townsfolk with ID %q not found", id.String())
}
