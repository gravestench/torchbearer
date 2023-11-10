package world

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type tuiWorldDeletePrompt struct {
	root  *tuiRoot
	World *World
}

func (t *tuiWorldDeletePrompt) Init() tea.Cmd {
	return nil
}

func (t *tuiWorldDeletePrompt) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if t.World == nil {
		t.root.mode = modeList
		return t, nil
	}

	switch v := msg.(type) {
	case tea.KeyMsg:
		switch v.String() {
		case "Y", "y":
			t.root.Service.DeleteWorld(t.World.Name)
			t.root.mode = modeList
		default:
			t.World = nil
		}
	}

	return t, nil
}

func (t *tuiWorldDeletePrompt) View() string {
	var header, content string

	header = t.headerLine()
	content = t.viewContent()

	return lipgloss.JoinVertical(lipgloss.Center, header, content)
}

func (t *tuiWorldDeletePrompt) headerLine() (output string) {
	style := lipgloss.NewStyle().
		Align(lipgloss.Center).
		Width(60).
		Padding(0, 1).
		Foreground(lipgloss.Color("#9a9a9a"))

	return style.Render("Are you sure you want to delete this world?\npress Y to confirm, any other key to cancel")
}

func (t *tuiWorldDeletePrompt) viewContent() string {
	return ""
}
