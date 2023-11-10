package adventurer

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	modeList = iota
	modeCreate
)

func (s *Service) ModalTui() (name string, model tea.Model) {
	return s.Name(), &tui{Service: s}
}

type tui struct {
	*Service
	mode      int
	tuiCreate *tuiCharacterCreation
}

func (m *tui) Init() tea.Cmd {
	return nil
}

func (m *tui) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	// Is it a key press?
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			m.mode--
		}
	}

	switch m.mode {
	case modeList:
		return m.updateList(msg)
	case modeCreate:
		return m.tuiCreate.Update(msg)
	}

	m.mode = modeList
	return m.Update(msg)
}

func (m *tui) updateList(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	// Is it a key press?
	case tea.KeyMsg:
		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "c":
			m.mode = modeCreate
			m.tuiCreate = &tuiCharacterCreation{Service: m.Service}
			m.tuiCreate.Init()
		}
	}

	return m, nil
}

func (m *tui) View() string {
	switch m.mode {
	case modeList:
		return m.viewList()
	case modeCreate:
		return m.tuiCreate.View()
	}

	m.mode = modeList
	return m.View()
}

func (m *tui) viewList() string {
	var output string

	styleHeader := lipgloss.NewStyle().Foreground(lipgloss.Color("#ef7aef"))

	rAlign := lipgloss.NewStyle().
		Width(20).
		Padding(0, 1).
		Align(lipgloss.Right)

	lAlign := lipgloss.NewStyle().
		Width(15).
		Padding(0, 1).
		Align(lipgloss.Left)

	header := rAlign.Render("Name") +
		lAlign.Render("Level") +
		lAlign.Render("Alive") +
		lAlign.Render("In Party")

	output += styleHeader.Render(header) + "\r\n"

	output += m.footerLine()

	return output
}

func (m *tui) footerLine() string {
	styleItem := lipgloss.NewStyle().Width(18).Align(lipgloss.Center)
	styleLabel := lipgloss.NewStyle().Padding(0, 1).Foreground(lipgloss.Color("#7D56F4")).Align(lipgloss.Right)
	styleHotkey := lipgloss.NewStyle().Padding(0, 1).Background(lipgloss.Color("#7D56F4")).Bold(true).Align(lipgloss.Left)

	var footer string

	hotkeys := []string{
		"Create::c",
		"Delete::delete",
		"Save::ctrl+s",
		"Reload::ctrl+l",
		"::⬆️",
		"Select::",
		"::⬇️",
	}

	for _, entry := range hotkeys {
		label := strings.Split(entry, "::")[0]
		hotkey := strings.Split(entry, "::")[1]
		label = styleLabel.Render(label)
		if hotkey != "" {
			hotkey = styleHotkey.Render(hotkey)
		}
		footer += styleItem.Render(label + hotkey)
	}

	styleBorder := lipgloss.NewStyle().Border(lipgloss.RoundedBorder())

	return styleBorder.Render(footer)
}
