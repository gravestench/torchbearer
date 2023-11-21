package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/gravestench/runtime"

	"torchbearer/pkg/services/adventurer"
	"torchbearer/pkg/services/config"
	"torchbearer/pkg/services/dice"
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
	rt := runtime.New("Game Master")

	a := &adventurer.Service{}

	rt.Add(&config.Service{RootDirectory: configDirectory})
	rt.Add(&phase.Service{})
	rt.Add(&session.Service{})
	rt.Add(&dice.Service{})
	rt.Add(a)
	rt.Add(&world.Service{})
	rt.Add(&records.Service{})
	rt.Add(&webRouter.Service{})
	rt.Add(&webServer.Service{})

	go func() {
		for !a.DependenciesResolved() {
			time.Sleep(time.Millisecond * 5)
		}

		for _, procedure := range a.Procedures() {
			p := procedure.New()

			for {
				step := p.NextStep()
				if step == nil {
					break
				}

				fmt.Println(step.Prompt)
				for _, choice := range step.Choices {
					fmt.Printf("%s: %s\r\n", choice.Name, choice.Description)
				}

				step.Answer = readInputFromStdin()
				if err := step.Complete(); err != nil {
					fmt.Println(err)
					continue
				}
			}
		}
	}()

	rt.Run()
}

func readInputFromStdin() string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter some text: ")

	if scanner.Scan() {
		return scanner.Text()
	}

	if scanner.Err() != nil {
		fmt.Println("Error:", scanner.Err())
	}

	return ""
}
