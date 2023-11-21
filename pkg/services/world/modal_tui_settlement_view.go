package world

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"torchbearer/pkg/models"
)

type tuiSettlementView struct {
	root       *tuiRoot
	parent     *tuiWorldView
	Settlement *models.Settlement
	input      struct {
		name       textinput.Model
		facilities struct {
			index int
		}
	}
}

func (t *tuiSettlementView) Init() tea.Cmd {
	return nil
}

func (t *tuiSettlementView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if t.Settlement == nil {
		t.parent.input.settlement.view = false
		return t, nil
	}

	switch v := msg.(type) {
	case tea.KeyMsg:
		switch v.String() {
		case "enter":
		case "up":
			t.input.facilities.index--
		case "down":
			t.input.facilities.index++
		case "backspace":
			t.parent.input.settlement.view = false
		}
	}

	return t, nil
}

func (t *tuiSettlementView) View() string {
	return lipgloss.JoinVertical(lipgloss.Center, t.headerLine(), t.viewContent(), t.viewFooter())
}

func (t *tuiSettlementView) headerLine() (output string) {
	style := lipgloss.NewStyle().
		Align(lipgloss.Center).
		Width(40).
		Padding(0, 1).
		BorderStyle(lipgloss.NormalBorder()).
		BorderBottom(true).
		Foreground(lipgloss.Color("#9a9a9a"))

	highlight := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#bababa"))

	return style.Render(fmt.Sprintf("Viewing settlement: '%s'", highlight.Render(t.Settlement.Name)))
}

func (t *tuiSettlementView) viewContent() string {
	styleColumn := lipgloss.NewStyle().Padding(0, 1).Border(lipgloss.RoundedBorder())

	return lipgloss.JoinHorizontal(lipgloss.Top,
		lipgloss.JoinVertical(lipgloss.Left,
			styleColumn.Render(t.viewColumnStats()),
			styleColumn.Render(t.viewColumnCulture()),
		),
		styleColumn.Render(t.viewColumnFacilities()),
	)
}

func (t *tuiSettlementView) viewColumnStats() string {
	styleTitle := lipgloss.NewStyle().
		Align(lipgloss.Center).
		Padding(0, 1).
		Underline(true).
		Foreground(lipgloss.Color("#9a9a9a"))

	title := styleTitle.Render("Stats")

	body := []string{
		fmt.Sprintf("Type: %v", t.Settlement.Type),
		fmt.Sprintf("Disasters: %v", t.Settlement.NumDisasters),
		fmt.Sprintf("Economic Level: %v", t.Settlement.EconomicLevel),
	}

	return lipgloss.JoinVertical(lipgloss.Center, title, lipgloss.JoinVertical(lipgloss.Left, body...))
}

func (t *tuiSettlementView) viewColumnFacilities() string {
	titleColor := lipgloss.Color("#FF0000")

	styleTitle := lipgloss.NewStyle().
		Align(lipgloss.Center).
		Padding(0, 1).
		Underline(true).
		Foreground(titleColor)

	focus := lipgloss.NewStyle().Background(lipgloss.Color("#7D56F4"))

	title := styleTitle.Render("Facilities")

	lines := make([]string, 0)

	for idx, facility := range t.Settlement.Facilities.List() {
		if idx == t.input.facilities.index {
			facility = focus.Render(facility)
		}

		lines = append(lines, facility)
	}

	body := lipgloss.JoinVertical(lipgloss.Left, lines...)

	return lipgloss.JoinVertical(lipgloss.Center, title, body)
}

func (t *tuiSettlementView) viewColumnCulture() string {
	titleColor := lipgloss.Color("#9a9a9a")

	styleTitle := lipgloss.NewStyle().
		Align(lipgloss.Center).
		Padding(0, 1).
		Underline(true).
		Foreground(titleColor)

	title := styleTitle.Render("Culture")

	lines := []string{
		fmt.Sprintf("Governemt: %s", t.Settlement.Culture.Government),
		fmt.Sprintf("Shadow Governemt: %s", t.Settlement.Culture.ShadowGovernment),
		fmt.Sprintf("#Laws: %v", len(t.Settlement.Culture.Laws)),
	}

	body := lipgloss.JoinVertical(lipgloss.Left, lines...)

	return lipgloss.JoinVertical(lipgloss.Center, title, body)
}

//func (t *tuiSettlementView) viewColumnGovernments() string {
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

func (m *tuiSettlementView) viewFooter() string {
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
