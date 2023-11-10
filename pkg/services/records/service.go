package records

import (
	"fmt"

	"github.com/gravestench/runtime"
	"github.com/rs/zerolog"

	"torchbearer/pkg/models"
	"torchbearer/pkg/services/config"
)

type Service struct {
	logger     *zerolog.Logger
	cfgManager config.Dependency
	SkillRecords
	StockRecords
	TraitRecords
	ready bool
}

func (s *Service) Init(rt runtime.Runtime) {
	s.SkillRecords = make(SkillRecords)
	s.StockRecords = make(StockRecords)
	s.TraitRecords = make(TraitRecords)

	s.initConfigFiles()
	s.ready = true
}

func (s *Service) Name() string {
	return "Records"
}

func (s *Service) Ready() bool {
	return s.ready
}

func (s *Service) BindLogger(logger *zerolog.Logger) {
	s.logger = logger
}

func (s *Service) Logger() *zerolog.Logger {
	return s.logger
}

func (s *Service) Skills() SkillRecords {
	return s.SkillRecords
}

func (s *Service) GetSkill(name string) (*models.SkillRecord, error) {
	skill, found := s.SkillRecords[name]
	if !found {
		return nil, fmt.Errorf("skill not found")
	}

	return &skill, nil
}

func (s *Service) Stocks() StockRecords {
	return s.StockRecords
}

func (s *Service) GetStock(name string) (*models.Stock, error) {
	stock, found := s.StockRecords[name]
	if !found {
		return nil, fmt.Errorf("stock not found")
	}

	return &stock, nil
}

func (s *Service) Traits() TraitRecords {
	return s.TraitRecords
}

func (s *Service) GetTraits(name string) (*models.TraitRecord, error) {
	trait, found := s.TraitRecords[name]
	if !found {
		return nil, fmt.Errorf("stock not found")
	}

	return &trait, nil
}
