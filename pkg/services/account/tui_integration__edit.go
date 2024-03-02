package account

import (
	tea "github.com/charmbracelet/bubbletea"
)

type tuiEditAccount struct {
	s *Service
}

func (t *tuiEditAccount) Init() tea.Cmd {
	return nil
}

func (t *tuiEditAccount) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return t, nil
}

func (t *tuiEditAccount) View() string {
	return ""
}
