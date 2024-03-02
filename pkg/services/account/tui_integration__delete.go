package account

import (
	tea "github.com/charmbracelet/bubbletea"
)

type tuiDeleteAccount struct {
	s *Service
}

func (t *tuiDeleteAccount) Init() tea.Cmd {
	return nil
}

func (t *tuiDeleteAccount) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return t, nil
}

func (t *tuiDeleteAccount) View() string {
	return ""
}
