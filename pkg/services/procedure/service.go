package procedure

import (
	"fmt"
	"log/slog"
	"sync"

	"github.com/google/uuid"
	"github.com/gravestench/servicemesh"
)

type Service struct {
	logger     *slog.Logger
	mux        sync.Mutex
	generators map[string]Procedure
	instances  map[uuid.UUID]Procedure
}

func (s *Service) SetLogger(logger *slog.Logger) {
	s.logger = logger
}

func (s *Service) Logger() *slog.Logger {
	return s.logger
}

func (s *Service) Init(mesh servicemesh.Mesh) {
	s.generators = make(map[string]Procedure)
	s.instances = make(map[uuid.UUID]Procedure)
}

func (s *Service) Name() string {
	return "Procedure Manager"
}

func (s *Service) Register(generator Procedure) error {
	s.mux.Lock()
	defer s.mux.Lock()

	if _, found := s.generators[generator.Name]; found {
		return fmt.Errorf("blueprint with name %q already exists", generator.Name)
	}

	s.generators[generator.Name] = generator

	return nil
}

func (s *Service) Deregister(generator Procedure) error {
	s.mux.Lock()
	defer s.mux.Lock()

	if _, found := s.generators[generator.Name]; !found {
		return fmt.Errorf("blueprint with name %q does not exist", generator.Name)
	}

	delete(s.generators, generator.Name)

	return nil
}

func (s *Service) Begin(name string) (*Procedure, error) {
	s.mux.Lock()
	defer s.mux.Lock()

	generator, found := s.generators[name]
	if !found {
		return nil, fmt.Errorf("procedure blueprint not found")
	}

	return generator.New(), nil
}

func (s *Service) End(instance Procedure) error {
	s.mux.Lock()
	defer s.mux.Lock()

	if _, found := s.instances[instance.UUID()]; !found {
		return fmt.Errorf("procedure not found")
	}

	delete(s.instances, instance.UUID())

	return nil
}
