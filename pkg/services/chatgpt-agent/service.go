package chatgpt_agent

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"

	"torchbearer/pkg/services/chatgpt"
	"torchbearer/pkg/services/config"
)

const (
	keyID       = "ID"
	keyContext  = "Context"
	keyMessages = "Messages"
)

func NewFromConfig(cfg config.Object) (*Service, error) {
	id, err := uuid.Parse(cfg.GetString(keyID))
	if err != nil {
		return nil, fmt.Errorf("parsing UUID from config: %v", err)
	}

	s := &Service{
		id:      &id,
		context: cfg.GetString(keyContext),
	}

	messageData := cfg.GetJson(keyMessages)
	if err = json.Unmarshal(messageData, &s.messages); err != nil {
		return nil, fmt.Errorf("unmarshalling messages from config: %v", err)
	}

	return s, nil
}

type Message struct {
	Time     time.Time
	Context  string
	Question string
	Answer   string
}

type Service struct {
	logger   *slog.Logger
	id       *uuid.UUID
	cfg      config.Dependency
	gpt      chatgpt.Dependency
	context  string
	messages []Message
}

func (s *Service) Init(mesh servicemesh.Mesh) {

}

func (s *Service) Name() string {
	const maxLength = 6

	id := s.UUID().String()

	if len(id) <= maxLength {
		return id
	}

	// Truncate the input string to the maxLength
	id = id[:maxLength]

	return fmt.Sprintf("ChatGPT Agent %s", id)
}

func (s *Service) DependenciesResolved() bool {
	if s.cfg == nil {
		return false
	}

	if s.gpt == nil {
		return false
	}

	return true
}

func (s *Service) ResolveDependencies(services []servicemesh.Service) {
	for _, service := range services {
		switch candidate := service.(type) {
		case config.Dependency:
			s.cfg = candidate
		case chatgpt.Dependency:
			s.gpt = candidate
		}
	}
}

func (s *Service) SetLogger(logger *slog.Logger) {
	s.logger = logger
}

func (s *Service) Logger() *slog.Logger {
	return s.logger
}

func (s *Service) ConfigFileName() string {
	return "chatgpt_agents.json"
}

func (s *Service) UUID() uuid.UUID {
	if s.id == nil {
		id := uuid.New()
		s.id = &id
	}

	return *s.id
}

func (s *Service) Context() string {
	return s.context
}

func (s *Service) SetContext(ctx string) {
	s.context = ctx
}

func (s *Service) Ask(question string) (response string, err error) {
	message := Message{
		Time:     time.Now(),
		Context:  s.context,
		Question: question,
	}

	res, err := s.gpt.Ask(s.context, question)
	if err != nil {
		return "", fmt.Errorf("asking GPT a question: %v", err)
	}

	if len(res.Choices) < 1 {
		return
	}

	message.Answer = res.Choices[0].Message.Content

	s.messages = append(s.messages, message)

	return message.Answer, nil
}

func (s *Service) Messages() []Message {
	return s.messages
}
