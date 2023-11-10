package config

import (
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func (s *Service) ModalTui() (name string, model tea.Model) {
	return s.Name(), &tui{Service: s}
}

type tui struct {
	*Service
}

func (m *tui) Init() tea.Cmd {
	return nil
}

func (m *tui) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m *tui) View() string {
	list := make([][]string, 0)

	for _, service := range m.rt.Services() {
		if candidate, ok := service.(HasConfig); ok {
			name := service.Name()
			path := filepath.Join(m.ConfigDirectory(), candidate.ConfigFileName())

			list = append(list, []string{name, path})
		}

		if candidate, ok := service.(HasConfigs); ok {
			for _, cfgPath := range candidate.ConfigFileNames() {
				name := service.Name()
				path := filepath.Join(m.ConfigDirectory(), cfgPath)

				list = append(list, []string{name, path})
			}
		}
	}

	styleHeader := lipgloss.NewStyle().Foreground(lipgloss.Color("#ef7aef"))

	rAlign := lipgloss.NewStyle().
		Padding(0, 1).
		Align(lipgloss.Right)

	lAlign := lipgloss.NewStyle().
		Padding(0, 1).
		Align(lipgloss.Left)

	var services, files string

	for idx := range list {
		name := list[idx][0]
		path := list[idx][1]
		if idx == 0 {
			services = styleHeader.Render(rAlign.Render("Service Name"))
			files = styleHeader.Render(lAlign.Render("File Path"))
		}

		services = lipgloss.JoinVertical(lipgloss.Right, services, rAlign.Render(name))
		files = lipgloss.JoinVertical(lipgloss.Left, files, lAlign.Render(path))
	}

	return lipgloss.JoinHorizontal(lipgloss.Top, services, files)
}
