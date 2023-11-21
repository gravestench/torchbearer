package procedure

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func NewTui(p *Procedure) *TuiProcedure {
	t := &TuiProcedure{
		procedure: p,
	}

	if step := t.procedure.NextStep(); step != nil {
		t.AddStep(step)
	}

	return t
}

type TuiProcedure struct {
	procedure *Procedure
	tuiSteps  []tuiStep
}

func (t *TuiProcedure) Init() tea.Cmd {
	if t.procedure == nil {
		return nil
	}

	for _, step := range t.procedure.Steps() {
		t.AddStep(step)
	}

	return nil
}

func (t *TuiProcedure) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	idx := t.procedure.StepIndex()

	if idx >= len(t.tuiSteps) {
		return t, CmdTerminate
	}

	tt, tcmd := t.tuiSteps[idx].Update(msg)
	if tcmd != nil {
		v := tcmd()
		switch v.(type) {
		case MsgTuiNextStep:
			nextStep := t.procedure.NextStep()
			if nextStep == nil {
				return t, CmdTerminate
			} else {
				t.AddStep(nextStep)
			}
		case MsgTerminateProcedure:
			return t, CmdTerminate
		}
	}

	return tt, nil
}

func (t *TuiProcedure) View() string {
	idx := t.procedure.StepIndex()

	if idx >= len(t.tuiSteps) {
		return ""
	}

	var header, content, footer string

	header = t.viewHeader()
	content = t.tuiSteps[idx].View()
	footer = t.viewFooter()

	return lipgloss.JoinVertical(lipgloss.Center, header, content, footer)
}

func (t *TuiProcedure) viewHeader() string {
	styleHeader := lipgloss.NewStyle().
		Align(lipgloss.Center).
		Foreground(lipgloss.Color("#ef7aef"))

	name := t.procedure.Name
	line := fmt.Sprintf("%s (step %d)", name, t.procedure.StepIndex()+1)

	return styleHeader.Render(line)
}

func (t *TuiProcedure) viewFooter() string {
	styleBorder := lipgloss.NewStyle().Border(lipgloss.RoundedBorder())
	return styleBorder.Render("press escape to cancel")
}

func (t *TuiProcedure) AddStep(step *Step) {
	tui := tuiStep{
		TuiProcedure: t,
		step:         step,
	}

	tui.Init()

	t.tuiSteps = append(t.tuiSteps, tui)
}
