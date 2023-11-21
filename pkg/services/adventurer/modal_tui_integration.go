package adventurer

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/evertras/bubble-table/table"

	"torchbearer/pkg/models"
	"torchbearer/pkg/services/procedure"
)

const (
	columnKeyName  = "name"
	columnKeyLevel = "level"
	columnKeyAlive = "alive?"

	colorNormal   = "#fa0"
	colorFire     = "#f64"
	colorElectric = "#ff0"
	colorWater    = "#44f"
	colorPlant    = "#8b8"
)

func makeRow(a *models.Adventurer) table.Row {
	return table.NewRow(table.RowData{
		columnKeyName:  a.Name,
		columnKeyLevel: len(a.Stock.ChosenLevelBenefits),
		columnKeyAlive: (a.Condition & models.Dead) == 0,
	})
}

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
	tuiList   table.Model
	tuiCreate tea.Model
}

func (m *tui) Init() tea.Cmd {
	m.tuiList = table.New([]table.Column{})
	m.tuiList.Init()

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
		_, cmd := m.tuiCreate.Update(msg)
		if cmd != nil {
			switch cmd().(type) {
			case procedure.MsgTerminateProcedure:
				m.tuiCreate = nil
				m.mode = modeList
			}
		}

		return m, nil
	}

	m.mode = modeList
	return m.Update(msg)
}

func (m *tui) updateList(msg tea.Msg) (tea.Model, tea.Cmd) {
	rows := make([]table.Row, 0)

	for _, adv := range m.Service.adventurers {
		rows = append(rows, makeRow(adv))
	}

	m.tuiList = m.tuiList.WithColumns([]table.Column{
		table.NewColumn(columnKeyName, "Name", 20),
		table.NewColumn(columnKeyLevel, "Level", 20),
		table.NewColumn(columnKeyAlive, "Alive?", 20),
	}).WithRows(rows).
		WithHeaderVisibility(true).
		WithFooterVisibility(true).
		SelectableRows(true).
		BorderRounded().
		WithPageSize(6).
		SortByDesc(columnKeyName).
		Focused(true)

	switch msg2 := msg.(type) {
	// Is it a key press?
	case tea.KeyMsg:
		// Cool, what was the actual key pressed?
		switch msg2.String() {

		// These keys should exit the program.
		case "c":
			m.mode = modeCreate
			m.tuiCreate = m.Service.CreateAdventurerProcedureTui()
			m.tuiCreate.Init()
		default:
			_, cmd := m.tuiList.Update(msg)
			return m, cmd
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

	content := m.tuiList.View()
	footer := m.footerLine()

	return lipgloss.JoinVertical(lipgloss.Center, styleHeader.Render(header), content, footer)
}

func (m *tui) footerLine() string {
	if m.mode == modeCreate {
		return ""
	}

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
