package procedure

import (
	tea "github.com/charmbracelet/bubbletea"
)

type MsgTerminateProcedure struct{}

func CmdTerminate() tea.Msg {
	return MsgTerminateProcedure{}
}

type MsgTuiNextStep struct{}

func CmdTuiNextStep() tea.Msg {
	return MsgTuiNextStep{}
}
