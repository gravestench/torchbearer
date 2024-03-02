package account

import (
	"sort"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type tuiListAccounts struct {
	s *Service
}

func (t *tuiListAccounts) Init() tea.Cmd {
	return nil
}

func (t *tuiListAccounts) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return t, nil
}

func (t *tuiListAccounts) View() string {
	footer := t.footerLine()

	return lipgloss.JoinVertical(lipgloss.Center, t.content(), footer)
}

func (t *tuiListAccounts) footerLine() string {
	styleItem := lipgloss.NewStyle().Width(18).Align(lipgloss.Center)
	styleLabel := lipgloss.NewStyle().Padding(0, 1).Foreground(lipgloss.Color("#7D56F4")).Align(lipgloss.Right)
	styleHotkey := lipgloss.NewStyle().Padding(0, 1).Background(lipgloss.Color("#7D56F4")).Bold(true).Align(lipgloss.Left)

	var footer string

	hotkeys := []string{
		"Create::c",
		"Delete::delete",
		"Save::s",
		"Reload::l",
		"::⬆",
		"View::enter",
		"::⬇",
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

func (t *tuiListAccounts) content() (content string) {
	accounts := t.s.Accounts()
	sort.Slice(accounts, func(i, j int) bool {
		return accounts[i].Username < accounts[j].Username
	})

	border := lipgloss.NewStyle().BorderStyle(lipgloss.RoundedBorder())
	header := lipgloss.NewStyle().Foreground(lipgloss.Color("#ef7aef"))
	pad := lipgloss.NewStyle().Padding(0, 1)

	var names, emails []string

	names = append(names, header.Render("Username"))
	emails = append(emails, header.Render("Email"))

	for _, account := range accounts {
		names = append(names, account.Username)
		emails = append(emails, account.Email)
	}

	colNames := lipgloss.JoinVertical(lipgloss.Right, names...)
	colEmails := lipgloss.JoinVertical(lipgloss.Left, emails...)
	content = lipgloss.JoinHorizontal(lipgloss.Top, pad.Render(colNames), pad.Render(colEmails))

	return border.Render(content)
}
