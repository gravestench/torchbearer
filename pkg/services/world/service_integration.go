package world

import (
	"github.com/google/uuid"
)

type Dependency = WorldManager

type WorldManager interface {
	LoadWorlds() error
	SaveWorlds() error
	NewWorld(name string) (*World, error)
	GetWorldByID(id uuid.UUID) (*World, error)
	GetWorldByName(name string) (*World, error)
	GetWorlds() []*World
	DeleteWorld(name string) error
	AddWorld(World)
	IsLoaded() bool
}
