package account

import (
	tea "github.com/charmbracelet/bubbletea"
)

type tuiCreateAccount struct {
	s *Service
}

func (t *tuiCreateAccount) Init() tea.Cmd {
	return nil
}

func (t *tuiCreateAccount) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return t, nil
}

func (t *tuiCreateAccount) View() string {
	return ""
}
