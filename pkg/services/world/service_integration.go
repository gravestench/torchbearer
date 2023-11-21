package world

type Dependency = WorldManager

type WorldManager interface {
	LoadWorlds() error
	SaveWorlds() error
	NewWorld(name string) (*World, error)
	GetWorldByName(name string) (*World, error)
	GetWorlds() []*World
	DeleteWorld(name string) error
	AddWorld(World)
	Ready() bool
}
