package dice

import (
	"log/slog"
	"time"

	"github.com/gravestench/servicemesh"

	"torchbearer/pkg/models"
)

type Service struct {
	logger *slog.Logger
	rolls  map[time.Time]models.Dice
}

func (s *Service) Init(mesh servicemesh.Mesh) {
	s.rolls = make(map[time.Time]models.Dice)
}

func (s *Service) Name() string {
	return "Dice"
}

func (s *Service) Roll(n int) models.Dice {
	dice := make(models.Dice, n)

	for idx := range dice {
		dice[idx] = &models.Die{}
		dice[idx].Roll()
	}

	s.logger.Info("rolling dice", "count", n, "result", dice.String())

	s.rolls[time.Now()] = dice

	return dice
}

func (s *Service) Reroll6s(dice models.Dice) models.Dice {
	var numAdditional int

	// determine additional dice to roll
	for _, d := range dice {
		if d.IsSix() {
			numAdditional++
		}
	}

	if numAdditional == 0 {
		return dice
	}

	additional := s.Roll(numAdditional)

	var continueRolling bool

	// roll those dice
	for _, d := range additional {
		d.Roll()
		if d.IsSix() {
			continueRolling = true
		}
	}

	if len(additional) > 0 {
		s.logger.Info("rerolling dice", "count", len(additional), "result", additional.String())
	}

	// recurse
	if continueRolling {
		additional = s.Reroll6s(additional)
	}

	return append(dice, additional...)
}

func (s *Service) SetLogger(logger *slog.Logger) {
	s.logger = logger
}

func (s *Service) Logger() *slog.Logger {
	return s.logger
}
