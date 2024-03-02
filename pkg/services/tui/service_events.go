package tui

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/gravestench/servicemesh"
)

func (s *Service) OnServiceMeshRunLoopInitiated() {
	p := tea.NewProgram(&s.modalUiModel)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

func (s *Service) OnServiceAdded(service servicemesh.Service) {
	s.attemptBindService(service)
}
