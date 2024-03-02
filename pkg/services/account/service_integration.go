package account

import (
	"github.com/google/uuid"
	"github.com/gravestench/servicemesh"

	"torchbearer/pkg/services/config"
	"torchbearer/pkg/services/webRouter"
)

type Dependency = AccountManager

type AccountManager interface {
	servicemesh.Service
	webRouter.ProvidesAuthMiddleware
	webRouter.IsRouteInitializer
	config.LoadsConfig
	LoadAccounts() error
	SaveAccounts() error
	Accounts() []Account
	CreateAccount(name, email string) (*Account, error)
	GetAccountByID(id uuid.UUID) (*Account, error)
	GetAccountByName(name string) (*Account, error)
	GetAccountByEmail(email string) (*Account, error)
	UpdateAccount(id uuid.UUID, updated Account) error
	InitiateAccountRecoveryByEmail(email string) (*string, error)
}
