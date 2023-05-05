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
	err = helpers.CreateDir(s.Name + "/manifests/k8s")
	if err != nil {
		s.Logger.Errorf("Error creating Kubernetes directory: %s", err.Error())
	}
	s.Logger.Infof("Created kubernetes directory.")

	// create the observability directory
	err = helpers.CreateDir(s.Name + "/manifests/observability")
	if err != nil {
		s.Logger.Errorf("Error creating observability directory: %s", err.Error())
	}
	s.Logger.Info("Created observability directory.")
	// create the Docker (Build) directory
	err = helpers.CreateDir(s.Name + "/build")
	if err != nil {
		s.Logger.Errorf("Error creating Docker (Build) directory: %s", err.Error())
	}
	s.Logger.Info("Created build directory.")

	// create Dockerfile
	err = helpers.CreateDockerfile(s.Name+"/build", s.Type, s.Name)
	if err != nil {
		s.Logger.Errorf("Error creating Dockerfile: %s", err.Error())
	}
	s.Logger.Info("Created Dockerfile")

	// create Makefile
	err = helpers.CreateMakefile(s.Name, s.Type)
	if err != nil {
		s.Logger.Errorf("Error creating Dockerfile: %s", err.Error())
	}
	s.Logger.Info("Created Makefile")

	// create .gitignore
	err = helpers.CreateGitIgnore(s.Name)
	if err != nil {
		s.Logger.Errorf("Error creating .gitignore: %s", err.Error())
	}
	s.Logger.Info("Created .gitignore")

	// create main
	err = helpers.CreateMainfile(s.Name, s.Type)
	if err != nil {
		s.Logger.Errorf("Error creating main file: %s", err.Error())
	}
	s.Logger.Info("Created main file.")

	// create servicemonitor
	err = helpers.CreateServiceMonitor(s.Name+"/manifests/observability", s.Name)
	if err != nil {
		s.Logger.Errorf("Error creating servicemonitor: %s", err.Error())
	}
	s.Logger.Info("Created ServiceMonitor.")

	// create Kubernetes manifests
	err = helpers.CreateKubernetesManifests(s.Name+"/manifests/k8s", s.Type, s.Controller, s.Port, s.Name)
	if err != nil {
		s.Logger.Errorf("Error creating kubernetes manifests: %s", err.Error())
	}
	s.Logger.Info("Created Kubernetes Manifests.")

	s.Logger.Print("Scaffolded project successfully!")

}
