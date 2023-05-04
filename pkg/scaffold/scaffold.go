package scaffold

import (
	"github.com/cmwylie19/create-ddis-app/pkg/helpers"
	"github.com/cmwylie19/create-ddis-app/pkg/log"
)

type Scaffold struct {
	Type       string `json:"type"`
	Port       string `json:"port"`
	Name       string `json:"name"`
	Telemetry  bool   `json:"telemetry"`
	Metrics    bool   `json:"metrics"`
	Controller string `json:"controller"`
	Verbose    bool   `json:"verbose"`
	// Logging interface
	Logger log.Logger `json:"logger"`
}

func NewScaffold(name, type_, port, controller string, telemetry, metrics, verbose bool, logger log.Logger) *Scaffold {
	return &Scaffold{
		Type:       type_,
		Port:       port,
		Name:       name,
		Telemetry:  telemetry,
		Metrics:    metrics,
		Controller: controller,
		Verbose:    verbose,
		Logger:     logger,
	}
}

func (s *Scaffold) Create() {

	// Create the project directory
	err := helpers.CreateDir(s.Name)
	if err != nil {
		s.Logger.Errorf("Error creating project directory: %s", err.Error())
	}
	s.Logger.Infof("Created project directory: %s", s.Name)

	// create the Kubernetes directory
	err = helpers.CreateDir(s.Name + "/kubernetes")
	if err != nil {
		s.Logger.Errorf("Error creating Kubernetes directory: %s", err.Error())
	}
	s.Logger.Infof("Created kubernetes directory and manifests: %s", s.Name)

	// create the Docker (Build) directory
	err = helpers.CreateDir(s.Name + "/build")
	if err != nil {
		s.Logger.Errorf("Error creating Docker (Build) directory: %s", err.Error())
	}
	s.Logger.Infof("Created build directory: %s", s.Name)

	s.Logger.Print("Scaffolded project successfully!")

}
