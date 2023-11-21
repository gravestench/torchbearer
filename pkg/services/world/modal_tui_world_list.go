package world

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type tuiWorldList struct {
	root     *tuiRoot
	step     int
	lastStep int
	input    struct {
		selection struct {
			list  []*World
			index int
		}
	}
}

func (t *tuiWorldList) Init() tea.Cmd {
	t.root.Service.LoadWorlds()
	t.updateList()

	return nil
}

func (t *tuiWorldList) updateList() {
	t.input.selection.list = make([]*World, 0)
	t.input.selection.list = append(t.input.selection.list, t.root.Service.GetSortedWorlds()...)
	t.input.selection.index = clamp(t.input.selection.index, 0, len(t.root.Worlds))
}

func (t *tuiWorldList) selectedWorld() *World {
	t.input.selection.list = append(t.input.selection.list, t.root.Service.GetSortedWorlds()...)
	if len(t.input.selection.list) < 1 {
		return nil
	}

	return t.input.selection.list[t.input.selection.index]
}

func (t *tuiWorldList) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	t.updateList()

	switch msg2 := msg.(type) {
	// Is it a key press?
	case tea.KeyMsg:
		switch msg2.String() {
		case "up":
			t.input.selection.index = clamp(t.input.selection.index-1, 0, len(t.root.Service.Worlds))
		case "down":
			t.input.selection.index = clamp(t.input.selection.index+1, 0, len(t.root.Service.Worlds))
		case "delete", "backspace":
			if len(t.root.Worlds) > 0 {
				t.root.delete.World = t.selectedWorld()
				t.root.mode = modePromptDeleteWorld
			}
		case "enter":
			if len(t.root.Worlds) > 0 {
				t.root.view.World = t.selectedWorld()
				t.root.mode = modeViewWorld
			}
		case "c":
			t.root.create.Init()
			t.root.mode = modeCreate
		}
	}

	return t, nil
}

func (t *tuiWorldList) headerLine() (output string) {
	style := lipgloss.NewStyle().
		Align(lipgloss.Left).
		Width(40).
		Padding(0, 1).
		Foreground(lipgloss.Color("#9a9a9a"))

	return style.Render(fmt.Sprintf("Create New World (step %d/%d)", t.step+1, stepCreateNumSteps))
}

func (t *tuiWorldList) View() string {
	styleHeader := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#ef7aef"))

	rAlign := lipgloss.NewStyle().
		Padding(0, 1).
		Align(lipgloss.Right)

	styleColumn := lipgloss.NewStyle().Padding(0, 1)
	styleMap := lipgloss.NewStyle().Border(lipgloss.OuterHalfBlockBorder())
	focus := lipgloss.NewStyle().Background(lipgloss.Color("#7D56F4"))

	var header, content, footer string
	//var colName, colSettlements, colSessions, colTests, colParties string
	var colName string

	{
		rows := []string{rAlign.Render(styleHeader.Render("Name"))}

		for _, world := range t.root.GetSortedWorlds() {
			row := world.Name

			if selected := t.selectedWorld(); selected != nil && world.UUID == selected.UUID {
				row = focus.Render(row)
			}

			rows = append(rows, row)
		}

		colName = lipgloss.JoinVertical(lipgloss.Right, rows...)
	}

	var asciiMap string

	if w := t.selectedWorld(); w != nil {
		asciiMap = ColorizeASCII(t.selectedWorld().AsciiMap)
	}

	content = lipgloss.JoinHorizontal(lipgloss.Top, styleColumn.Render(colName), styleMap.Render(asciiMap))

	footer = t.footerLine()

	return lipgloss.JoinVertical(lipgloss.Left, header, content, footer)
}

func (m *tuiWorldList) footerLine() string {
	styleItem := lipgloss.NewStyle().Align(lipgloss.Center)
	styleLabel := lipgloss.NewStyle().Padding(0, 1).Foreground(lipgloss.Color("#7D56F4")).Align(lipgloss.Right)
	styleHotkey := lipgloss.NewStyle().Padding(0, 1).Background(lipgloss.Color("#7D56F4")).Bold(true).Align(lipgloss.Center)

	var footer string

	hotkeys := []string{
		"Create::c",
		"Delete::delete",
		"Save::ctrl+s",
		"Reload::ctrl+l",
		"::↑",
		"Select::",
		"::↓",
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

func clamp(value, min, max int) int {
	if value < min {
		return min
	} else if value > max-1 {
		return max - 1
	}
	return value
}
