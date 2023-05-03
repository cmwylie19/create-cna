package cmd

import (
	"github.com/cmwylie19/create-ddis-app/interal/log"
	"github.com/spf13/cobra"
)

var (
	Name       string
	Type       string
	Port       string
	Telemetry  bool
	Metrics    bool
	Controller string
)

func getScaffoldCommand(log log.Logger) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "scaffold",
		Short: "Scaffold a new project",
		Long: `Scaffold a new project based on the type of project you want to create.

Usage:
  create-ddis-app scaffold --name=<project-name> --type=<project-type> --port=<port> --telemetry=<telemetry> --metrics=<metrics> --controller=<controller>
`,
		Run: func(cmd *cobra.Command, args []string) {
		},
	}
	cmd.Flags().StringVar(&Name, "name", "", "Name of the project")
	cmd.Flags().StringVar(&Type, "type", "", "Type of the project")
	cmd.Flags().StringVar(&Port, "port", "", "Port of the project")
	cmd.Flags().BoolVar(&Telemetry, "telemetry", false, "Telemetry of the project")
	cmd.Flags().BoolVar(&Metrics, "metrics", false, "Metrics of the project")
	cmd.Flags().StringVar(&Controller, "controller", "", "Controller of the project")
	err := cmd.MarkFlagRequired("name")
	if err != nil {
		log.Errorf("Name required: %s", err.Error())
	}

	err = cmd.MarkFlagRequired("type")
	if err != nil {
		log.Errorf("Type required: %s", err.Error())
	}

	if Type != "python-batch" {
		err = cmd.MarkFlagRequired("port")
		if err != nil {
			log.Errorf("Port required: %s", err.Error())
		}
	}
	return cmd
}
