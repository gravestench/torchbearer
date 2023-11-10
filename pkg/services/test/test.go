package test

import (
	"torchbearer/pkg/models"
)

type Test struct {
	models.Dice
	Description string
	models.Obstruction
	Factors
}

type Factors []models.Factor
