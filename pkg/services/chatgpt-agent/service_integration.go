package chatgpt_agent

import (
	"github.com/google/uuid"

	"torchbearer/pkg/services/config"
)

type ChatGptAgent interface {
	servicemesh.Service
	servicemesh.HasDependencies
	servicemesh.HasLogger
	config.HasConfig
	UUID() uuid.UUID
	Context() string
	SetContext(string)
	Messages() []Message
	Ask(question string) (response string, err error)
}
