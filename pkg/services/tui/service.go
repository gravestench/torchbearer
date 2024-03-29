package tui

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/gravestench/servicemesh"

	"torchbearer/pkg/services/config"
)

type Service struct {
	mesh   servicemesh.Mesh
	logger *slog.Logger
	cfg    config.Dependency
	isInit bool
	mux    sync.Mutex
	modalUiModel
}

func (s *Service) Init(mesh servicemesh.Mesh) {
	s.mesh = mesh
	s.modalUiModel.modals = make(map[string]tea.Model)

	dir := s.cfg.ConfigDirectory()
	logPath := expandHomeDirectory(filepath.Join(dir, "output.log"))

	redirect, err := os.OpenFile(logPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		s.logger.Error("opening or creating the file", "error", err)
		s.mesh.Shutdown()
	}

	s.mesh.SetLogHandler(slog.NewJSONHandler(redirect, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))

	clearScreen()

	// bubbletea kicked off using servicemesh event handler

	for _, service := range s.mesh.Services() {
		go s.attemptBindService(service)
	}

	s.isInit = true
}

func (s *Service) Name() string {
	return "Modal TUI"
}

func (s *Service) IsInitialized() bool {
	return s.isInit
}

func (s *Service) OnShutdown() {
	s.mesh.SetLogDestination(os.Stdout)

	dir := s.cfg.ConfigDirectory()
	logPath := filepath.Join(dir, "output.log")

	s.logger.Info("tui disabled, output logged to file", "log path", logPath)
}

// the following methods are boilerplate, but they are used
// by the servicemesh to enforce a standard logging format.

func (s *Service) SetLogger(logger *slog.Logger) {
	s.logger = logger
}

func (s *Service) Logger() *slog.Logger {
	return s.logger
}

func clearScreen() {
	// ANSI escape code to clear the screen
	fmt.Print("\033[H\033[2J")
}

func (s *Service) attemptBindService(service servicemesh.Service) {
	s.mux.Lock()
	defer s.mux.Unlock()

	for !s.isInit {
		time.Sleep(time.Second)
	}

	tui, ok := service.(HasModalTextUserInterface)
	if !ok {
		return
	}

	name, modal := tui.ModalTui()

	s.modals[name] = modal

	if s.selectedModal == "" {
		s.selectedModal = name
	}

	if current, exists := s.modals[s.selectedModal]; exists {
		current.Update(tea.Msg(1))
	}
}

func expandHomeDirectory(path string) string {
	if strings.HasPrefix(path, "~") {
		homeDir, err := os.UserHomeDir()
		if err == nil {
			path = strings.Replace(path, "~", homeDir, 1)
		}
	}
	return path
}
