package chatgpt

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/ayush6624/go-chatgpt"
	"github.com/gravestench/servicemesh"

	"torchbearer/pkg/services/config"
)

type Service struct {
	logger     *slog.Logger
	client     *chatgpt.Client
	cfgManager config.Dependency
}

func (s *Service) Init(mesh servicemesh.Mesh) {
	cfg, err := s.cfgManager.GetConfigByFileName(s.ConfigFileName())
	if err != nil {
		cfg, err = s.cfgManager.CreateConfigWithFileName(s.ConfigFileName())
		if err != nil {
			s.logger.Error("creating skill records", "error", err)
			mesh.Shutdown()
		}
	}

	key := cfg.Group("api").GetString("key")

	c, err := chatgpt.NewClient(key)
	if err != nil {
		s.logger.Error("could not init client", "error", err)
		return
	}

	s.client = c
}

func (s *Service) Name() string {
	return "ChatGPT"
}

func (s *Service) SetLogger(logger *slog.Logger) {
	s.logger = logger
}

func (s *Service) Logger() *slog.Logger {
	return s.logger
}

func (s *Service) DependenciesResolved() bool {
	if s.cfgManager == nil {
		return false
	}

	return true
}

func (s *Service) ResolveDependencies(services []servicemesh.Service) {
	for _, service := range services {
		switch candidate := service.(type) {
		case config.Dependency:
			s.cfgManager = candidate
		}
	}
}

func (s *Service) ConfigFileName() string {
	return "openai.chatgpt.json"
}

func (s *Service) DefaultConfig() (cfg config.Config) {
	cfg.Group("api").SetDefault("name", "")
	cfg.Group("api").SetDefault("key", "")

	return cfg
}

func (s *Service) Ask(prompt, question string) (*chatgpt.ChatResponse, error) {
	if s.client == nil {
		return nil, fmt.Errorf("client not initialized. Set your api key in %s", s.cfgManager.GetFilePath(s.ConfigFileName()))
	}

	// Set a deadline of 2 seconds from now
	deadline := time.Now().Add(10 * time.Second)

	// Create a context with the deadline
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	res, err := s.client.Send(ctx, &chatgpt.ChatCompletionRequest{
		Model: chatgpt.GPT35Turbo,
		Messages: []chatgpt.ChatMessage{
			{
				Role:    chatgpt.ChatGPTModelRoleSystem,
				Content: prompt,
			},
			{
				Role:    chatgpt.ChatGPTModelRoleUser,
				Content: question,
			},
		},
	})

	if err != nil {
		return nil, fmt.Errorf("sending message: %v", err)
	}

	return res, nil
}
