package main

import (
	"math/rand"
	"time"

	"github.com/gravestench/runtime"

	"torchbearer/pkg/services/adventurer"
	"torchbearer/pkg/services/config"
	"torchbearer/pkg/services/dice"
	"torchbearer/pkg/services/phase"
	"torchbearer/pkg/services/records"
	"torchbearer/pkg/services/session"
	"torchbearer/pkg/services/tui"
	"torchbearer/pkg/services/world"
)

const (
	configDirectory = "~/.config/torchbearer"
)

func main() {
	rand.Seed(time.Now().Unix())

	rt := runtime.New("Game Master")

	rt.Add(&config.Service{RootDirectory: configDirectory})
	rt.Add(&tui.Service{})
	rt.Add(&phase.Service{})
	rt.Add(&session.Service{})
	rt.Add(&dice.Service{})
	rt.Add(&adventurer.Service{})
	rt.Add(&world.Service{})
	rt.Add(&records.Service{})

	rt.Run()
}
