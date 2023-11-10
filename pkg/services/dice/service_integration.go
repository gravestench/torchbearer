package dice

import (
	"torchbearer/pkg/models"
)

type DiceManager interface {
	Roll(int) models.Dice
	Reroll6s(models.Dice) models.Dice
}
