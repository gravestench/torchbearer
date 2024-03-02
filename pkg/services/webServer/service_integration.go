package webServer

import (
	"github.com/gravestench/servicemesh"

	"torchbearer/pkg/services/config"
)

var (
	_ servicemesh.Service     = &Service{}
	_ servicemesh.HasLogger   = &Service{}
	_ config.HasDefaultConfig = &Service{}
	_ IsWebServer             = &Service{}
)

type Dependency = IsWebServer

type IsWebServer interface {
	RestartServer()
	StartServer()
	StopServer()
}
