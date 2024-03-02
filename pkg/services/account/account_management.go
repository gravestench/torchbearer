package account

import (
	"fmt"

	"github.com/google/uuid"
)

func (s *Service) LoadAccounts() error {
	cfg, err := s.cfgManager.GetConfigByFileName(s.ConfigFileName())
	if err != nil {
		return fmt.Errorf("loading accounts: %v", err)
	}

	s.LoadConfig(cfg)

	return nil
}

func (s *Service) SaveAccounts() error {
	cfg, err := s.cfgManager.GetConfigByFileName(s.ConfigFileName())
	if err != nil {
		return fmt.Errorf("loading accounts: %v", err)
	}

	for key, account := range s.accounts {
		g := cfg.Group(key.String())

		g.Set("ID", account.ID)
		g.Set("Username", account.Username)
		g.Set("Password", account.Password)
		g.Set("Email", account.Email)
		g.Set("Adventurers", account.Adventurers)
	}

	return s.cfgManager.SaveConfigWithFileName(s.ConfigFileName())
}

func (s *Service) Accounts() (a []Account) {
	for _, account := range s.accounts {
		a = append(a, account)
	}

	return a
}

func (s *Service) CreateAccount(name, email string) (*Account, error) {
	if s.accounts == nil {
		s.accounts = make(map[uuid.UUID]Account)
	}

	for _, account := range s.Accounts() {
		if account.Username == name {
			return nil, fmt.Errorf("account with name %q already exists", name)
		}

		if account.Email == email {
			return nil, fmt.Errorf("account with email %q already exists", email)
		}
	}

	a := Account{
		ID:       uuid.New(),
		Username: name,
		Email:    email,
	}

	s.accounts[a.ID] = a

	if err := s.SendOneTimePassCode(a.Email); err != nil {
		delete(s.accounts, a.ID)
		return nil, fmt.Errorf("sending one-time pass code: %v", err)
	}

	if err := s.SaveAccounts(); err != nil {
		delete(s.accounts, a.ID)
		return nil, fmt.Errorf("saving accounts: %v", err)
	}

	return &a, nil
}

func (s *Service) GetAccountByID(id uuid.UUID) (*Account, error) {
	if account, found := s.accounts[id]; found {
		return &account, nil
	}

	return nil, fmt.Errorf("account with id %q not found", id.String())
}

func (s *Service) GetAccountByName(name string) (*Account, error) {
	for _, account := range s.Accounts() {
		if account.Username == name {
			return &account, nil
		}
	}

	return nil, fmt.Errorf("account with name %q not found", name)
}

func (s *Service) GetAccountByEmail(email string) (*Account, error) {
	for _, account := range s.Accounts() {
		if account.Email == email {
			return &account, nil
		}
	}

	return nil, fmt.Errorf("account with email %q not found", email)
}

func (s *Service) UpdateAccount(id uuid.UUID, updated Account) error {
	if _, found := s.accounts[id]; !found {
		return fmt.Errorf("account not found")
	}

	s.accounts[id] = updated

	return s.SaveAccounts()
}
