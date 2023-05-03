package scaffold

import (
	"github.com/create-ddis-app/pkg/helpers"
)

type Scaffold struct {
	Type       string `json:"type"`
	Port       string `json:"port"`
	Name       string `json:"name"`
	Telemetry  bool   `json:"telemetry"`
	Metrics    bool   `json:"metrics"`
	Controller string `json:"controller"`
}

func newScaffold(name, type_, port string, telemetry, metrics bool) *Scaffold {
	return &Scaffold{
		Type:       type_,
		Port:       port,
		Name:       name,
		Telemetry:  telemetry,
		Metrics:    metrics,
		Controller: "controller",
	}
}

func (s *Scaffold) Create() {

	// Create the project directory
	err := helpers.CreateDir(s.Name)
	if err != nil {

	}

}
