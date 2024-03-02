package account

import (
	tea "github.com/charmbracelet/bubbletea"
)

const (
	tuiModeList = iota
	tuiModeView
	tuiModeEdit
	tuiModeCreate
	tuiModeDelete
)

func (s *Service) ModalTui() (name string, model tea.Model) {
	switch s.tui.mode {
	case tuiModeList:
		s.tui.list.s = s
		model = &s.tui.list
	case tuiModeView:
		s.tui.view.s = s
		model = &s.tui.view
	case tuiModeCreate:
		s.tui.create.s = s
		model = &s.tui.create
	case tuiModeEdit:
		s.tui.edit.s = s
		model = &s.tui.edit
	case tuiModeDelete:
		s.tui.delete.s = s
		model = &s.tui.delete
	default:
		s.tui.mode = 0
		s.tui.list.s = s
		model = &s.tui.list
	}

	return s.Name(), model
}
