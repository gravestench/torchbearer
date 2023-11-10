package adventurer

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"torchbearer/pkg/models"
)

const (
	chooseName int = iota
	chooseClass
	humanUpbringing
	chooseHometown
	chooseSocialGraces
	chooseSpecialty
	chooseWises
	answerNatureQuestions
	answerCirclesQuestions
	noteResources
	chooseGear
	chooseWeapon
	chooseArmor
	rollForSpellsAndRelics
	numSteps
	cancelConfirm
)

type tuiCharacterCreation struct {
	*Service
	Adventurer *models.Adventurer
	step       int
	lastStep   int
	input      struct {
		name  textinput.Model
		stock struct {
			choices []models.Stock
			index   int
		}
		humanUpbringing struct {
			choices []models.SkillRecord
			index   int
		}
	}
}

func (t *tuiCharacterCreation) Init() tea.Cmd {
	t.Adventurer = &models.Adventurer{}

	t.input.name = textinput.New()
	t.input.name.Placeholder = t.GenerateAdventurerName()
	t.input.name.Prompt = ""
	t.input.name.Focus()
	t.input.name.CharLimit = 30
	t.input.name.Width = 30

	for _, stock := range t.records.Stocks() {
		t.input.stock.choices = append(t.input.stock.choices, stock)
	}

	for _, skillName := range []string{
		"Criminal",
		"Laborer",
		"Haggler",
		"Pathfinder",
		"Peasant",
		"Survivalist",
	} {
		record, err := t.records.GetSkill(skillName)
		if err != nil {
			t.logger.Fatal().Msgf("selecting skill %q: %v", skillName, err)
		}

		t.input.humanUpbringing.choices = append(t.input.humanUpbringing.choices, *record)
	}

	return nil
}

func (t *tuiCharacterCreation) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if t.step < chooseName {
		t.step = chooseName
	}

	switch t.step {
	case chooseName:
		return t.updateChooseName(msg)
	case chooseClass:
		return t.updateChooseClass(msg)
	case humanUpbringing:
		return t.updateHumanUpbringing(msg)
	case chooseHometown:
		return t.updateChooseHometown(msg)
	case chooseSocialGraces:
		return t.updateChooseSocialGraces(msg)
	case chooseSpecialty:
		return t.updateChooseSpecialty(msg)
	case chooseWises:
		return t.updateChooseWises(msg)
	case answerNatureQuestions:
		return t.updateAnswerNatureQuestions(msg)
	case answerCirclesQuestions:
		return t.updateAnswerCirclesQuestions(msg)
	case noteResources:
		return t.updateNoteResources(msg)
	case chooseGear:
		return t.updateChooseGear(msg)
	case chooseWeapon:
		return t.updateChooseWeapon(msg)
	case chooseArmor:
		return t.updateChooseArmor(msg)
	case rollForSpellsAndRelics:
		return t.updateRollForSpellsAndRelics(msg)
	}

	return t, nil
}

func (t *tuiCharacterCreation) headerLine() (output string) {
	style := lipgloss.NewStyle().
		Align(lipgloss.Left).
		Width(40).
		Padding(0, 1).
		Foreground(lipgloss.Color("#9a9a9a"))

	return style.Render(fmt.Sprintf("Create Character (step %d/%d)", t.step+1, numSteps))
}

func (t *tuiCharacterCreation) View() string {
	header := t.headerLine()

	var content string

	switch t.step {
	case chooseName:
		content = t.viewChooseName()
	case chooseClass:
		content = t.viewChooseStock()
	case humanUpbringing:
		content = t.viewHumanUpbringing()
	case chooseHometown:
		content = t.viewChooseHometown()
	case chooseSocialGraces:
		content = t.viewChooseSocialGraces()
	case chooseSpecialty:
		content = t.viewChooseSpecialty()
	case chooseWises:
		content = t.viewChooseWises()
	case answerNatureQuestions:
		content = t.viewAnswerNatureQuestions()
	case answerCirclesQuestions:
		content = t.viewAnswerCirclesQuestions()
	case noteResources:
		content = t.viewNoteResources()
	case chooseGear:
		content = t.viewChooseGear()
	case chooseWeapon:
		content = t.viewChooseWeapon()
	case chooseArmor:
		content = t.viewChooseArmor()
	case rollForSpellsAndRelics:
		content = t.viewRollForSpellsAndRelics()
	}

	return lipgloss.JoinVertical(lipgloss.Left, header, content)
}

func (t *tuiCharacterCreation) updateChooseName(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			t.Adventurer.Name = t.input.name.Value()
			t.step++
		}

		t.input.name, cmd = t.input.name.Update(msg)
	}

	return t, cmd
}

func (t *tuiCharacterCreation) viewChooseName() string {
	question := lipgloss.NewStyle().Margin(0, 1).Foreground(lipgloss.Color("#ababab"))

	return lipgloss.JoinVertical(lipgloss.Top, question.Render("What is your name? ", t.input.name.View()))
}

func (t *tuiCharacterCreation) updateChooseClass(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up":
			t.input.stock.index--
			if t.input.stock.index < 0 {
				t.input.stock.index = len(t.input.stock.choices) - 1
			}
		case "down":
			t.input.stock.index++
			if t.input.stock.index >= len(t.input.stock.choices) {
				t.input.stock.index = t.input.stock.index % len(t.input.stock.choices)
			}
		case "enter":
			t.Adventurer.Stock = t.input.stock.choices[t.input.stock.index]
			t.step++
		}
	}

	return t, nil
}

func (t *tuiCharacterCreation) viewChooseStock() (output string) {
	question := lipgloss.NewStyle().Padding(1, 0).Foreground(lipgloss.Color("#ababab"))
	center := lipgloss.NewStyle().Width(60).Align(lipgloss.Center)

	top := center.Render(question.Render("Choose your stock"))

	styleFocused := lipgloss.NewStyle().Bold(true).Background(lipgloss.Color("#AA2299"))
	styleLeft := lipgloss.NewStyle().Padding(0, 1).Align(lipgloss.Right)
	styleRight := lipgloss.NewStyle().Width(60).Padding(0, 1).Align(lipgloss.Left)

	var left string
	var right string

	lines := make([]string, 0)

	for idx, choice := range t.input.stock.choices {
		line := fmt.Sprintf("%s %s", choice.Race, choice.Class)
		selected := t.input.stock.index % len(t.input.stock.choices)

		if idx == selected {
			line = styleFocused.Render(line)
			right = choice.Description
		}

		lines = append(lines, line)
	}

	left = lipgloss.JoinVertical(lipgloss.Top, lines...)
	bottom := lipgloss.JoinHorizontal(lipgloss.Top, styleLeft.Render(left), styleRight.Render(right))

	return lipgloss.JoinVertical(lipgloss.Left, top, bottom)
}

func (t *tuiCharacterCreation) updateHumanUpbringing(msg tea.Msg) (tea.Model, tea.Cmd) {
	if t.Adventurer.Stock.Race != models.Human {
		t.step++
		return t, nil
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up":
			t.input.humanUpbringing.index--
			if t.input.humanUpbringing.index < 0 {
				t.input.humanUpbringing.index = len(t.input.humanUpbringing.choices) - 1
			}
		case "down":
			t.input.humanUpbringing.index++
			if t.input.humanUpbringing.index >= len(t.input.humanUpbringing.choices) {
				t.input.humanUpbringing.index = t.input.humanUpbringing.index % len(t.input.humanUpbringing.choices)
			}
		case "enter":
			t.Adventurer.Skills = append(t.Adventurer.Skills, models.AdventurerSkill{
				SkillRecord: t.input.humanUpbringing.choices[t.input.humanUpbringing.index],
				Level:       1,
			})

			t.step++
		}
	}

	return t, nil
}

func (t *tuiCharacterCreation) viewHumanUpbringing() string {
	pad := lipgloss.NewStyle().Padding(0, 1)
	question := pad.Copy().Foreground(lipgloss.Color("#ababab"))
	styleFocused := lipgloss.NewStyle().Bold(true).Background(lipgloss.Color("#AA2299"))

	var content string

	var skillNames []string

	for idx, record := range t.input.humanUpbringing.choices {
		if idx%len(t.input.humanUpbringing.choices) == t.input.humanUpbringing.index {
			skillNames = append(skillNames, styleFocused.Render(record.Name))
			continue
		}

		skillNames = append(skillNames, record.Name)
	}

	content = pad.Copy().BorderRight(true).Render(lipgloss.JoinVertical(lipgloss.Right, skillNames...))

	selected := t.input.humanUpbringing.choices[t.input.humanUpbringing.index]
	styleDescription := lipgloss.NewStyle().Padding(0, 1).Width(40)
	content = lipgloss.JoinHorizontal(lipgloss.Top, content, styleDescription.Render(selected.Description))

	return lipgloss.JoinVertical(lipgloss.Left, question.Padding(1, 0).Render("Your human upbringing grants you one of the following skills:"), content)
}

func (t *tuiCharacterCreation) updateChooseHometown(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		}
	}

	return t, nil
}

func (t *tuiCharacterCreation) viewChooseHometown() (output string) {
	return
}

func (t *tuiCharacterCreation) updateChooseSocialGraces(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		}
	}

	return t, nil
}

func (t *tuiCharacterCreation) viewChooseSocialGraces() (output string) {
	return
}

func (t *tuiCharacterCreation) updateChooseSpecialty(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		}
	}

	return t, nil
}

func (t *tuiCharacterCreation) viewChooseSpecialty() (output string) {
	return
}

func (t *tuiCharacterCreation) updateChooseWises(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		}
	}

	return t, nil
}

func (t *tuiCharacterCreation) viewChooseWises() (output string) {
	return
}

func (t *tuiCharacterCreation) updateAnswerNatureQuestions(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		}
	}

	return t, nil
}

func (t *tuiCharacterCreation) viewAnswerNatureQuestions() (output string) {
	return
}

func (t *tuiCharacterCreation) updateAnswerCirclesQuestions(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		}
	}

	return t, nil
}

func (t *tuiCharacterCreation) viewAnswerCirclesQuestions() (output string) {
	return
}

func (t *tuiCharacterCreation) updateNoteResources(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		}
	}

	return t, nil
}

func (t *tuiCharacterCreation) viewNoteResources() (output string) {
	return
}

func (t *tuiCharacterCreation) updateChooseGear(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		}
	}

	return t, nil
}

func (t *tuiCharacterCreation) viewChooseGear() (output string) {
	return
}

func (t *tuiCharacterCreation) updateChooseWeapon(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		}
	}

	return t, nil
}

func (t *tuiCharacterCreation) viewChooseWeapon() (output string) {
	return
}

func (t *tuiCharacterCreation) updateChooseArmor(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		}
	}

	return t, nil
}

func (t *tuiCharacterCreation) viewChooseArmor() (output string) {
	return
}

func (t *tuiCharacterCreation) updateRollForSpellsAndRelics(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		}
	}

	return t, nil
}

func (t *tuiCharacterCreation) viewRollForSpellsAndRelics() (output string) {
	return
}
