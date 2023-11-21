package procedure

import (
	"fmt"
	"regexp"
	"strings"
)

type Step struct {
	Key             string
	Procedure       *Procedure
	Prompt          string
	Default         string
	Choices         []Choice
	Answer          string
	ValidatorRegex  string
	ValidatorPrompt string
	ShouldSkip      func() bool
	OnComplete      func()
	isComplete      bool
}

type Choice struct {
	Name, Description string
}

func (s *Step) Complete() error {
	if err := s.Validate(); err != nil {
		return err
	}

	if s.OnComplete != nil && !s.isComplete {
		s.OnComplete()
	}

	s.isComplete = true

	return nil
}

func (s *Step) Validate() (err error) {
	if s.ShouldSkip != nil {
		if s.ShouldSkip() {
			return nil
		}
	}

	if s.ValidatorRegex == "" && len(s.Choices) == 0 {
		return nil
	}

	if len(s.Choices) > 0 {
		var choiceNames []string
		for _, choice := range s.Choices {
			choiceNames = append(choiceNames, choice.Name)
		}

		s.ValidatorRegex = fmt.Sprintf("%s", strings.Join(choiceNames, "|"))
		s.ValidatorPrompt = fmt.Sprintf("Must be one of [%s]", strings.Join(choiceNames, ", "))
	}

	matched, err := regexp.MatchString(s.ValidatorRegex, s.Answer)
	if !matched {
		err = fmt.Errorf("invalid answer")
		if s.ValidatorPrompt != "" {
			err = fmt.Errorf("invalid answer: %s", s.ValidatorPrompt)
		}
		return
	}

	return nil
}
