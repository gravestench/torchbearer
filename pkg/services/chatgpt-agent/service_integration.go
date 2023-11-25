package chatgpt_agent

import (
	"github.com/google/uuid"
	"github.com/gravestench/runtime"

	"torchbearer/pkg/services/config"
)

type ChatGptAgent interface {
	runtime.Service
	runtime.HasDependencies
	runtime.HasLogger
	config.HasConfig
	UUID() uuid.UUID
	Context() string
	SetContext(string)
	Messages() []Message
	Ask(question string) (response string, err error)
}
