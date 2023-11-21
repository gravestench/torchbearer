package world

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"torchbearer/pkg/models"
)

type tuiWorldView struct {
	root           *tuiRoot
	World          *World
	settlementView *tuiSettlementView
	input          struct {
		name       textinput.Model
		settlement struct {
			index int
			view  bool
		}
	}
}

func (t *tuiWorldView) Init() tea.Cmd {
	t.settlementView = &tuiSettlementView{
		parent: t,
	}

	return nil
}

func (t *tuiWorldView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if t.World == nil {
		t.root.mode = modeList
		return t, nil
	}

	if t.input.settlement.view {
		return t.settlementView.Update(msg)
	}

	switch v := msg.(type) {
	case tea.KeyMsg:
		switch v.String() {
		case "enter":
			t.input.settlement.view = true
		case "up":
			t.input.settlement.index--
		case "down":
			t.input.settlement.index++
		case "backspace":
			t.root.mode = modeList
		}
	}

	t.input.settlement.index = clamp(t.input.settlement.index, 0, len(t.World.Settlements))

	return t, nil
}

func (t *tuiWorldView) View() string {
	if t.input.settlement.view {
		t.settlementView.Settlement = t.World.Settlements[t.input.settlement.index]
		return t.settlementView.View()
	}

	return lipgloss.JoinVertical(lipgloss.Center, t.headerLine(), t.viewContent(), t.viewFooter())
}

func (t *tuiWorldView) headerLine() (output string) {
	style := lipgloss.NewStyle().
		Align(lipgloss.Center).
		Width(40).
		Padding(0, 1).
		BorderStyle(lipgloss.NormalBorder()).
		BorderBottom(true).
		Foreground(lipgloss.Color("#9a9a9a"))

	highlight := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#bababa"))

	return style.Render(fmt.Sprintf("Viewing world: '%s'", highlight.Render(t.World.Name)))
}

func (t *tuiWorldView) viewContent() string {
	selectedSettlement := t.World.Settlements[t.input.settlement.index]
	return lipgloss.JoinHorizontal(lipgloss.Top,
		t.viewColumnStats(),
		t.viewColumnSettlements(),
		t.viewColumnSettlementDescription(selectedSettlement))
}

func (t *tuiWorldView) viewColumnStats() string {
	styleTitle := lipgloss.NewStyle().
		Align(lipgloss.Center).
		Padding(0, 1).
		Underline(true).
		Foreground(lipgloss.Color("#9a9a9a"))

	title := styleTitle.Render("Stats")

	body := strings.Join([]string{
		fmt.Sprintf("Sessions Played: %v", t.World.Stats.SessionsPlayed),
		fmt.Sprintf("Tests Rolled: %v", t.World.Stats.TestsRolled),
	}, "\n")

	return lipgloss.JoinVertical(lipgloss.Center, title, body)
}

func (t *tuiWorldView) viewColumnSettlements() string {
	titleColor := lipgloss.Color("#9a9a9a")

	styleTitle := lipgloss.NewStyle().
		Align(lipgloss.Center).
		Padding(0, 1).
		Underline(true).
		Foreground(titleColor)

	focus := lipgloss.NewStyle().Background(lipgloss.Color("#7D56F4"))

	title := styleTitle.Render("Settlements")

	settlementLines := make([]string, 0)

	for idx, settlement := range t.World.Settlements {
		name := settlement.Name
		if idx == t.input.settlement.index {
			name = focus.Render(name)
		}
		settlementLines = append(settlementLines, name)
	}

	body := lipgloss.JoinVertical(lipgloss.Left, settlementLines...)

	return lipgloss.JoinVertical(lipgloss.Center, title, body)
}

func (t *tuiWorldView) viewColumnSettlementDescription(settlement *models.Settlement) string {
	titleColor := lipgloss.Color("#9a9a9a")

	styleTitle := lipgloss.NewStyle().
		Align(lipgloss.Center).
		Padding(0, 1).
		Underline(true).
		Foreground(titleColor)

	styleDescription := lipgloss.NewStyle().
		Align(lipgloss.Center).
		Padding(0, 1).
		Width(40)

	desc := settlement.Description()

	return lipgloss.JoinVertical(lipgloss.Center, styleTitle.Render("Description"), styleDescription.Render(desc))
}

//func (t *tuiWorldView) viewColumnGovernments() string {
//	titleColor := lipgloss.Color("#9a9a9a")
//	if t.input.submode == worldViewSubModeGovernment {
//		titleColor = lipgloss.Color("#9f9f9f")
//	}
//
//	styleTitle := lipgloss.NewStyle().
//		Align(lipgloss.Center).
//		Padding(0, 1).
//		Underline(true).
//		Foreground(titleColor)
//
//	focus := lipgloss.NewStyle().Background(lipgloss.Color("#7D56F4"))
//
//	title := styleTitle.Render("Government")
//
//	body := lipgloss.JoinVertical(lipgloss.Left, governmentLines...)
//
//	return lipgloss.JoinVertical(lipgloss.Center, title, body)
//}

func (m *tuiWorldView) viewFooter() string {
	styleTitle := lipgloss.NewStyle().
		Align(lipgloss.Center).
		Padding(0, 1).
		Underline(true).
		Foreground(lipgloss.Color("#9a9a9a"))

	styleItem := lipgloss.NewStyle().Align(lipgloss.Center)
	styleLabel := lipgloss.NewStyle().Padding(0, 1).Foreground(lipgloss.Color("#7D56F4")).Align(lipgloss.Right)
	styleHotkey := lipgloss.NewStyle().Padding(0, 1).Background(lipgloss.Color("#7D56F4")).Bold(true).Align(lipgloss.Center)

	var footer string

	hotkeys := [][2]string{
		{"Create", "c"},
		{"View", "v"},
		{"Delete", "delete"},
		{"Save", "ctrl+s"},
		{"Reload", "ctrl+l"},
		{"", "↑"},
		{"Select", ""},
		{"", "↓"},
	}

	for _, entry := range hotkeys {
		label := entry[0]
		hotkey := entry[1]
		label = styleLabel.Render(label)
		if hotkey != "" {
			hotkey = styleHotkey.Render(hotkey)
		}
		footer += styleItem.Render(label + hotkey)
	}

	styleBorder := lipgloss.NewStyle().Border(lipgloss.RoundedBorder())

	return lipgloss.JoinVertical(lipgloss.Left, styleTitle.Render(), styleBorder.Render(footer))
}
