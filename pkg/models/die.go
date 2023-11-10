package models

import (
	"math/rand"
	"strings"
)

const sides = 6

type Die struct {
	state int
}

func (d *Die) Roll() int {
	d.state = rand.Intn(sides)

	return d.state
}

func (d *Die) String() string {
	if d.state < sides>>1 {
		return "Wyrm"
	}

	if d.state == sides-1 {
		return "Success*"
	}

	return "Success"
}

func (d *Die) IsSix() bool {
	return d.state == sides-1
}

type Dice []*Die

func (dd Dice) String() string {
	var s []string

	for _, d := range dd {
		s = append(s, d.String())
	}

	return strings.Join(s, ", ")
}

func (dd Dice) NumFail() (s int) {
	for _, d := range dd {
		if d.state < 3 {
			s++
		}
	}

	return
}

func (dd Dice) NumSuccess() (s int) {
	for _, d := range dd {
		if d.state > 2 {
			s++
		}
	}

	return
}

func (dd Dice) NumSuccess6() (s int) {
	for _, d := range dd {
		if d.state >= 5 {
			s++
		}
	}

	return
}
