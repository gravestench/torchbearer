package world

import (
	tea "github.com/charmbracelet/bubbletea"
)

const (
	modeList = iota
	modeCreate
	modeViewWorld
	modePromptDeleteWorld
)

func (s *Service) ModalTui() (name string, model tea.Model) {
	model = &tuiRoot{Service: s}
	model.Init()
	return s.Name(), model
}

type tuiRoot struct {
	*Service
	mode   int
	list   *tuiWorldList
	create *tuiWorldCreate
	view   *tuiWorldView
	delete *tuiWorldDeletePrompt
}

func (m *tuiRoot) Init() tea.Cmd {
	m.list = &tuiWorldList{
		root: m,
	}

	m.create = &tuiWorldCreate{
		root: m,
	}

	m.view = &tuiWorldView{
		root: m,
	}

	m.delete = &tuiWorldDeletePrompt{
		root: m,
	}

	m.list.Init()
	m.create.Init()
	m.view.Init()
	m.delete.Init()

	return nil
}

func (m *tuiRoot) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	// Is it a key press?
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			m.mode--
		}
	}

	switch m.mode {
	case modeList:
		return m.list.Update(msg)
	case modeCreate:
		mm, res := m.create.Update(msg)
		if res == nil {
			return mm, res
		}
		switch res().(type) {
		case WorldCreationComplete:
			m.Service.SaveWorlds()
			return mm, nil
		}
	case modeViewWorld:
		return m.view.Update(msg)
	case modePromptDeleteWorld:
		return m.delete.Update(msg)
	}

	m.mode = modeList
	return m.Update(msg)
}

func (m *tuiRoot) View() string {
	switch m.mode {
	case modeList:
		return m.list.View()
	case modeCreate:
		return m.create.View()
	case modeViewWorld:
		return m.view.View()
	case modePromptDeleteWorld:
		return m.delete.View()
	}

	m.mode = modeList
	return m.View()
}
