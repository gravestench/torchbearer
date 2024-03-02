package account

import (
	"fmt"
	"log/slog"

	"github.com/google/uuid"
	"github.com/gravestench/servicemesh"

	"torchbearer/pkg/services/config"
	"torchbearer/pkg/services/email"
	"torchbearer/pkg/services/webRouter"
)

var _ servicemesh.Service = &Service{}

type Service struct {
	logger                   *slog.Logger
	cfgManager               config.Dependency
	email                    email.Dependency
	router                   webRouter.Dependency
	accounts                 map[uuid.UUID]Account
	accountRecoveryPasscodes map[uuid.UUID]string // one-time passcodes
	tui                      struct {
		mode   int
		list   tuiListAccounts
		create tuiCreateAccount
		view   tuiViewAccount
		edit   tuiEditAccount
		delete tuiDeleteAccount
	}
}

func (s *Service) Name() string {
	return "Accounts"
}

func (s *Service) Init(mesh servicemesh.Mesh) {
	if err := s.LoadAccounts(); err != nil {
		s.logger.Error("error", fmt.Sprintf("loading accounts: %v", err))
		panic(err)
	}
}

func (s *Service) SetLogger(logger *slog.Logger) {
	s.logger = logger
}

func (s *Service) Logger() *slog.Logger {
	return s.logger
}
