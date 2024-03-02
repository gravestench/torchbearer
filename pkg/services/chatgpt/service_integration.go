package chatgpt

import (
	"github.com/ayush6624/go-chatgpt"
	"github.com/gravestench/servicemesh"

	"torchbearer/pkg/services/config"
)

type Dependency = ChatGptApi

type ChatGptApi interface {
	servicemesh.Service
	servicemesh.HasLogger
	servicemesh.HasDependencies
	config.HasDefaultConfig
	Ask(prompt, question string) (*chatgpt.ChatResponse, error)
}
