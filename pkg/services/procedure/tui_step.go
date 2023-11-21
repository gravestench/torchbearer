package procedure

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type tuiStep struct {
	*TuiProcedure
	step  *Step
	input struct {
		text        textinput.Model
		choiceIndex int
	}
	lastError string
}

func (t *tuiStep) Init() tea.Cmd {
	t.input.text = textinput.New()
	t.input.text.Focus()
	t.input.text.SetValue("")
	t.input.choiceIndex = 0

	return nil
}

func (t *tuiStep) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg2 := msg.(type) {
	// Is it a key press?
	case tea.KeyMsg:
		switch msg2.String() {
		case "esc":
			return t, CmdTerminate
		case "enter":
			return t.updateApplyAnswer()
		default:
			return t.updateInput(msg)
		}
	}

	return t, nil
}

func (t *tuiStep) updateApplyAnswer() (tea.Model, tea.Cmd) {
	switch len(t.step.Choices) {
	case 0:
		t.step.Answer = t.input.text.Value()
	default:
		choice := t.step.Choices[t.input.choiceIndex]
		t.step.Answer = choice.Name
	}

	if err := t.step.Complete(); err == nil {
		t.lastError = ""

		if t.step.Procedure.NextStep() == nil {
			if t.step.Procedure.OnComplete != nil {
				t.step.Procedure.OnComplete()
			}

			return t, CmdTerminate
		}

		return t, CmdTuiNextStep
	} else {
		t.lastError = err.Error()
	}

	return t, nil
}

func (t *tuiStep) updateInput(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch len(t.step.Choices) {
	case 0:
		t.input.text, _ = t.input.text.Update(msg)
		return t, nil
	default:
		switch msg2 := msg.(type) {
		// Is it a key press?
		case tea.KeyMsg:
			switch msg2.String() {
			case "up":
				t.input.choiceIndex--
			case "down":
				t.input.choiceIndex++
			}
		}

		clamped := clamp(t.input.choiceIndex, 0, len(t.step.Choices))
		t.input.choiceIndex = clamped
	}

	return t, nil
}

func (t *tuiStep) View() string {
	var header, content, footer string

	header = t.viewHeader()

	switch len(t.step.Choices) {
	case 0:
		content = t.viewPromptText()
	default:
		content = t.viewPromptChoice()
	}

	if t.lastError != "" {
		footer = t.viewFooter()
	}

	return lipgloss.JoinVertical(lipgloss.Center, header, content, footer)
}

func (t *tuiStep) viewHeader() string {
	styleHeader := lipgloss.NewStyle().
		Width(60).
		Align(lipgloss.Center).
		Padding(1, 1, 0, 1).
		Foreground(lipgloss.Color("#ef7aef"))

	return styleHeader.Render(t.step.Prompt)
}

func (t *tuiStep) viewFooter() string {
	styleFooter := lipgloss.NewStyle().
		Background(lipgloss.Color("#9F4444")).
		Foreground(lipgloss.Color("#440000"))

	return styleFooter.Render(t.lastError)
}

func (t *tuiStep) viewPromptText() string {
	style := lipgloss.NewStyle().
		Width(40).
		Align(lipgloss.Center).
		Padding(1)

	return style.Render(t.input.text.View())
}

func (t *tuiStep) viewPromptChoice() string {
	var leftColumn, rightColumn string

	pad := lipgloss.NewStyle().Padding(1)
	styleColumn := lipgloss.NewStyle().Padding(0, 1).Border(lipgloss.RoundedBorder())
	styleFocused := lipgloss.NewStyle().Bold(true).Background(lipgloss.Color("#AA2299"))

	var choices []string
	for idx, choice := range t.step.Choices {
		strChoice := choice.Name

		if idx == t.input.choiceIndex {
			strChoice = styleFocused.Render(choice.Name)
		}

		choices = append(choices, strChoice)
	}

	leftColumn = styleColumn.Render(lipgloss.JoinVertical(lipgloss.Right, choices...))
	rightColumn = styleColumn.Copy().Width(40).Render(t.step.Choices[t.input.choiceIndex].Description)

	return pad.Render(lipgloss.JoinHorizontal(lipgloss.Top, leftColumn, rightColumn))
}

func clamp(value, min, max int) int {
	if value < min {
		return min
	} else if value > max-1 {
		return max - 1
	}
	return value
}
