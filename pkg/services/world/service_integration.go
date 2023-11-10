package world

type WorldManager interface {
	LoadWorlds() error
	SaveWorlds() error
	NewWorld(name string) (*World, error)
	GetWorld(name string) (*World, error)
	DeleteWorld(name string) error
	AddWorld(World)
}
