package main

import (
	"log/slog"

	"github.com/gravestench/servicemesh"

	"torchbearer/pkg/services/account"
	"torchbearer/pkg/services/adventurer"
	"torchbearer/pkg/services/config"
	"torchbearer/pkg/services/dice"
	"torchbearer/pkg/services/email"
	"torchbearer/pkg/services/phase"
	"torchbearer/pkg/services/records"
	"torchbearer/pkg/services/session"
	"torchbearer/pkg/services/webRouter"
	"torchbearer/pkg/services/webServer"
	"torchbearer/pkg/services/world"
)

const (
	configDirectory = "~/.config/torchbearer"
)

func main() {
	rt := servicemesh.New("Game Master")
	rt.SetLogLevel(slog.LevelInfo)

	// generic plumbing services, used by most other services
	rt.Add(&config.Service{RootDirectory: configDirectory})
	//rt.Add(&tui.Service{})

	// account/identity management services
	rt.Add(&account.Service{})
	rt.Add(&email.Service{})

	// web services
	rt.Add(&webRouter.Service{})
	rt.Add(&webServer.Service{})

	// external API integration services
	//rt.Add(&chatgpt.Service{})

	// torchbearer game services
	rt.Add(&records.Service{})
	rt.Add(&phase.Service{})
	rt.Add(&session.Service{})
	rt.Add(&dice.Service{})
	rt.Add(&adventurer.Service{})
	rt.Add(&world.Service{})

	rt.Run()
}
