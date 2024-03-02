package records

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gravestench/servicemesh"

	"torchbearer/pkg/models"
	"torchbearer/pkg/services/config"
)

type Service struct {
	logger     *slog.Logger
	cfgManager config.Dependency
	SkillRecords
	StockRecords
	TraitRecords
	WisesRecords
	loaded bool
}

func (s *Service) Slug() string {
	return "records"
}

func (s *Service) InitRoutes(group *gin.RouterGroup) {
	group.GET("skills", func(c *gin.Context) {
		c.JSON(http.StatusOK, s.SkillRecords)
	})
}

func (s *Service) Init(mesh servicemesh.Mesh) {
	s.SkillRecords = make(SkillRecords)
	s.StockRecords = make(StockRecords)
	s.TraitRecords = make(TraitRecords)
	s.WisesRecords = make(WisesRecords)

	s.initConfigFiles()

	s.loaded = true
}

func (s *Service) Name() string {
	return "Records"
}

func (s *Service) RecordsLoaded() bool {
	return s.loaded
}

func (s *Service) SetLogger(logger *slog.Logger) {
	s.logger = logger
}

func (s *Service) Logger() *slog.Logger {
	return s.logger
}

func (s *Service) Skills() SkillRecords {
	return s.SkillRecords
}

func (s *Service) GetSkillByName(name string) (*models.SkillRecord, error) {
	skill, found := s.SkillRecords[name]
	if !found {
		return nil, fmt.Errorf("skill not found")
	}

	return &skill, nil
}

func (s *Service) Stocks() StockRecords {
	return s.StockRecords
}

func (s *Service) GetStockByName(name string) (*models.Stock, error) {
	stock, found := s.StockRecords[name]
	if !found {
		return nil, fmt.Errorf("stock not found")
	}

	return &stock, nil
}

func (s *Service) Traits() TraitRecords {
	return s.TraitRecords
}

func (s *Service) GetTraitByName(name string) (*models.TraitRecord, error) {
	trait, found := s.TraitRecords[name]
	if !found {
		return nil, fmt.Errorf("trait not found")
	}

	return &trait, nil
}

func (s *Service) Wises() WisesRecords {
	return s.WisesRecords
}

func (s *Service) GetWiseByName(name string) (*models.WiseRecord, error) {
	record, found := s.WisesRecords[name]
	if !found {
		return nil, fmt.Errorf("wise not found")
	}

	return &record, nil
}
