package world

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	stepCreateChooseName = iota
	stepCreateNumSteps
)

type tuiWorldCreate struct {
	root     *tuiRoot
	World    *World
	step     int
	lastStep int
	input    struct {
		name textinput.Model
	}
}

func (t *tuiWorldCreate) Init() tea.Cmd {
	t.step = stepCreateChooseName

	t.input.name = textinput.New()
	t.World, _ = t.root.Service.NewWorld(t.input.name.Placeholder)
	t.input.name.Placeholder = t.World.generateNewSettlementName()
	t.input.name.Prompt = ""
	t.input.name.Focus()
	t.input.name.CharLimit = 30
	t.input.name.Width = 30

	return nil
}

type WorldCreationComplete struct{}

func (t *tuiWorldCreate) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if t.step < stepCreateChooseName {
		t.step = stepCreateChooseName
	}

	switch t.step {
	case stepCreateChooseName:
		return t.updateChooseName(msg)
	}

	return t, func() tea.Msg { return &WorldCreationComplete{} }
}

func (t *tuiWorldCreate) headerLine() (output string) {
	style := lipgloss.NewStyle().
		Align(lipgloss.Left).
		Width(40).
		Padding(0, 1).
		Foreground(lipgloss.Color("#9a9a9a"))

	return style.Render(fmt.Sprintf("Create New World (step %d/%d)", t.step+1, stepCreateNumSteps))
}

func (t *tuiWorldCreate) View() string {
	header := t.headerLine()

	var content string

	switch t.step {
	case stepCreateChooseName:
		content = t.viewChooseName()
	}

	return lipgloss.JoinVertical(lipgloss.Left, header, content)
}

func (t *tuiWorldCreate) updateChooseName(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg2 := msg.(type) {
	case tea.KeyMsg:
		switch msg2.String() {
		case "ctrl+r":
			t.input.name.Placeholder = t.World.generateNewSettlementName()
		case "enter":
			t.World.Name = t.input.name.Value()

			if t.input.name.Value() == "" {
				t.World.Name = t.input.name.Placeholder
			}

			t.root.mode = modeList
			if t.World != nil {
				t.root.Service.AddWorld(*t.World)
				t.World = nil
				t.root.Service.SaveWorlds()
			}

			t.step++
		}

		t.input.name, cmd = t.input.name.Update(msg)
	}

	return t, cmd
}

func (t *tuiWorldCreate) viewChooseName() string {
	question := lipgloss.NewStyle().Margin(0, 1).Foreground(lipgloss.Color("#ababab"))

	return lipgloss.JoinVertical(lipgloss.Top, question.Render("What is the name of this world? ", t.input.name.View()))
}
