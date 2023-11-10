package dice

import (
	"fmt"
	"sort"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func (s *Service) ModalTui() (name string, model tea.Model) {
	return s.Name(), &tui{Service: s}
}

type tui struct {
	*Service
	toRoll int
}

func (m *tui) Init() tea.Cmd {
	return nil
}

func (m *tui) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.toRoll < 1 {
		m.toRoll = 1
	}

	switch msg := msg.(type) {
	// Is it a key press?
	case tea.KeyMsg:
		// Cool, what was the actual key pressed?
		switch msg.String() {

		case "-":
			if m.toRoll >= 1 {
				m.toRoll--
			}

			return m, nil

		case "+":
			m.toRoll++

			return m, nil

		case "r":
			m.Roll(m.toRoll)
			return m, nil

		case "R":
			times := make([]time.Time, 0)
			for t := range m.rolls {
				times = append(times, t)
			}

			sort.Slice(times, func(i, j int) bool {
				return times[j].Sub(times[i]) > 0
			})

			m.Reroll6s(m.rolls[times[len(times)-1]])
			return m, nil
		}
	}

	return m, nil
}

func (m *tui) View() string {
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

	header := rAlign.Render("Time") +
		lAlign.Render("#Dice") +
		lAlign.Render("Fail") +
		lAlign.Render("Success") +
		lAlign.Render("Success*")

	output += styleHeader.Render(header) + "\r\n"

	times := make([]time.Time, 0)
	for t := range m.rolls {
		times = append(times, t)
	}

	sort.Slice(times, func(i, j int) bool {
		return times[j].Sub(times[i]) > 0
	})

	if len(times) > 10 {
		times = times[len(times)-10:]
	}

	styleFail := lipgloss.NewStyle().Foreground(lipgloss.Color("#FF2222"))
	styleSuccess := lipgloss.NewStyle().Foreground(lipgloss.Color("#22BB22"))
	styleSuccess6 := lipgloss.NewStyle().Foreground(lipgloss.Color("#22dd44"))

	for _, t := range times {
		roll := m.rolls[t]

		numFail := roll.NumFail()
		numSuccess := roll.NumSuccess()
		numSuccess6 := roll.NumSuccess6()

		if numSuccess > 0 && numSuccess6 > 0 {
			numSuccess -= numSuccess6
		}

		fails := fmt.Sprintf("%d", numFail)
		if numFail > 0 {
			fails = styleFail.Render(fails)
		} else {
			fails = ""
		}

		success := fmt.Sprintf("%d", numSuccess)
		if numSuccess > 0 {
			success = styleSuccess.Render(success)
		} else {
			success = ""
		}

		success6 := fmt.Sprintf("%d", numSuccess6)
		if numSuccess6 > 0 {
			success6 = styleSuccess6.Render(success6)
		} else {
			success6 = ""
		}

		output += rAlign.Render(t.Format("15:04:05")) +
			lAlign.Render(fmt.Sprintf("%d", len(roll))) +
			lAlign.Render(fails) +
			lAlign.Render(success) +
			lAlign.Render(success6) +
			"\r\n"
	}

	output += m.footerLine()

	return output
}

func (m *tui) footerLine() string {
	styleItem := lipgloss.NewStyle().Width(18).Align(lipgloss.Center)
	styleLabel := lipgloss.NewStyle().Padding(0, 1).Foreground(lipgloss.Color("#7D56F4")).Align(lipgloss.Right)
	styleHotkey := lipgloss.NewStyle().Padding(0, 1).Background(lipgloss.Color("#7D56F4")).Bold(true).Align(lipgloss.Left)

	var footer string

	hotkeys := []string{
		"Roll::r",
		"Reroll::shift+r",
		"::-",
		fmt.Sprintf("Amount %d::", m.toRoll),
		"::+",
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
