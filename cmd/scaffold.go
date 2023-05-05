package cmd

import (
	"github.com/cmwylie19/create-ddis-app/pkg/log"
	"github.com/cmwylie19/create-ddis-app/pkg/scaffold"
	"github.com/spf13/cobra"
)

var (
	Name       string
	Type       string
	Port       string
	Telemetry  bool
	Metrics    bool
	Controller string
	Verbose    bool
)

func getScaffoldCommand(logger log.Logger) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "scaffold",
		Short: "Scaffold a new project",
		Long: `Scaffold a new project based on the langauge and type of project you want to create.

Usage:
  create-ddis-app scaffold --name=<project-name> --type=<project-type> --port=<port> --telemetry=<telemetry> --metrics=<metrics> --controller=<controller>

Example:
  create-ddis-app scaffold --name=go-server --type=go --port=8080 --telemetry=true --metrics=true --controller=deployment

  create-ddis-app scaffold --name=migration-job --type=python-batch --port=8080 --telemetry=false --metrics=false --controller=pod --verbose=false

  create-ddis-app scaffold --name=salesforce-accounts --type=python --port=8080 --telemetry=true --metrics=true --controller=deployment

  create-ddis-app scaffold --name=image-analyzer --type=rust --port=8080 --telemetry=true --metrics=true --controller=daemonset
`,
		Run: func(cmd *cobra.Command, args []string) {
			// call create scaffold
			var logg log.Logger = &log.Log{Debug: Verbose}
			scaffold.NewScaffold(Name, Type, Port, Controller, Telemetry, Metrics, Verbose, logg).Create()

		},
	}
	cmd.Flags().StringVar(&Name, "name", "", "Name of the project")
	cmd.Flags().StringVar(&Type, "type", "", "Type of the project: go, rust, python-batch, python, frontend")
	cmd.Flags().StringVar(&Port, "port", "", "Port of the project: 8080")
	cmd.Flags().BoolVar(&Telemetry, "telemetry", false, "Telemetry of the project: true or false")
	cmd.Flags().BoolVar(&Metrics, "metrics", false, "Metrics of the project: true or false")
	cmd.Flags().StringVar(&Controller, "controller", "", "Controller of the project: Deployment (default), DaemonSet, StatefulSet, CronJob, Job, Pod (python-batch)")
	cmd.Flags().BoolVar(&Verbose, "verbose", true, "Verbose logs: true (default) or false")

	err := cmd.MarkFlagRequired("name")
	if err != nil {
		logger.Errorf("Name required: %s", err.Error())
	}

	err = cmd.MarkFlagRequired("type")
	if err != nil {
		logger.Errorf("Type required: %s", err.Error())
	}

	if Type != "python-batch" {
		err = cmd.MarkFlagRequired("port")
		if err != nil {
			logger.Errorf("Port required: %s", err.Error())
		}
	}
	return cmd
}
