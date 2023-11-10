package dice

import (
	"time"

	"github.com/gravestench/runtime"
	"github.com/rs/zerolog"

	"torchbearer/pkg/models"
)

type Service struct {
	logger *zerolog.Logger
	rolls  map[time.Time]models.Dice
}

func (s *Service) Roll(n int) models.Dice {
	dice := make(models.Dice, n)

	for idx := range dice {
		dice[idx] = &models.Die{}
		dice[idx].Roll()
	}

	s.logger.Info().Msgf("Rolling (%d dice): %s", n, dice.String())

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
		s.logger.Info().Msgf("Reroll 6s (%d additional): %s", len(additional), additional.String())
	}

	// recurse
	if continueRolling {
		additional = s.Reroll6s(additional)
	}

	return append(dice, additional...)
}

func (s *Service) Init(rt runtime.Runtime) {
	s.rolls = make(map[time.Time]models.Dice)
}

func (s *Service) Name() string {
	return "Dice"
}

func (s *Service) BindLogger(logger *zerolog.Logger) {
	s.logger = logger
}

func (s *Service) Logger() *zerolog.Logger {
	return s.logger
}
