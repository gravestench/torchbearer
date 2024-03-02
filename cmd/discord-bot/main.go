package main

import (
	"github.com/rs/zerolog"

	"torchbearer/pkg/services/account"
	"torchbearer/pkg/services/adventurer"
	"torchbearer/pkg/services/config"
	"torchbearer/pkg/services/dice"
	"torchbearer/pkg/services/discord"
	"torchbearer/pkg/services/phase"
	"torchbearer/pkg/services/records"
	"torchbearer/pkg/services/session"
	"torchbearer/pkg/services/world"
)

const (
	configDirectory = "~/.config/torchbearer"
)

func main() {
	rt := servicemesh.New("Torchbearer Discord Bot")
	rt.SetLogLevel(zerolog.InfoLevel)

	rt.Add(&config.Service{RootDirectory: configDirectory})
	//rt.Add(&tui.Service{})
	rt.Add(&phase.Service{})
	rt.Add(&session.Service{})
	rt.Add(&dice.Service{})
	rt.Add(&adventurer.Service{})
	rt.Add(&world.Service{})
	rt.Add(&records.Service{})
	rt.Add(&discord.Service{})
	rt.Add(&account.Service{})

	rt.Run()
}
