package webServer

import (
	"github.com/gravestench/runtime"

	"torchbearer/pkg/services/config"
)

var (
	_ runtime.Service         = &Service{}
	_ runtime.HasLogger       = &Service{}
	_ config.HasDefaultConfig = &Service{}
	_ IsWebServer             = &Service{}
)

type Dependency = IsWebServer

type IsWebServer interface {
	RestartServer()
	StartServer()
	StopServer()
}
