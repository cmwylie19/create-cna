package cmd

import (
	"github.com/cmwylie19/create-ddis-app/pkg/log"
	"github.com/cmwylie19/create-ddis-app/pkg/scaffold"
	"github.com/spf13/cobra"
)

var (
	Name    string
	Type    string
	Port    string
	Metrics bool
	Verbose bool
)

func getScaffoldCommand(logger log.Logger) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "scaffold",
		Short: "Scaffold a new project",
		Long: `Scaffold a new project based on the langauge and type of project you want to create.

Usage:
  create-cna scaffold --name=<project-name> --type=<project-type> --port=<port> --metrics=<metrics>

Example:
  create-cna scaffold --name=go-server --type=go --port=8080 --metrics=true

  create-cna scaffold --name=migration-job --type=go --port=9191 --metrics=false --verbose=false
`,
		Run: func(cmd *cobra.Command, args []string) {
			// call create scaffold
			var logg log.Logger = &log.Log{Debug: Verbose}
			scaffold.NewScaffold(Name, Type, Port, Metrics, Verbose, logg).Create()

		},
	}
	cmd.Flags().StringVar(&Name, "name", "", "Name of the project")
	cmd.Flags().StringVar(&Type, "type", "", "Type of the project: go, rust, python-batch, python, frontend")
	cmd.Flags().StringVar(&Port, "port", "", "Port of the project: 8080")
	cmd.Flags().BoolVar(&Metrics, "metrics", false, "Metrics of the project: true or false")
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
