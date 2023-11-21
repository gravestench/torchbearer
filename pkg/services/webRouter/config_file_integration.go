package webRouter

import (
	"fmt"

	"torchbearer/pkg/services/config"
)

func (s *Service) ConfigFileName() string {
	return "web_router.json"
}

func (s *Service) Config() (*config.Config, error) {
	if s.cfgManager == nil {
		return nil, fmt.Errorf("no config manager service bound")
	}

	return s.cfgManager.GetConfigByFileName(s.ConfigFileName())
}

func (s *Service) DefaultConfig() (cfg config.Config) {
	g := cfg.Group("Gin Route Handler")

	g.SetDefault("debug", true)

	return
}
