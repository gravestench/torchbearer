package adventurer

import (
	tea "github.com/charmbracelet/bubbletea"

	"torchbearer/pkg/services/procedure"
)

func (s *Service) CreateAdventurerProcedureTui() tea.Model {
	return &tuiCharacterCreation{Service: s}
}

type tuiCharacterCreation struct {
	*Service
	creation  *procedureCreateAdventurer
	tuiCreate *procedure.TuiProcedure
}

func (t *tuiCharacterCreation) Init() tea.Cmd {
	t.creation = t.Service.CreateAdventurerProcedure()
	t.tuiCreate = procedure.NewTui(t.creation.Procedure)

	return nil
}

func (t *tuiCharacterCreation) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return t.tuiCreate.Update(msg)
}

func (t *tuiCharacterCreation) View() string {
	return t.tuiCreate.View()
}
