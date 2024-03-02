package email

import (
	"torchbearer/pkg/services/config"
)

func (s *Service) ConfigFileName() string {
	return "email.json"
}

func (s *Service) DefaultConfig() (config config.Config) {
	g := config.Group("auth")

	g.Set("user", "changeme@gmail.com")
	g.Set("password", "changeme")

	return
}
