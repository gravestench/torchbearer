package account

import (
	tea "github.com/charmbracelet/bubbletea"
)

type tuiViewAccount struct {
	s *Service
}

func (t *tuiViewAccount) Init() tea.Cmd {
	return nil
}

func (t *tuiViewAccount) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return t, nil
}

func (t *tuiViewAccount) View() string {
	return ""
}
