package chatgpt

import (
	"github.com/ayush6624/go-chatgpt"
	"github.com/gravestench/runtime"

	"torchbearer/pkg/services/config"
)

type Dependency = ChatGptApi

type ChatGptApi interface {
	runtime.Service
	runtime.HasLogger
	runtime.HasDependencies
	config.HasDefaultConfig
	Ask(prompt, question string) (*chatgpt.ChatResponse, error)
}
